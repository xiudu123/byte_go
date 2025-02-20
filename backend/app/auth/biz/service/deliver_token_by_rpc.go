package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"byte_go/backend/utils"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	token, jti, err := utils.GenerateToken(req.UserId)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.TokenCreateError
	}

	redis.SetJTI(s.ctx, req.UserId, jti)

	return &auth.DeliveryResp{
		Token: token,
	}, nil
}
