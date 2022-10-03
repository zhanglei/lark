package client

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
	"net/url"
	"runtime/debug"
	"sync"
	"time"
)

type Client struct {
	rwLock    sync.RWMutex
	mgr       *Manager
	conn      *websocket.Conn
	uid       int64 // 用户ID
	platform  int32 // 平台ID
	key       string
	onlineAt  int64 // 上线时间戳（毫秒）
	sendChan  chan []byte
	hbChan    chan int
	closeChan chan []byte
	closed    bool
	nickname  string
}

func NewClient(uid int64, mgr *Manager) (client *Client) {
	var (
		u url.URL
		//q      url.Values
		ts     int64
		conn   *websocket.Conn
		header map[string][]string
		token  string
		err    error
	)
	ts = time.Now().Unix()
	client = &Client{
		mgr:       mgr,
		conn:      nil,
		uid:       uid,
		platform:  1,
		onlineAt:  ts,
		sendChan:  make(chan []byte, WS_CHAN_CLIENT_SEND_MESSAGE),
		closeChan: make(chan []byte),
		hbChan:    make(chan int, WS_CHAN_HEARTBEAT_MESSAGE),
		closed:    false,
		nickname:  "昵称:" + utils.Int64ToStr(uid),
	}

	u = url.URL{Scheme: "ws", Host: "127.0.01:32001", Path: "/"}
	/*
		q := u.Query()
		q.Set("uid", uid)
		q.Set("platform", "1")
		u.RawQuery = q.Encode()
	*/

	token, _ = xjwt.CreateToken(uid, 1)
	header = make(map[string][]string)
	header[WS_KEY_UID] = []string{utils.Int64ToStr(uid)}
	header[WS_KEY_PLATFORM] = []string{"1"}
	header[WS_KEY_COOKIE] = []string{token}

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		client.closed = true
		fmt.Println("创建连接失败")
		return
	}
	client.conn = conn
	go client.write()
	go client.read()
	return
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
	c.mgr.unregister <- c
}

/*
func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket read err:", r, string(debug.Stack()))
		}
		c.closeConn()
	}()

	var (
		msgType int
		r       io.Reader
		all     []byte
		slices  [][]byte
		buf     []byte
		err     error
	)

	c.conn.SetReadLimit(WS_MAX_MESSAGE_SIZE)
	c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT)); return nil })

	for {
		msgType, r, err = c.conn.NextReader()
		if err != nil {
			return
		}
		if c.closed == true {
			break
		}
		if msgType == websocket.PongMessage || msgType == websocket.PingMessage {
			c.hbChan <- msgType
			continue
		}
		if msgType == websocket.CloseMessage {
			c.closeChan <- nil
			return
		}
		if msgType != websocket.BinaryMessage {
			continue
		}

		all, err = ioutil.ReadAll(r)
		if err != nil {
			return
		}
		if len(all) == 0 {
			continue
		}
		slices = bytes.Split(all, WS_MSG_BUF_NEWLINE)
		if len(slices) == 0 {
			continue
		}
		for _, buf = range slices {
			c.messageHandler(buf)
		}
	}
}
*/

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket read err:", r, string(debug.Stack()))
		}
		c.closeConn()
	}()

	var (
		msgType int
		bufMsg  []byte
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
			c.closeChan <- nil
			return
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		if len(bufMsg) == 0 {
			continue
		}
		c.decode(bufMsg)
	}
}

const (
	MessageLength   uint32 = 4
	MessageTopic    uint32 = 8
	MessageSubtopic uint32 = 12
	MessageType     uint32 = 16
)

func (c *Client) decode(buf []byte) {
	msgTimer.UpdateEndTime()
	var (
		lengthBuff   []byte
		length       uint32
		topicBuff    []byte
		topic        uint32
		subtopicBuff []byte
		subtopic     uint32
		msgTypeBuff  []byte
		msgType      uint32
		body         []byte
		totalLength  uint32
	)
	var (
		msgCount  int
		msgLength = len(buf)
		resp      *pb_msg.MessageResp
		msg       *pb_msg.SrvChatMessage
	)
	for {
		totalLength = uint32(len(buf))
		if totalLength < MessageType {
			break
		}
		lengthBuff = buf[:MessageLength]
		length = binary.LittleEndian.Uint32(lengthBuff)
		if totalLength < MessageType+length {
			break
		}
		topicBuff = buf[MessageLength:MessageTopic]
		topic = binary.LittleEndian.Uint32(topicBuff)

		subtopicBuff = buf[MessageTopic:MessageSubtopic]
		subtopic = binary.LittleEndian.Uint32(subtopicBuff)

		msgTypeBuff = buf[MessageSubtopic:MessageType]
		msgType = binary.LittleEndian.Uint32(msgTypeBuff)

		body = buf[MessageType : MessageType+length]
		if length > 0 && topic > 0 && subtopic > 0 && msgType > 0 && len(body) > 0 {

		}
		switch pb_enum.MESSAGE_TYPE(msgType) {
		case pb_enum.MESSAGE_TYPE_RESP:
			resp = new(pb_msg.MessageResp)
			proto.Unmarshal(body, resp)
			xlog.Info("应答消息:", topic, subtopic, resp.Code, resp.Msg)
		case pb_enum.MESSAGE_TYPE_NEW:
			msg = new(pb_msg.SrvChatMessage)
			proto.Unmarshal(body, msg)
			xlog.Info("新消息:", topic, subtopic, msg.SeqId, string(msg.Body))
		}
		msgCount++
		if totalLength < MessageType+length+MessageType {
			break
		}
		buf = buf[MessageType+length:]
	}
	if msgCount > 1 {
		xlog.Info("收到合并消息", msgLength, msg.String())
	}
}

func (c *Client) write() {
	pingTicker := time.NewTicker(WS_PING_PERIOD)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket write err:", r, string(debug.Stack()))
		}
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err     error
		message []byte
		msgType int
		ok      bool
	)

	for {
		select {
		case message, ok = <-c.sendChan:
			if ok == false {
				return
			}
			if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
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
		case <-pingTicker.C:
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.closeChan:
			c.conn.WriteMessage(websocket.CloseMessage, nil)
			return
		}
	}
}

func (c *Client) SendUser(receiverId int64) (err error) {
	var (
		ts      int64
		msgBuf  []byte
		msgBody *pb_msg.CliChatMessage
	)
	if c.conn == nil {
		return
	}
	if c.closed == true {
		return
	}
	ts = utils.MillisFromTime(time.Now())
	msgBody = &pb_msg.CliChatMessage{
		CliMsgId:       xsnowflake.NewSnowflakeID(), //客户端唯一消息号
		SenderId:       c.uid,
		ReceiverId:     receiverId,
		SenderPlatform: pb_enum.PLATFORM_TYPE(c.platform),
		ChatId:         3333336666669999990,
		ChatType:       pb_enum.CHAT_TYPE_GROUP,
		MsgType:        1,
		Body:           utils.Str2Bytes("文本聊天消息" + utils.Int64ToStr(c.uid)),
		SentTs:         ts,
	}
	msgBuf, _ = proto.Marshal(msgBody)
	msgBuf, _ = utils.Encode(int32(pb_enum.TOPIC_CHAT), 0, int32(pb_enum.MESSAGE_TYPE_NEW), msgBuf)
	//xlog.Info(len(msgBuf), pb_enum.TOPIC_CHAT, 0)
	c.Send(msgBuf)
	return
}

func (c *Client) Send(message []byte) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed {
		return
	}
	if len(c.sendChan) >= WS_CHAN_CLIENT_SEND_MESSAGE {
		return
	}
	c.sendChan <- message
}

func (c *Client) Close() {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed {
		return
	}
	c.closeChan <- WS_MSG_BUF_CLOSE
}
