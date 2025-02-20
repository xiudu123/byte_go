package user

import (
	"byte_go/backend/app/front/hertz_gen/frontend/common_hertz"
	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"byte_go/backend/app/front/infra/rpc"
	"byte_go/backend/rpc_gen/kitex_gen/common"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *user.LogoutReq) (resp *common_hertz.Empty, err error) {

	_, err = rpc.UserClient.Logout(h.Context, &common.Empty{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
