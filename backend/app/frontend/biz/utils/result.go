package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/2/4 22:18
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 封装统一接口返回值
 */

// Result 封装统一接口返回值
type Result struct {
	Timestamp int64       `json:"timestamp"` // 响应时间(当前毫秒数)
	Code      int         `json:"code"`      // 返回状态码
	Msg       string      `json:"msg"`       // 返回错误信息
	Data      interface{} `json:"data"`      // 返回数据类型
}

// NewResult 创建一个 Result 实例
func NewResult(status int, errorMessage string, data interface{}) *Result {
	return &Result{
		Timestamp: time.Now().UnixMilli(),
		Code:      status,
		Msg:       errorMessage,
		Data:      data,
	}
}

func (e *Result) Error() string {
	JsonErr, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("code: %d, msg: %s, timestamp: %d", e.Code, e.Msg, e.Timestamp)
	}
	return string(JsonErr)
}

// Success 无数据返回的成功结果
func Success() *Result {
	return NewResult(200, "success", nil)
}

// SuccessWithData 带数据返回的成功结果
func SuccessWithData(data interface{}) *Result {
	return NewResult(200, "success", data)
}

// Error 返回错误结果
func Error(message string) *Result {
	return NewResult(500, message, nil)
}

// ErrorWithCode 返回带状态码的错误结果
func ErrorWithCode(code int, message string) *Result {
	return NewResult(code, message, nil)
}

// ErrorWithData 返回带数据和状态码的错误结果
func ErrorWithData(code int, message string, data interface{}) *Result {
	return NewResult(code, message, data)
}

// ErrorDefault 返回默认错误结果
func ErrorDefault() *Result {
	return NewResult(500, "系统错误", nil)
}

// ErrorFromEnum 返回枚举定义的错误结果
func ErrorFromEnum(resultStatusEnum ResultStatusEnum) *Result {
	return NewResult(resultStatusEnum.Status, resultStatusEnum.Message, nil)
}

// ErrorFromEnumWithData 返回枚举定义的错误结果并带数据
func ErrorFromEnumWithData(resultStatusEnum ResultStatusEnum, data interface{}) *Result {
	return NewResult(resultStatusEnum.Status, resultStatusEnum.Message, data)
}

// ResultStatusEnum 错误状态枚举
type ResultStatusEnum struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
