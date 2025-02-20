package auth

import (
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	"context"

	"byte_go/backend/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
	DeleteTokenByRPC(ctx context.Context, Req *auth.DeleteTokenReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	DeleteTokenListByRPC(ctx context.Context, Req *auth.DeleteTokenListReq, callOptions ...callopt.Option) (r *common.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.DeliverTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	return c.kitexClient.VerifyTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteTokenByRPC(ctx context.Context, Req *auth.DeleteTokenReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.DeleteTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteTokenListByRPC(ctx context.Context, Req *auth.DeleteTokenListReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.DeleteTokenListByRPC(ctx, Req, callOptions...)
}
