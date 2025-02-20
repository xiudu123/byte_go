package kitex_err

import "github.com/cloudwego/kitex/pkg/kerrors"

/**
 * @author: 锈渎
 * @date: 2025/2/11 21:27
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 枚举 kitex 错误
 */
var (
	ValidateError            = kerrors.NewBizStatusError(10001, "参数校验错误")
	ValidatePasswordNotEqual = kerrors.NewBizStatusError(10002, "两次密码不一致")

	UserNotExist      = kerrors.NewBizStatusError(20001, "用户不存在")
	UserPasswordError = kerrors.NewBizStatusError(20002, "密码错误")
	TokenCreateError  = kerrors.NewBizStatusError(20003, "token 生成错误")
	TokenValidError   = kerrors.NewBizStatusError(20004, "token 验证错误")
	TokenExpiredError = kerrors.NewBizStatusError(20005, "token 过期")
	PermissionError   = kerrors.NewBizStatusError(20006, "权限不足")

	ProductNotExist   = kerrors.NewBizStatusError(30001, "商品不存在")
	OrderItemEmpty    = kerrors.NewBizStatusError(30002, "订单商品为空")
	CardValidError    = kerrors.NewBizStatusError(30003, "信用卡验证错误")
	CartEmptyError    = kerrors.NewBizStatusError(30004, "购物车为空")
	ProductEmptyError = kerrors.NewBizStatusError(30005, "商品为空")

	SystemError = kerrors.NewBizStatusError(50001, "系统错误")
)
