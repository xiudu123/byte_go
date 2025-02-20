package service

import (
	"context"
	"testing"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
	user "byte_go/backend/rpc_gen/kitex_gen/user"
)

func TestDeleteUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	req := &user.DeleteUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
