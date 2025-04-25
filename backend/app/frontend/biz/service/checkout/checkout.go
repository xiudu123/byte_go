package checkout

import (
	checkout "byte_go/backend/app/frontend/hertz_gen/frontend/checkout"
	"byte_go/backend/app/frontend/infra/rpc"
	rpcCheckout "byte_go/backend/rpc_gen/kitex_gen/checkout"
	rpcPayment "byte_go/backend/rpc_gen/kitex_gen/payment"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	checkoutResult, err := rpc.CheckoutClient.Checkout(h.Context, &rpcCheckout.CheckoutReq{
		UserId:       req.UserId,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Email:        req.Email,
		UserCurrency: req.UserCurrency,
		Address: &rpcCheckout.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		CreditCard: &rpcPayment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	})

	if err != nil {
		hlog.CtxErrorf(h.Context, "checkout failed, err: %v", err.Error())
		return nil, err
	}

	return &checkout.CheckoutResp{
		OrderId:       checkoutResult.OrderId,
		TransactionId: checkoutResult.TransactionId,
	}, nil
}
