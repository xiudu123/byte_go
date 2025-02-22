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

// Run 实现令牌验证服务
// 通过token里面的jti判断用户是否在黑名单中
// 如果在黑名单中则返回验证错误
// 如果不在黑名单中则返回用户id和jti
// @param req 包含token的请求
// @return 包含用户id和jti的响应
// @error token为空/解析错误/过期/验证错误/redis错误
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {

	// 解析token
	if req == nil || req.Token == "" {
		return nil, kitex_err.TokenEmptyError
	}
	claims, err := utils.ParseToken(req.Token)
	if err != nil {
		klog.Error("token parse failed: ", err.Error())
		return nil, kitex_err.TokenParseError
	}

	// 验证 token 是否过期, 预留30s缓冲期, 以应对服务器时间不同步
	if claims.ExpiresAt.Unix() <= time.Now().Add(30*time.Second).Unix() {
		return nil, kitex_err.TokenExpiredError
	}

	// 验证 jti 是否在黑名单中
	if claims == nil || claims.JTI == "" {
		return nil, kitex_err.TokenValidError
	}
	exist, err := redis.CheckJITIsBlackListed(s.ctx, claims.JTI)
	if err != nil {
		klog.Error("check jti is blacklisted failed: ", err.Error())
		return nil, kitex_err.RedisError
	}
	if exist {
		return nil, kitex_err.TokenValidError
	}

	// 返回
	return &auth.VerifyResp{
		UserId: claims.UserId,
		Jti:    claims.JTI,
	}, nil
}
