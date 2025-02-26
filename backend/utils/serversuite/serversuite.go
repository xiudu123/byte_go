package serversuite

import (
	"byte_go/backend/utils/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

/**
 * @author: 锈渎
 * @date: 2025/2/26 15:57
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

type CommonServerSuite struct {
	CurrentServerName string
	RegistryAddr      string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServerName,
		}),
		server.WithTracer(prometheus.NewServerTracer("",
			"",
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry))),
		server.WithSuite(tracing.NewServerSuite()),
	}

	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r), server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	return opts
}
