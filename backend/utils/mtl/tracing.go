package mtl

/**
 * @author: 锈渎
 * @date: 2025/2/26 19:06
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p
}
