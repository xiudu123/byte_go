package service

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/model"
	"byte_go/backend/app/user/infra/rpc"
	"byte_go/backend/rpc_gen/kitex_gen/auth"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {

	// 参数校验
	if req.Email == "" || req.Password == "" {
		return nil, kitex_err.RequestParamError
	}

	// 从数据库中获取用户信息
	userInfo, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		klog.Errorf("login get user by email of %s, err: %v", req.Email, err.Error())
		return nil, kitex_err.ValidateLoginError
	}

	// 校验密码
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.PasswordHash), []byte(req.Password)); err != nil {
		return nil, kitex_err.UserPasswordError
	}

	// 生成token
	tokenResp, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: uint32(userInfo.ID)})

	if err != nil {
		klog.Errorf("create token by rpc failed,user_id:%v,  err: %v", userInfo.ID, err.Error())
		return nil, err
	}

	// 返回结果
	return &user.LoginResp{
		User: &user.User{
			UserId:    uint32(userInfo.ID),
			Username:  userInfo.Username,
			Email:     userInfo.Email,
			AvatarUrl: userInfo.Avatar,
		},
		Token: tokenResp.Token,
	}, nil
}
