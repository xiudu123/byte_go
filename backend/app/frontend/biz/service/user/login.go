package user

import (
	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"byte_go/backend/app/front/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {

	userResp, err := rpc.UserClient.Login(h.Context, &rpcUser.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		klog.CtxErrorf(h.Context, "login user [%s] failed, err: %v", req.Email, err.Error())
		return nil, err
	}

	return &user.LoginResp{
		User: &user.User{
			UserId:    userResp.User.UserId,
			Email:     userResp.User.Email,
			Username:  userResp.User.Username,
			AvatarUrl: userResp.User.AvatarUrl,
		},
		AccessToken: userResp.Token,
	}, nil
}
