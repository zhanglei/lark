package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"runtime/debug"
	"sync"
	"time"
)

type Client struct {
	rwLock    sync.RWMutex
	hub       *Hub
	conn      *websocket.Conn
	uid       int64 // 用户ID
	platform  int32 // 平台ID
	key       string
	onlineAt  int64 // 上线时间戳（毫秒）
	sendChan  chan []byte
	hbChan    chan int
	closeChan chan []byte
	closed    bool
}

func newClient(hub *Hub, conn *websocket.Conn, uid int64, platform int32) *Client {
	cli := &Client{
		hub:       hub,
		conn:      conn,
		uid:       uid,
		platform:  platform,
		key:       clientKey(uid, platform),
		onlineAt:  time.Now().UnixNano() / 1e6,
		sendChan:  make(chan []byte, WS_CHAN_CLIENT_SEND_MESSAGE),
		closeChan: make(chan []byte),
		hbChan:    make(chan int, WS_CHAN_HEARTBEAT_MESSAGE),
	}
	//cli.debug()
	return cli
}

func (c *Client) debug() {
	go func() {
		allTicker := time.NewTicker(time.Second * 5)
		defer allTicker.Stop()
		for {
			select {
			case <-allTicker.C:
				fmt.Println(c.uid, c.closed, len(c.sendChan))
			}
		}
	}()
}

func (c *Client) closeConn() {
	c.rwLock.Lock()
	defer func() {
		c.rwLock.Unlock()
	}()

	if c.closed == true {
		return
	}
	c.closed = true

	close(c.sendChan)
	close(c.closeChan)
	close(c.hbChan)

	c.conn.Close()
	c.hub.unregisterChan <- c
}

func (c *Client) readLoop() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket read err:", r, string(debug.Stack()))
		}
		c.closeConn()
	}()

	var (
		msgType int
		bufMsg  []byte
		message *Message
		err     error
	)

	c.conn.SetReadLimit(WS_MAX_MESSAGE_SIZE)
	c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT)); return nil })

	for {
		if msgType, bufMsg, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			break
		}
		if c.closed == true {
			break
		}
		if msgType == websocket.PongMessage || msgType == websocket.PingMessage {
			c.hbChan <- msgType
			continue
		}
		if msgType == websocket.CloseMessage {
			return
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		if len(bufMsg) == 0 {
			continue
		}
		message = &Message{
			Client: c,
			Body:   bufMsg,
		}
		c.hub.readChan <- message
	}
}

func (c *Client) writeLoop() {
	pingTicker := time.NewTicker(WS_PING_PERIOD)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket write err:", r, string(debug.Stack()))
		}
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err        error
		message    []byte
		msgType    int
		ok         bool
		wc         io.WriteCloser
		chanLength int
		msgLength  int
		index      int
	)

	for {
		select {
		case message, ok = <-c.sendChan:
			if ok == false {
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			wc, err = c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			msgLength = len(message)
			wc.Write(message)
			chanLength = len(c.sendChan)
			for index = 0; index < chanLength; index++ {
				if msgLength >= WS_WRITE_BUFFER_SIZE {
					break
				}
				message = <-c.sendChan
				msgLength += len(message)
				wc.Write(message)
			}
			if err = wc.Close(); err != nil {
				return
			}
		case msgType, ok = <-c.hbChan:
			if ok == false {
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			if msgType == websocket.PingMessage {
				if err = c.conn.WriteMessage(websocket.PongMessage, nil); err != nil {
					return
				}
			}
		case _, ok = <-pingTicker.C:
			if ok == false {
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case message, ok = <-c.closeChan:
			if ok == false {
				return
			}
			c.conn.WriteMessage(websocket.CloseMessage, message)
			return
		}
	}
}

func (c *Client) Send(message []byte) {
	c.rwLock.RLock()
	defer func() {
		c.rwLock.RUnlock()
		if r := recover(); r != nil {
			fmt.Println("socket send err:", r, string(debug.Stack()))
		}
	}()

	if c.closed == true {
		return
	}
	if len(c.sendChan) >= WS_CHAN_CLIENT_SEND_MESSAGE {
		return
	}
	c.sendChan <- message
}

func (c *Client) Close() {
	c.rwLock.RLock()
	defer func() {
		c.rwLock.RUnlock()
	}()

	if c.closed == true {
		return
	}
	c.closeChan <- WS_MSG_BUF_CLOSE
}
