package service

import (
	"context"
	"testing"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
)

func TestLogout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLogoutService(ctx)
	// init req and assert value

	req := &user.LogoutReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
