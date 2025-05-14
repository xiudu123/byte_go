package service

import (
	"context"
	"testing"
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	common "byte_go/backend/rpc_gen/kitex_gen/common"
)

func TestClearUserTokensByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewClearUserTokensByRPCService(ctx)
	// init req and assert value

	req := &auth.ClearUserTokensReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
