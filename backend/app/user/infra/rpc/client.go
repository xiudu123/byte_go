package rpc

import (
	"byte_go/backend/app/user/conf"
	"byte_go/backend/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

/**
 * @author: 锈渎
 * @date: 2025/2/11 22:16
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var (
	AuthClient authservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		panic(err)
	}
}
