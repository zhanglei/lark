```
package client

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"io"
	"io/ioutil"
	"lark/pkg/common/xjwt"
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

	u = url.URL{Scheme: "ws", Host: "lark.com:32001", Path: "/"}
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

func (c *Client) messageHandler(message []byte) {
	msgTimer.UpdateEndTime()

	var (
		svrMsg = new(pb_msg.SrvMessage)
		err    error
	)
	err = proto.Unmarshal(message, svrMsg)
	if err != nil {
		fmt.Println("解析消息错误:", err.Error())
		return
	} else {
		fmt.Println("解析消息成功:", svrMsg.Topic, svrMsg.MsgType)
	}
	return
	if svrMsg.Topic == pb_enum.TOPIC_UNKNOWN_TOPIC {
		return
	}
	if svrMsg.Data == nil {
		return
	}
	if svrMsg.MsgType == pb_enum.SRV_MESSAGE_TYPE_RESP {
		var (
			resp = new(pb_msg.MessageResp)
		)
		proto.Unmarshal(svrMsg.Data, resp)
		//fmt.Println(fmt.Sprintf("收到响应消息 发送者UID:%d, 消息ID:%d, code:%d, msg:%s", c.Uid, resp.MsgId, resp.Code, resp.Msg))
	} else {
		var (
			newMsg = new(pb_msg.SrvChatMessage)
		)
		proto.Unmarshal(svrMsg.Data, newMsg)
		//fmt.Println(fmt.Sprintf("新消息:%v,\n%v", time.Now(), newMsg.String()))
	}
	c.replyMessages(svrMsg)
}

func (c *Client) replyMessages(svrMsg *pb_msg.SrvMessage) (err error) {
	var (
		msgBody     = new(pb_msg.SrvChatMessage)
		sendMsgReq  *pb_msg.CliMessage
		sendMsgBody = new(pb_msg.CliChatMessage)
		sendMsgBuf  []byte
	)
	err = proto.Unmarshal(svrMsg.Data, msgBody)
	if err != nil {
		fmt.Println("解析消息错误 02")
		return
	}

	sendMsgReq = &pb_msg.CliMessage{
		Topic:    svrMsg.Topic,
		SubTopic: svrMsg.SubTopic,
	}
	copier.Copy(sendMsgBody, msgBody)
	sendMsgBody.CliMsgId = xsnowflake.NewSnowflakeID()
	sendMsgBody.ReceiverId = sendMsgBody.SenderId
	sendMsgBody.SenderId = c.uid
	sendMsgReq.Data, _ = proto.Marshal(sendMsgBody)
	sendMsgBuf, _ = proto.Marshal(sendMsgReq)
	time.Sleep(time.Second * 2)
	c.Send(sendMsgBuf)
	return
}

func (c *Client) SendUser(receiverId int64) (err error) {
	var (
		ts       int64
		reqBytes []byte
		sendReq  *pb_msg.CliMessage
		msgBody  *pb_msg.CliChatMessage
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

	sendReq = &pb_msg.CliMessage{
		Topic:    pb_enum.TOPIC_CHAT,
		SubTopic: 0,
	}
	sendReq.Data, _ = proto.Marshal(msgBody)
	reqBytes, _ = proto.Marshal(sendReq)
	c.Send(reqBytes)
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

```