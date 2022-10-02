package ws

import (
	"lark/pkg/common/xgin"
	"lark/pkg/middleware"
)

type WServer struct {
	port     int
	serverId int
	hub      *Hub
	gin      *xgin.GinServer
}

func NewWServer(port int, serverId int, callback MessageCallback) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		port:     port,
		serverId: serverId,
		hub:      NewHub(serverId, callback),
		gin:      xgin.NewGinServer(),
	}
	//TODO:待开启验证
	ws.gin.Engine.Use(middleware.JwtAuth())
	ws.gin.Engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	ws.hub.Run()
	ws.gin.Run(ws.port)
}

func (ws *WServer) SendMessage(uid int64, platform int32, message []byte) (result int32) {
	return ws.hub.SendMessage(uid, platform, message)
}

func (ws *WServer) IsOnline(uid int64, platform int32) (ok bool) {
	return ws.hub.IsOnline(uid, platform)
}
