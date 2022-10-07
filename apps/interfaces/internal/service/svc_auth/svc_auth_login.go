package svc_auth

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_auth"
	"lark/pkg/proto/pb_mq"
	"lark/pkg/xhttp"
)

func (s *authService) Login(params *dto_auth.LoginReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req       = new(pb_auth.LoginReq)
		reply     *pb_auth.LoginResp
		loginResp = new(dto_auth.AuthResp)
		ok        bool
	)
	copier.Copy(req, params)
	reply = s.authClient.Login(req)
	if reply == nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetRespInfo(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}

	//TODO:获取服务器ID 测试数据 ServerId: 1
	wsServer := &dto_auth.ServerInfo{
		ServerId: 1,
		Address:  "lark-ws-server.com:32001",
	}
	//更新 wsServer 和 登录平台
	onlineMsg := &pb_mq.UserOnline{
		Uid:      reply.UserInfo.Uid,
		Platform: params.Platform,
		ServerId: wsServer.ServerId,
	}
	ok = s.UserOnline(onlineMsg, resp)
	if ok == false {
		return
	}

	copier.Copy(loginResp, reply)
	loginResp.Server = wsServer
	resp.Data = loginResp
	return
}

func (s *authService) UserOnline(onlineMsg *pb_mq.UserOnline, resp *xhttp.Resp) (ok bool) {
	var (
		err error
	)
	_, _, err = s.producer.EnQueue(onlineMsg, "")
	if err != nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_MESSAGE_ENQUEUE_FAILED, xhttp.ERROR_HTTP_MESSAGE_ENQUEUE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_MESSAGE_ENQUEUE_FAILED, xhttp.ERROR_HTTP_MESSAGE_ENQUEUE_FAILED, err.Error())
	}
	ok = true
	return
}
