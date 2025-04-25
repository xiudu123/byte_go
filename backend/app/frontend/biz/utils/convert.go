package utils

import (
	hertzCart "byte_go/backend/app/frontend/hertz_gen/frontend/cart"
	hertzOrder "byte_go/backend/app/frontend/hertz_gen/frontend/order"
	hertzProduct "byte_go/backend/app/frontend/hertz_gen/frontend/product"
	rpcCart "byte_go/backend/rpc_gen/kitex_gen/cart"
	rpcOrder "byte_go/backend/rpc_gen/kitex_gen/order"
	rpcProduct "byte_go/backend/rpc_gen/kitex_gen/product"
)

/**
 * @author: 锈渎
 * @date: 2025/2/15 22:54
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description:
 */

func ProductGen2Hertz(productGen *rpcProduct.Product) (productHertz *hertzProduct.Product) {
	productHertz = &hertzProduct.Product{
		ProductId:   productGen.ProductId,
		Name:        productGen.Name,
		Description: productGen.Description,
		Picture:     productGen.Picture,
		Price:       productGen.Price,
		Categories:  productGen.Categories,
	}
	return
}

func CartItemHertz2Gen(cartItemHertz *hertzCart.CartItem) (cartItemGen *rpcCart.CartItem) {

	cartItemGen = &rpcCart.CartItem{
		ProductId: cartItemHertz.ProductId,
		Quantity:  cartItemHertz.Quantity,
	}

	return
}

func CartGen2Hertz(cartGen *rpcCart.Cart) (cartHertz *hertzCart.Cart) {
	cartHertz = &hertzCart.Cart{
		UserId: cartGen.UserId,
		Items:  make([]*hertzCart.CartItem, len(cartGen.Items)),
	}
	for i, item := range cartGen.Items {
		cartHertz.Items[i] = &hertzCart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}

	return
}

func OrderListGen2Hertz(orderGen []*rpcOrder.Order) (orderHertz []*hertzOrder.Order) {

	orderHertz = make([]*hertzOrder.Order, len(orderGen))
	for i, order := range orderGen {

		orderItems := make([]*hertzOrder.OrderItem, len(order.OrderItems))
		for j, item := range order.OrderItems {
			orderItems[j] = &hertzOrder.OrderItem{
				Item: &hertzCart.CartItem{
					ProductId: item.Item.ProductId,
					Quantity:  item.Item.Quantity,
				},
				Cost: item.Cost,
			}
		}

		orderHertz[i] = &hertzOrder.Order{
			OrderId:      order.OrderId,
			UserId:       order.UserId,
			UserCurrency: order.UserCurrency,
			Email:        order.Email,
			CreatedAt:    order.CreatedAt,
			MarkedPaid:   order.MarkedPaid,
			Address: &hertzOrder.Address{
				StreetAddress: order.Address.StreetAddress,
				City:          order.Address.City,
				State:         order.Address.State,
				ZipCode:       order.Address.ZipCode,
				Country:       order.Address.Country,
			},
			OrderItems: orderItems,
		}
	}

	return orderHertz
}
