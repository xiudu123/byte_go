package user

import (
	user "byte_go/backend/app/front/hertz_gen/frontend/user"
	"byte_go/backend/app/front/infra/rpc"
	rpcUser "byte_go/backend/rpc_gen/kitex_gen/user"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	if req.UserId <= 0 && req.Email == "" {
		return nil, errors.New("参数错误")
	}

	var userResp *rpcUser.GetUserInfoResp
	if req.Email != "" {

		userResp, err = rpc.UserClient.GetUserInfo(h.Context, &rpcUser.GetUserInfoReq{
			Identifier: &rpcUser.GetUserInfoReq_Email{
				Email: req.Email,
			},
		})
	} else {
		userResp, err = rpc.UserClient.GetUserInfo(h.Context, &rpcUser.GetUserInfoReq{
			Identifier: &rpcUser.GetUserInfoReq_UserId{
				UserId: req.UserId,
			},
		})
	}

	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResp{
		User: &user.User{
			UserId:    userResp.User.UserId,
			Email:     userResp.User.Email,
			Username:  userResp.User.Username,
			AvatarUrl: userResp.User.AvatarUrl,
		},
	}, nil
}
