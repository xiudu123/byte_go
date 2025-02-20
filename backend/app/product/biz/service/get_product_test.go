package service

import (
	"byte_go/backend/rpc_gen/kitex_gen/product"
	"context"
	"fmt"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	fmt.Println("!!!!")
	//fmt.Println(s.ctx)
	// init req and assert value

	req := &product.GetProductReq{ProductId: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

	fmt.Println(resp)

}
