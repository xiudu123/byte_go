package service

import (
	auth "byte_go/backend/rpc_gen/kitex_gen/auth"
	"context"
	"testing"
)

func TestDeleteTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeleteTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
