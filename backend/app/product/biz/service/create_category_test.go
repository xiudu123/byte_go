package service

import (
	"context"
	"testing"
	product "byte_go/backend/rpc_gen/kitex_gen/product"
)

func TestCreateCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateCategoryService(ctx)
	// init req and assert value

	req := &product.CreateCategoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
