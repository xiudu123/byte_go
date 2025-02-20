package user

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {

	userResp, err := rpc.UserClient.Register(h.Context, &rpcUser.RegisterReq{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})

	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &user.RegisterResp{
		User: &user.User{
			UserId:    userResp.User.UserId,
			Email:     userResp.User.Email,
			Username:  userResp.User.Username,
			AvatarUrl: userResp.User.AvatarUrl,
		},
		AccessToken: userResp.Token,
	}, nil
}
