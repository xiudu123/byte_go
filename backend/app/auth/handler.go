package main

import (
	"byte_go/backend/app/auth/biz/service"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"byte_go/backend/rpc_gen/kitex_gen/common"
	"context"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// DeleteTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeleteTokenByRPC(ctx context.Context, req *auth.DeleteTokenReq) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteTokenByRPCService(ctx).Run(req)

	return resp, err
}
