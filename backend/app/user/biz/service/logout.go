package service

import (
	"byte_go/backend/app/user/infra/rpc"
	"byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"byte_go/backend/utils"
	"context"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	// Finish your business logic.
	claims, _ := utils.GetClaims(s.ctx)
	_, _ = rpc.AuthClient.DeleteTokenByRPC(s.ctx, &auth.DeleteTokenReq{Jti: claims.JTI})
	return
}
