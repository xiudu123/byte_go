package auth

import (
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteTokenByRPC(ctx context.Context, req *auth.DeleteTokenReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
