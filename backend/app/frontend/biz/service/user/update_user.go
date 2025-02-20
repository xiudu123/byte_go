package user

import (
	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"byte_go/backend/app/front/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserService {
	return &UpdateUserService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {

	userResp, err := rpc.UserClient.UpdateUser(h.Context, &rpcUser.UpdateUserReq{
		UserId:    req.UserId,
		Username:  req.Username,
		AvatarUrl: req.AvatarUrl,
	})
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResp{
		User: &user.User{
			UserId:    userResp.User.UserId,
			Email:     userResp.User.Email,
			Username:  userResp.User.Username,
			AvatarUrl: userResp.User.AvatarUrl,
		},
	}, nil
}
