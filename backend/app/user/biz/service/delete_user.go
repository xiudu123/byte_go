package service

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/model"
	"byte_go/backend/app/user/infra/rpc"
	rpcAuth "byte_go/backend/rpc_gen/kitex_gen/auth"
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
	// Finish your business logic.
	userId := req.UserId
	err = model.DeleteById(mysql.DB, uint(userId))
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}
	_, _ = rpc.AuthClient.DeleteTokenListByRPC(s.ctx, &rpcAuth.DeleteTokenListReq{UserId: userId})
	return
}
