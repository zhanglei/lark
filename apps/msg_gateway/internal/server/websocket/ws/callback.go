package ws

type Message struct {
	Client *Client `json:"client"`
	// 消息本体
	Body []byte `json:"body"`
}

type MessageCallback func(msg *Message)

type SendResult struct {
	Platform int32
	Code     int32
}
