package kitex_err

import "github.com/cloudwego/kitex/pkg/kerrors"

/**
 * @author: 锈渎
 * @date: 2025/2/11 21:27
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 枚举 kitex 错误
 */
var (
	RequestParamError = kerrors.NewBizStatusError(10001, "请求参数错误")

	ValidateLoginError       = kerrors.NewBizStatusError(20001, "登录错误")
	ValidatePasswordNotEqual = kerrors.NewBizStatusError(20002, "两次密码不一致")
	EmailExistError          = kerrors.NewBizStatusError(20003, "邮箱已存在")
	UserNotExist             = kerrors.NewBizStatusError(20004, "用户不存在")
	UserPasswordError        = kerrors.NewBizStatusError(20005, "密码错误")
	TokenCreateError         = kerrors.NewBizStatusError(20006, "token 生成错误")
	TokenValidError          = kerrors.NewBizStatusError(20007, "token 验证错误")
	TokenExpiredError        = kerrors.NewBizStatusError(20008, "token 过期")
	TokenParseError          = kerrors.NewBizStatusError(20009, "token 解析错误")
	JTIEmptyError            = kerrors.NewBizStatusError(20010, "jti 为空")
	TokenEmptyError          = kerrors.NewBizStatusError(30011, "token 为空")
	InvalidAuthError         = kerrors.NewBizStatusError(20012, "无效的认证")

	ProductNotExist   = kerrors.NewBizStatusError(30001, "商品不存在")
	CategoryNotExist  = kerrors.NewBizStatusError(40004, "商品分类不存在")
	CategoryExist     = kerrors.NewBizStatusError(40005, "商品分类已存在")
	OrderItemEmpty    = kerrors.NewBizStatusError(30002, "订单商品为空")
	CartEmptyError    = kerrors.NewBizStatusError(30004, "购物车为空")
	ProductEmptyError = kerrors.NewBizStatusError(30005, "商品为空")
	OrderNotExist     = kerrors.NewBizStatusError(40001, "订单不存在")
	OrderPaidError    = kerrors.NewBizStatusError(40002, "订单已支付")
	CardValidError    = kerrors.NewBizStatusError(30003, "信用卡验证错误")

	SystemError     = kerrors.NewBizStatusError(50001, "系统错误")
	RedisError      = kerrors.NewBizStatusError(50002, "系统错误")
	MysqlError      = kerrors.NewBizStatusError(50003, "系统错误")
	IdGenerateError = kerrors.NewBizStatusError(50004, "系统错误")
)
