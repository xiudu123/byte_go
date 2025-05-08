package service

import (
	"byte_go/backend/app/user/biz/dal/repository"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *common.Empty, err error) {
	// 参数校验
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	// 删除用户
	userId := req.UserId
	userRepository := repository.NewUserRepository(s.ctx)
	if err = userRepository.DeleteById(uint(userId)); err != nil {
		klog.Errorf("user delete failed, user_id:%+v,  err: %v", userId, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 删除用户的token
	// TODO 修改用户权限版本号
	//_, err = rpc.AuthClient.DeleteTokenListByRPC(s.ctx, &rpcAuth.DeleteTokenListReq{UserId: userId})
	//if err != nil {
	//	klog.Errorf("user delete token list failed, user_id:%+v,  err: %v", userId, err.Error())
	//	return nil, err
	//}

	return &common.Empty{}, nil
}
