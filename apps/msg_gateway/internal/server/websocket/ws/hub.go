package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Hub struct {
	serverId          int
	rwLock            sync.RWMutex
	upgrader          websocket.Upgrader
	registerChan      chan *Client
	unregisterChan    chan *Client    // 只在Client调用closeConn()函数时触发
	readChan          chan *Message   // 客户端发送的消息
	msgCallback       MessageCallback // 回调
	clients           *RwMap          // key1:uid key2:platform
	access            map[int64]int64 // 访问间隔
	onlineConnections int64           // 在线连接数
}

func NewHub(serverId int, msgCallback MessageCallback) *Hub {
	return &Hub{
		serverId: serverId,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  WS_READ_BUFFER_SIZE,
			WriteBufferSize: WS_WRITE_BUFFER_SIZE,
		},
		registerChan:   make(chan *Client, WS_CHAN_CLIENT_REGISTER),
		unregisterChan: make(chan *Client, WS_CHAN_CLIENT_UNREGISTER),
		readChan:       make(chan *Message, WS_CHAN_SERVER_READ_MESSAGE),
		msgCallback:    msgCallback,
		clients:        NewRwMap(),
		access:         make(map[int64]int64),
	}
}

func (h *Hub) registerClient(client *Client) {
	var (
		ok  bool
		cli *Client
	)

	if cli, ok = h.clients.Get(client.key); ok == false {
		h.clients.Set(client.key, client)
		atomic.AddInt64(&h.onlineConnections, 1)
		fmt.Println("R在线用户数量:", h.onlineConnections)
		return
	}

	if client.onlineAt > cli.onlineAt {
		h.clients.Set(client.key, client)
		h.close(cli)
		return
	}
	h.close(client)
}

func (h *Hub) close(client *Client) {
	client.Close()
}

func (h *Hub) unregisterClient(client *Client) {
	var (
		ok  bool
		cli *Client
	)
	if cli, ok = h.clients.Get(client.key); ok == false {
		return
	}
	if client == cli {
		h.clients.Delete(client.key)
		atomic.AddInt64(&h.onlineConnections, -1)
	}
	fmt.Println("U在线用户数量:", h.onlineConnections)
}

func (h *Hub) Online(online bool) {
	if online {
		atomic.AddInt64(&h.onlineConnections, 1)
	} else {
		atomic.AddInt64(&h.onlineConnections, -1)
	}
}

func (h *Hub) Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("hub run err:", r, string(debug.Stack()))
		}
	}()

	var (
		index int
		loop  = 100
	)
	for index = 0; index < loop; index++ {
		var (
			client *Client
			msg    *Message
		)
		go func() {
			for {
				select {
				case client = <-h.registerChan:
					h.registerClient(client)
				case client = <-h.unregisterChan:
					h.unregisterClient(client)
				case msg = <-h.readChan:
					h.messageHandler(msg)
				}
			}
		}()
	}
}

func (h *Hub) IsOnline(uid int64, platform int32) (ok bool) {
	_, ok = h.clients.Get(clientKey(uid, platform))
	return
}

func (h *Hub) SendMessage(uid int64, platform int32, message []byte) (result int32) {
	result = WS_CLIENT_OFFLINE
	var (
		cli *Client
		ok  bool
	)
	if cli, ok = h.clients.Get(clientKey(uid, platform)); ok == false {
		return
	}
	cli.Send(message)
	result = WS_SEND_MSG_SUCCESS
	return
}

func (h *Hub) messageHandler(msg *Message) {
	h.msgCallback(msg)
}

func (h *Hub) wsHandler(c *gin.Context) {
	var (
		uidVal   interface{}
		pidVal   interface{}
		exists   bool
		uid      int64
		platform int32
		conn     *websocket.Conn
		client   *Client
		lastTs   int64
		nowTs    int64
		err      error
	)

	if h.onlineConnections >= WS_MAX_CONNECTIONS {
		httpError(c, ERROR_CODE_WS_EXCEED_MAX_CONNECTIONS, ERROR_WS_EXCEED_MAX_CONNECTIONS)
		return
	}
	uidVal, exists = c.Get(WS_KEY_UID)
	if exists == false {
		httpError(c, ERROR_CODE_HTTP_UID_DOESNOT_EXIST, ERROR_HTTP_UID_DOESNOT_EXIST)
		return
	}
	pidVal, exists = c.Get(WS_KEY_PLATFORM)
	if exists == false {
		httpError(c, ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	uid = int64(uidVal.(float64))
	if uid == 0 {
		httpError(c, ERROR_CODE_HTTP_UID_DOESNOT_EXIST, ERROR_HTTP_UID_DOESNOT_EXIST)
		return
	}
	platform = int32(pidVal.(float64))

	nowTs = time.Now().UnixNano() / 1e6
	h.rwLock.Lock()
	lastTs, _ = h.access[uid]
	h.access[uid] = nowTs
	h.rwLock.Unlock()
	if nowTs-lastTs < WS_MINIMUM_TIME_INTERVAL {
		httpError(c, ERROR_CODE_HTTP_REQUEST_TOO_MUNDANE, ERROR_HTTP_REQUEST_TOO_MUNDANE)
		return
	}

	if conn, err = h.upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		// 协议升级失败
		httpError(c, ERROR_CODE_HTTP_UPGRADER_FAILED, err.Error())
		return
	}
	client = newClient(h, conn, uid, platform)
	h.registerChan <- client

	go client.writeLoop()
	go client.readLoop()
}
