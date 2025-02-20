package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"context"
)

type DeleteTokenByRPCService struct {
	ctx context.Context
} // NewDeleteTokenByRPCService new DeleteTokenByRPCService
func NewDeleteTokenByRPCService(ctx context.Context) *DeleteTokenByRPCService {
	return &DeleteTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeleteTokenByRPCService) Run(req *auth.DeleteTokenReq) (resp *common.Empty, err error) {
	// Finish your business logic.
	redis.AddJTI2BlackListed(s.ctx, req.Jti)
	return
}
