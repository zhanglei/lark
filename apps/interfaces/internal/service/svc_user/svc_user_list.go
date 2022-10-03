package svc_user

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_user"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_user"
	"lark/pkg/xhttp"
)

func (s *userService) UserList(params *dto_user.UserListReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req             = &pb_user.GetUserListReq{UidList: params.UidList}
		getUserListResp *pb_user.GetUserListResp
		list            = make([]dto_user.UserInfo, 0)
	)
	getUserListResp = s.userClient.GetUserList(req)
	if getUserListResp == nil {
		resp.SetRespInfo(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if getUserListResp.Code > 0 {
		resp.SetRespInfo(getUserListResp.Code, getUserListResp.Msg)
		return
	}
	copier.Copy(list, getUserListResp.List)
	resp.Data = list
	return
}
