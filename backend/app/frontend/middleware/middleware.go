package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

/**
 * @author: 锈渎
 * @date: 2025/2/12 16:37
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

func Register(h *server.Hertz) {
	h.Use(JwtAuthMiddleware())
}
