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
	"strconv"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
// TODO: 生成 UserId
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 参数校验
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, kitex_err.ValidateError
	}
	if req.Password != req.ConfirmPassword {
		return nil, kitex_err.ValidatePasswordNotEqual
	}

	// 哈希密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 将注册用户写入数据库
	newUser := model.User{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Username:     req.Username,
		Avatar:       "https://p3-passport.byteacctimg.com/img/mosaic-legacy/3796/2975850990~50x50.awebp",
	}
	err = model.Create(mysql.DB, &newUser)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 生成token
	tokenResp, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: uint32(newUser.ID)})

	if err != nil {
		klog.Error(strconv.Itoa(int(newUser.ID)) + " " + err.Error())
		return nil, err
	}

	// 返回结果
	return &user.RegisterResp{
		User: &user.User{
			UserId:    uint32(newUser.ID),
			Username:  newUser.Username,
			Email:     newUser.Email,
			AvatarUrl: newUser.Avatar,
		},
		Token: tokenResp.Token,
	}, nil
}
