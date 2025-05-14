package service

import (
	"byte_go/backend/app/auth/biz/dal/repository"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ClearUserTokensByRPCService struct {
	ctx context.Context
} // NewClearUserTokensByRPCService new ClearUserTokensByRPCService
func NewClearUserTokensByRPCService(ctx context.Context) *ClearUserTokensByRPCService {
	return &ClearUserTokensByRPCService{ctx: ctx}
}

// Run create note info
func (s *ClearUserTokensByRPCService) Run(req *auth.ClearUserTokensReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	authRepo := repository.NewAuthRepository(s.ctx)
	if err := authRepo.IncrementPermissionVersion(s.ctx, req.UserId); err != nil {
		klog.Errorf("redis permission version increment failed for userId:%d, error:%v", req.UserId, err)
		return nil, kitex_err.RedisError
	}

	return &common.Empty{}, nil
}
