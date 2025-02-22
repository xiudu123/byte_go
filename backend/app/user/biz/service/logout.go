package service

import (
	"byte_go/backend/app/user/infra/rpc"
	"byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"byte_go/backend/utils"
	"byte_go/kitex_err"
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

	// 从token中获取jti
	claims, err := utils.GetClaims(s.ctx)
	if err != nil || claims == nil || claims.JTI == "" {
		return nil, kitex_err.InvalidAuthError
	}

	// 将jti加入黑名单
	_, err = rpc.AuthClient.DeleteTokenByRPC(s.ctx, &auth.DeleteTokenReq{Jti: claims.JTI})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &common.Empty{}, nil
}
