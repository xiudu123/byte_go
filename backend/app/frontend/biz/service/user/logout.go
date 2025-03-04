package user

import (
	"byte_go/backend/app/front/hertz_gen/frontend/common_hertz"
	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"byte_go/backend/app/front/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *user.LogoutReq) (resp *common_hertz.Empty, err error) {
	jti, _ := h.RequestContext.Get("jti")
	_, err = rpc.UserClient.Logout(h.Context, &rpcUser.LogoutReq{
		Jti: jti.(string),
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "logout failed, err: %v", err.Error())
		return nil, err
	}
	return nil, nil
}
