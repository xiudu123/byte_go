package service

import (
	"context"
	"testing"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
)

func TestListProductByIds_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductByIdsService(ctx)
	// init req and assert value

	req := &product.ListProductByIdsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
