package rpc

import (
	"byte_go/backend/app/front/conf"
	"byte_go/backend/rpc_gen/kitex_gen/auth/authservice"
	"byte_go/backend/rpc_gen/kitex_gen/cart/cartservice"
	"byte_go/backend/rpc_gen/kitex_gen/checkout/checkoutservice"
	"byte_go/backend/rpc_gen/kitex_gen/order/orderservice"
	"byte_go/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"byte_go/backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"byte_go/backend/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

/**
 * @author: 锈渎
 * @date: 2025/2/4 17:20
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var (
	UserClient     userservice.Client
	AuthClient     authservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	OrderClient    orderservice.Client
	PaymentClient  paymentservice.Client
	CheckoutClient checkoutservice.Client
	Once           sync.Once
	err            error
)

func Init() {
	Once.Do(func() {
		var r discovery.Resolver
		r, err = consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
		if err != nil {
			hlog.Fatal(err)
		}
		initUserClient(r)
		initAuthClient(r)
		initProductCatalogClient(r)
		initCartClient(r)
		initOrderClient(r)
		initPaymentClient(r)
		initCheckoutClient(r)
	})
}

func initUserClient(r discovery.Resolver) {
	UserClient, err = userservice.NewClient(
		"user",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initAuthClient(r discovery.Resolver) {
	AuthClient, err = authservice.NewClient(
		"auth",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initProductCatalogClient(r discovery.Resolver) {
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient(r discovery.Resolver) {
	CartClient, err = cartservice.NewClient(
		"cart",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		hlog.Fatal(err)
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
		hlog.Fatal(err)
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
		hlog.Fatal(err)
	}
}

func initCheckoutClient(r discovery.Resolver) {
	CheckoutClient, err = checkoutservice.NewClient(
		"checkout",
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}
