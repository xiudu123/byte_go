package rpc

import (
	"byte_go/backend/app/checkout/conf"
	"byte_go/backend/rpc_gen/kitex_gen/cart/cartservice"
	"byte_go/backend/rpc_gen/kitex_gen/order/orderservice"
	"byte_go/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"byte_go/backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"byte_go/backend/utils/clientsuite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

/**
 * @author: 锈渎
 * @date: 2025/2/19 21:02
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func InitClient() {
	once.Do(func() {
		opts := []client.Option{
			client.WithSuite(clientsuite.CommonClientSuite{
				CurrentServiceName: ServiceName,
				RegistryAddr:       RegistryAddr,
			}),
		}
		initCartClient(opts)
		initProductClient(opts)
		initOrderClient(opts)
		initPaymentClient(opts)
	})
}

func initCartClient(opts []client.Option) {
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		klog.Fatalf("init cart client failed: %v", err)
	}
}

func initProductClient(opts []client.Option) {
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatalf("init product client failed: %v", err)
	}
}

func initOrderClient(opts []client.Option) {
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Fatalf("init order client failed: %v", err)
	}
}

func initPaymentClient(opts []client.Option) {
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		klog.Fatalf("init payment client failed: %v", err)
	}
}
