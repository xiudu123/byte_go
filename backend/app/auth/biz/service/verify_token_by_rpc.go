package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"byte_go/backend/utils"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	claims, err := utils.ParseToken(req.Token)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.TokenValidError
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, kitex_err.TokenExpiredError
	}
	// 验证 jti 是否在黑名单中
	exist, err := redis.CheckJITIsBlackListed(s.ctx, claims.JTI)
	if err != nil {
		klog.Error(err)
		return nil, kitex_err.SystemError
	}
	if exist {
		return nil, kitex_err.TokenValidError
	}

	return &auth.VerifyResp{
		UserId: claims.UserId,
		Jti:    claims.JTI,
	}, nil
}
