package service

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/model"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
	"byte_go/kitex_err"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {

	// 参数校验
	if req == nil || req.Identifier == nil {
		return nil, kitex_err.RequestParamError
	}

	// 根据 id 或者 email 获取用户信息
	var userInfo model.User
	switch v := req.Identifier.(type) {
	case *user.GetUserInfoReq_UserId:
		userInfo, err = model.GetById(mysql.DB, uint(v.UserId))
	case *user.GetUserInfoReq_Email:
		userInfo, err = model.GetByEmail(mysql.DB, v.Email)
	default:
		return nil, kitex_err.RequestParamError
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kitex_err.UserNotExist
	}
	if err != nil {
		klog.Errorf("user get info failed, param:%+v,  err: %v", req.Identifier, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回结果
	return &user.GetUserInfoResp{
		User: &user.User{
			UserId:    uint32(userInfo.ID),
			Username:  userInfo.Username,
			Email:     userInfo.Email,
			AvatarUrl: userInfo.Avatar,
		},
	}, nil
}
