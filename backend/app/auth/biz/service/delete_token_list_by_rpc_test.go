package service

import (
	"context"
	"testing"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
)

func TestDeleteTokenListByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteTokenListByRPCService(ctx)
	// init req and assert value

	req := &auth.DeleteTokenListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
