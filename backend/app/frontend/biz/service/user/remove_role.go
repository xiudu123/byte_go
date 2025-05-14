package user

import (
	"byte_go/backend/app/frontend/casbin"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	common_hertz "byte_go/backend/app/frontend/hertz_gen/frontend/common_hertz"
	user "byte_go/backend/app/frontend/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RemoveRoleService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRemoveRoleService(Context context.Context, RequestContext *app.RequestContext) *RemoveRoleService {
	return &RemoveRoleService{RequestContext: RequestContext, Context: Context}
}

func (h *RemoveRoleService) Run(req *user.RemoveRoleReq) (resp *common_hertz.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	err = casbin.DeleteRoleForUser(req.UserId, req.Role)
	if err != nil {
		hlog.CtxErrorf(h.Context, "remove role for user [%d] failed, err: %v", req.UserId, err.Error())
		return &common_hertz.Empty{}, err
	}

	return &common_hertz.Empty{}, nil
}
