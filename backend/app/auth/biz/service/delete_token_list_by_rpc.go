package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"context"
)

type DeleteTokenListByRPCService struct {
	ctx context.Context
} // NewDeleteTokenListByRPCService new DeleteTokenListByRPCService
func NewDeleteTokenListByRPCService(ctx context.Context) *DeleteTokenListByRPCService {
	return &DeleteTokenListByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeleteTokenListByRPCService) Run(req *auth.DeleteTokenListReq) (resp *common.Empty, err error) {
	// Finish your business logic.
	jtiList, _ := redis.ListJTIList(s.ctx, req.UserId)
	redis.AddJTIList2BlackListed(s.ctx, jtiList)
	return
}
