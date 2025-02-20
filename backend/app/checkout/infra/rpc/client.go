package rpc

import (
	"byte_go/backend/app/checkout/conf"
	"byte_go/backend/rpc_gen/kitex_gen/cart/cartservice"
	"byte_go/backend/rpc_gen/kitex_gen/order/orderservice"
	"byte_go/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"byte_go/backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
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
)

func InitClient() {
	once.Do(func() {
		var r discovery.Resolver
		r, err = consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
		if err != nil {
			klog.Fatal(err)
		}
		initCartClient(r)
		initProductClient(r)
		initOrderClient(r)
		initPaymentClient(r)
	})
}

func initCartClient(r discovery.Resolver) {
	CartClient, err = cartservice.NewClient(
		"cart",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		klog.Fatal(err)
	}
}

func initProductClient(r discovery.Resolver) {
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		klog.Fatal(err)
	}
}

func initOrderClient(r discovery.Resolver) {
	OrderClient, err = orderservice.NewClient(
		"order",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		klog.Fatal(err)
	}
}

func initPaymentClient(r discovery.Resolver) {
	PaymentClient, err = paymentservice.NewClient(
		"payment",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		klog.Fatal(err)
	}
}
