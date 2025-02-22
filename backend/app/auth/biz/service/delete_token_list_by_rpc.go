package service

import (
	"byte_go/backend/app/auth/biz/dal/redis"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteTokenListByRPCService struct {
	ctx context.Context
} // NewDeleteTokenListByRPCService new DeleteTokenListByRPCService
func NewDeleteTokenListByRPCService(ctx context.Context) *DeleteTokenListByRPCService {
	return &DeleteTokenListByRPCService{ctx: ctx}
}

// Run 执行令牌批量吊销服务
// 通过用户的id获取其所有的jti并将其加入黑名单使其对应令牌立即失效
// @param req 包含用户id的请求
// @return 空响应
// @error 参数错误/redis错误
func (s *DeleteTokenListByRPCService) Run(req *auth.DeleteTokenListReq) (resp *common.Empty, err error) {

	// 校验参数
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	// 获取用户的 JTI 列表
	jtiList, err := redis.ListJTIList(s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("redis get jti list failed for user_id:%d, error:%v", req.UserId, err)
		return nil, kitex_err.RedisError
	}

	// 加入黑名单
	if err = redis.AddJTIList2BlackListed(s.ctx, jtiList); err != nil {
		klog.Errorf("redis blacklist add failed for user_id:%d, error:%v", req.UserId, err)
		return nil, kitex_err.RedisError
	}

	// 返回
	return &common.Empty{}, nil
}
