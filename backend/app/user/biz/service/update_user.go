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

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {

	// 参数校验
	if req == nil {
		return nil, kitex_err.RequestParamError
	}

	// 检查用户是否存在
	userInfo, err := model.GetById(mysql.DB, uint(req.UserId))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kitex_err.UserNotExist
	}
	if err != nil {
		klog.Errorf("user get info failed, param:%+v,  err: %v", req.UserId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 检查参数
	var updateInfo = make(map[string]interface{})
	if req.Username != "" {
		updateInfo["username"] = req.Username
		userInfo.Username = req.Username
	}
	if req.AvatarUrl != "" {
		updateInfo["avatar"] = req.AvatarUrl
		userInfo.Avatar = req.AvatarUrl
	}

	// 更新
	if err = model.UpdateById(mysql.DB, uint(req.UserId), updateInfo); err != nil {
		klog.Errorf("user update failed, param:%+v,  err: %v", req.UserId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回结果
	return &user.UpdateUserResp{
		User: &user.User{
			UserId:    uint32(userInfo.ID),
			Username:  userInfo.Username,
			Email:     userInfo.Email,
			AvatarUrl: userInfo.Avatar,
		},
	}, nil
}
