package rpc

import (
	"byte_go/backend/app/frontend/conf"
	"byte_go/backend/rpc_gen/kitex_gen/auth/authservice"
	"byte_go/backend/rpc_gen/kitex_gen/cart/cartservice"
	"byte_go/backend/rpc_gen/kitex_gen/checkout/checkoutservice"
	"byte_go/backend/rpc_gen/kitex_gen/order/orderservice"
	"byte_go/backend/rpc_gen/kitex_gen/payment/paymentservice"
	"byte_go/backend/rpc_gen/kitex_gen/product/productcatalogservice"
	"byte_go/backend/rpc_gen/kitex_gen/user/userservice"
	"byte_go/backend/utils/clientsuite"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
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
	ServiceName    = conf.GetConf().Hertz.Service
	MetricsPort    = conf.GetConf().Hertz.MetricsPort
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
)

func Init() {
	Once.Do(func() {
		initUserClient()
		initAuthClient()
		initProductCatalogClient()
		initCartClient()
		initOrderClient()
		initPaymentClient()
		initCheckoutClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient(
		"user",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient(
		"auth",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initProductCatalogClient() {
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient(
		"cart",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient(
		"order",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient(
		"payment",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient(
		"checkout",
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	)
	if err != nil {
		hlog.Fatal(err)
	}
}
