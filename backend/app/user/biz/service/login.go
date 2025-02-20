package service

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/model"
	"byte_go/backend/app/user/infra/rpc"
	"byte_go/backend/rpc_gen/kitex_gen/auth"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
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
		return nil, kitex_err.ValidateError
	}

	// 从数据库中获取用户信息
	userRow, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kitex_err.UserNotExist
		}
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, kitex_err.UserPasswordError
	}

	// 生成token
	tokenResp, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: uint32(userRow.ID)})

	if err != nil {
		klog.Error(strconv.Itoa(int(userRow.ID)) + " " + err.Error())
		return nil, err
	}

	// 返回结果
	return &user.LoginResp{
		User: &user.User{
			UserId:    uint32(userRow.ID),
			Username:  userRow.Username,
			Email:     userRow.Email,
			AvatarUrl: userRow.Avatar,
		},
		Token: tokenResp.Token,
	}, nil
}
