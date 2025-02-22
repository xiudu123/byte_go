package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteTokenByRPCService struct {
	ctx context.Context
} // NewDeleteTokenByRPCService new DeleteTokenByRPCService
func NewDeleteTokenByRPCService(ctx context.Context) *DeleteTokenByRPCService {
	return &DeleteTokenByRPCService{ctx: ctx}
}

// Run 执行令牌吊销服务
// 通过jti将其加入黑名单使其对应令牌立即失效
// @param req 包含jti的请求
// @return 空响应
// @error jti为空/redis错误
func (s *DeleteTokenByRPCService) Run(req *auth.DeleteTokenReq) (resp *common.Empty, err error) {
	// 校验参数
	if req == nil || req.Jti == "" {
		return nil, kitex_err.JTIEmptyError
	}
	// 加入黑名单
	if err = redis.AddJTI2BlackListed(s.ctx, req.Jti); err != nil {
		klog.Errorf("redis blacklist add failed for jti:%s, error:%v", req.Jti, err)
		return nil, kitex_err.RedisError
	}
	return &common.Empty{}, nil
}
