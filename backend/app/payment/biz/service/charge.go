package service

import (
	"byte_go/backend/app/payment/biz/dal/mysql"
	"byte_go/backend/app/payment/biz/model"
	payment "byte_go/backend/rpc_gen/kitex_gen/payment"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"strconv"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {

	// 校验参数
	if req == nil {
		return nil, kitex_err.RequestParamError
	}

	// 校验信用卡
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		klog.Errorf("card valid error: %v", err.Error())
		return nil, kitex_err.CardValidError
	}

	// 生成 交易id
	var transactionId uuid.UUID
	transactionId, err = uuid.NewRandom()
	if err != nil {
		klog.Errorf("id generate error: %v", err.Error())
		return nil, kitex_err.IdGenerateError
	}

	// 写入数据库
	paymentQuery := model.NewPaymentQuery(s.ctx, mysql.DB)
	err = paymentQuery.CreatePaymentLog(&model.PaymentLog{
		UserId:        int32(req.UserId),
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
	})
	if err != nil {
		klog.Errorf("用户 %v 使用卡号 %v 支付失败, error: %v", req.UserId, req.CreditCard.CreditCardNumber, err.Error())
		return nil, kitex_err.MysqlError
	}

	// 返回
	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
