package payment

import (
	"byte_go/backend/app/front/infra/rpc"
	rpcPayment "byte_go/backend/rpc_gen/kitex_gen/payment"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	payment "byte_go/backend/app/front/hertz_gen/frontend/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type ChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChargeService(Context context.Context, RequestContext *app.RequestContext) *ChargeService {
	return &ChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {

	chargeResult, err := rpc.PaymentClient.Charge(h.Context, &rpcPayment.ChargeReq{
		Amount:  req.Amount,
		UserId:  req.UserId,
		OrderId: req.OrderId,
		CreditCard: &rpcPayment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &payment.ChargeResp{
		TransactionId: chargeResult.TransactionId,
	}, nil
}
