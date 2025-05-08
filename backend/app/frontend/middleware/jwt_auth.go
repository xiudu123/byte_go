package middleware

import (
	"byte_go/backend/app/frontend/casbin"
	"byte_go/backend/app/frontend/infra/rpc"
	rpcAuth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	"/ping",          // 测试接口
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
		userId := ctx.GetUint32("userId")
		// 验证token
		claims, err := rpc.AuthClient.VerifyTokenByRPC(c, &rpcAuth.VerifyTokenReq{Token: authHeader, UserId: userId})
		if err != nil {
			ctx.AbortWithStatusJSON(401, "请重新登录")
			return
		}

		//claimsJson, _ := json.Marshal(claims)

		ok, err := casbin.CheckPermission(claims.UserId, currentPath, string(ctx.Method()))
		if err != nil {
			hlog.Errorf("user [%d] visit [%s] method [%s] check permission failed: %v", claims.UserId, currentPath, string(ctx.Method()), err.Error())
			ctx.AbortWithStatusJSON(401, "请重新登录")
			return
		}
		if !ok {
			hlog.Errorf("user [%d] visit [%s] method [%s] no permission", claims.UserId, currentPath, string(ctx.Method()))
			ctx.AbortWithStatusJSON(401, "无权限")
			return
		}

		ctx.Set("jti", claims.Jti)
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
