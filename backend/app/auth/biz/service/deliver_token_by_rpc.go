package service

import (
	"byte_go/backend/app/auth/biz/dal/repository"
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

// Run 实现令牌颁发服务
// 通过用户的id生成一个token并将其存入redis
// @param req 包含用户id的请求
// @return 包含token的响应
// @error 参数错误/令牌生成错误/redis错误
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	// 校验参数
	if req == nil || req.UserId == 0 {
		return nil, kitex_err.RequestParamError
	}

	authRepo := repository.NewAuthRepository(s.ctx)

	// 生成 token
	permissionVersion, err := authRepo.GetPermissionVersion(s.ctx, req.UserId)
	if err != nil {
		klog.Error("get permission version failed: ", err.Error())
		return nil, kitex_err.RedisError
	}
	token, _, err := utils.GenerateToken(req.UserId, permissionVersion)
	if err != nil {
		klog.Error("token generate failed: ", err.Error())
		return nil, kitex_err.TokenCreateError
	}

	// 返回
	return &auth.DeliveryResp{
		Token: token,
	}, nil
}
