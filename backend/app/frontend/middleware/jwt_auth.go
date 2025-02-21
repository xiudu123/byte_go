package middleware

import (
	"byte_go/backend/app/front/infra/rpc"
	"byte_go/backend/constants"
	rpcAuth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"context"
	"encoding/json"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"regexp"
	"strings"
)

/**
 * @author: 锈渎
 * @date: 2025/2/12 16:23
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

var whitelist = []string{
	"/user/login",    // 登录
	"/user/register", // 注册
	"/user/get/*",    // 获取用户信息
}

func JwtAuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		// 检测是否在白名单中
		currentPath := string(ctx.Path())
		if isWhiteListed(currentPath) {
			ctx.Next(c)
			return
		}

		// 从header中获取token
		authHeader := string(ctx.GetHeader("Authorization"))
		// 验证token
		claims, _ := rpc.AuthClient.VerifyTokenByRPC(c, &rpcAuth.VerifyTokenReq{Token: authHeader})
		//if err != nil {
		//	ctx.AbortWithStatusJSON(401, "请重新登录")
		//	return
		//}

		claimsJson, _ := json.Marshal(claims)

		c = metainfo.WithValue(c, constants.JwtClaimsKey, string(claimsJson))
		ctx.Next(c)
	}
}

func isWhiteListed(path string) bool {
	for _, pattern := range whitelist {
		// 转换通配符为正则表达式
		regexPattern := "^" + strings.ReplaceAll(pattern, "*", ".*") + "$"

		matched, _ := regexp.MatchString(regexPattern, path)
		if matched {
			return true
		}
	}
	return false
}
