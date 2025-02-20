package service

import (
	"byte_go/backend/app/checkout/infra/rpc"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	checkout "byte_go/backend/rpc_gen/kitex_gen/checkout"
	"byte_go/backend/rpc_gen/kitex_gen/order"
	"byte_go/backend/rpc_gen/kitex_gen/payment"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
	"byte_go/kitex_err"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	// 获取购物车
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &rpcCart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if cartResult == nil || cartResult.Cart == nil || cartResult.Cart.Items == nil || len(cartResult.Cart.Items) == 0 {
		klog.Error("cart is empty")
		return nil, kitex_err.CartEmptyError
	}

	// 获取商品信息
	productIds := make([]uint32, len(cartResult.Cart.Items))
	for idx, item := range cartResult.Cart.Items {
		productIds[idx] = item.ProductId
	}
	products, err := rpc.ProductClient.ListProductByIds(s.ctx, &rpcProduct.ListProductByIdsReq{
		ProductIds: productIds,
	})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	// 构建订单
	var productMap = make(map[uint32]*rpcProduct.Product)
	for _, p := range products.Products {
		productMap[p.ProductId] = p
	}
	var totalPrice float32
	oi := make([]*order.OrderItem, len(cartResult.Cart.Items))
	for i, item := range cartResult.Cart.Items {
		price := productMap[item.ProductId].Price * float32(item.Quantity)
		totalPrice += price
		oi[i] = &order.OrderItem{
			Item: &rpcCart.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: price,
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Email:        req.Email,
		OrderItems:   oi,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Nickname: req.Firstname + req.Lastname,
	})

	if err != nil {
		klog.Error(err)
		return nil, err
	}

	// 清空购物车
	var orderId string
	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &rpcCart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	// 支付
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		Amount:  totalPrice,
		OrderId: orderId,
		UserId:  req.UserId,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	// 返回
	return &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}, nil
}
