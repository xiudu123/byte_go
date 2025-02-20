package service

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/model"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
	"byte_go/backend/utils"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {

	// 检查用户是否存在
	userUp, err := model.GetById(mysql.DB, uint(req.UserId))
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.UserNotExist
	}

	// 检查用户是否有权限
	// 检查是否水平越权
	claims, _ := utils.GetClaims(s.ctx)
	if claims.UserId != req.UserId {
		return nil, kitex_err.PermissionError
	}

	// 检查参数
	var updateInfo = make(map[string]interface{})
	if req.Username != "" {
		updateInfo["username"] = req.Username
		userUp.Username = req.Username
	}
	if req.AvatarUrl != "" {
		updateInfo["avatar"] = req.AvatarUrl
		userUp.Avatar = req.AvatarUrl
	}

	// 更新
	err = model.UpdateById(mysql.DB, uint(req.UserId), updateInfo)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}

	// 返回结果
	return &user.UpdateUserResp{
		User: &user.User{
			UserId:    uint32(userUp.ID),
			Username:  userUp.Username,
			Email:     userUp.Email,
			AvatarUrl: userUp.Avatar,
		},
	}, nil
}
