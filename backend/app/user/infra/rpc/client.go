package rpc

import (
	"byte_go/backend/app/user/conf"
	"byte_go/backend/rpc_gen/kitex_gen/auth/authservice"
	"byte_go/backend/utils/clientsuite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

/**
 * @author: 锈渎
 * @date: 2025/2/11 22:16
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var (
	AuthClient   authservice.Client
	once         sync.Once
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		klog.Fatalf("init auth client failed: %v", err)
	}
}
