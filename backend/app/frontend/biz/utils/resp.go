package utils

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	er := strings.TrimPrefix(err.Error(), "biz error: ")
	parts := strings.Split(er, ", ")
	var status int
	var msg string
	status, _ = strconv.Atoi(strings.TrimPrefix(parts[0], "code="))
	msg = strings.TrimPrefix(parts[1], "msg=")

	result := NewResult(status, msg, nil)
	c.JSON(code, result)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, SuccessWithData(data))
}
