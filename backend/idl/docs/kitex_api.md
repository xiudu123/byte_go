# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth.proto](#auth-proto)
    - [DeleteTokenListReq](#auth-DeleteTokenListReq)
    - [DeleteTokenReq](#auth-DeleteTokenReq)
    - [DeliverTokenReq](#auth-DeliverTokenReq)
    - [DeliveryResp](#auth-DeliveryResp)
    - [VerifyResp](#auth-VerifyResp)
    - [VerifyTokenReq](#auth-VerifyTokenReq)
  
    - [AuthService](#auth-AuthService)
  
- [cart.proto](#cart-proto)
    - [AddItemReq](#cart-AddItemReq)
    - [AddItemResp](#cart-AddItemResp)
    - [Cart](#cart-Cart)
    - [CartItem](#cart-CartItem)
    - [EmptyCartReq](#cart-EmptyCartReq)
    - [EmptyCartResp](#cart-EmptyCartResp)
    - [GetCartReq](#cart-GetCartReq)
    - [GetCartResp](#cart-GetCartResp)
  
    - [CartService](#cart-CartService)
  
- [checkout.proto](#checkout-proto)
    - [Address](#checkout-Address)
    - [CheckoutReq](#checkout-CheckoutReq)
    - [CheckoutResp](#checkout-CheckoutResp)
  
    - [CheckoutService](#checkout-CheckoutService)
  
- [common.proto](#common-proto)
    - [Empty](#common-Empty)
  
- [order.proto](#order-proto)
    - [Address](#order-Address)
    - [ListOrderReq](#order-ListOrderReq)
    - [ListOrderResp](#order-ListOrderResp)
    - [MarkOrderPaidReq](#order-MarkOrderPaidReq)
    - [MarkOrderPaidResp](#order-MarkOrderPaidResp)
    - [Order](#order-Order)
    - [OrderItem](#order-OrderItem)
    - [OrderResult](#order-OrderResult)
    - [PlaceOrderReq](#order-PlaceOrderReq)
    - [PlaceOrderResp](#order-PlaceOrderResp)
  
    - [OrderService](#order-OrderService)
  
- [payment.proto](#payment-proto)
    - [ChargeReq](#payment-ChargeReq)
    - [ChargeResp](#payment-ChargeResp)
    - [CreditCardInfo](#payment-CreditCardInfo)
  
    - [PaymentService](#payment-PaymentService)
  
- [product.proto](#product-proto)
    - [CreateCategoryReq](#product-CreateCategoryReq)
    - [CreateCategoryResp](#product-CreateCategoryResp)
    - [CreateProductReq](#product-CreateProductReq)
    - [CreateProductResp](#product-CreateProductResp)
    - [DeleteProductReq](#product-DeleteProductReq)
    - [DeleteProductResp](#product-DeleteProductResp)
    - [GetProductReq](#product-GetProductReq)
    - [GetProductResp](#product-GetProductResp)
    - [ListProductByIdsReq](#product-ListProductByIdsReq)
    - [ListProductByIdsResp](#product-ListProductByIdsResp)
    - [ListProductsReq](#product-ListProductsReq)
    - [ListProductsResp](#product-ListProductsResp)
    - [Product](#product-Product)
    - [SearchProductsReq](#product-SearchProductsReq)
    - [SearchProductsResp](#product-SearchProductsResp)
    - [UpdateProductReq](#product-UpdateProductReq)
    - [UpdateProductResp](#product-UpdateProductResp)
  
    - [ProductCatalogService](#product-ProductCatalogService)
  
- [user.proto](#user-proto)
    - [DeleteUserReq](#user-DeleteUserReq)
    - [GetUserInfoReq](#user-GetUserInfoReq)
    - [GetUserInfoResp](#user-GetUserInfoResp)
    - [LoginReq](#user-LoginReq)
    - [LoginResp](#user-LoginResp)
    - [RegisterReq](#user-RegisterReq)
    - [RegisterResp](#user-RegisterResp)
    - [UpdateUserReq](#user-UpdateUserReq)
    - [UpdateUserResp](#user-UpdateUserResp)
    - [User](#user-User)
  
    - [UserService](#user-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth.proto



<a name="auth-DeleteTokenListReq"></a>

### DeleteTokenListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="auth-DeleteTokenReq"></a>

### DeleteTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jti | [string](#string) |  |  |






<a name="auth-DeliverTokenReq"></a>

### DeliverTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="auth-DeliveryResp"></a>

### DeliveryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="auth-VerifyResp"></a>

### VerifyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| jti | [string](#string) |  |  |






<a name="auth-VerifyTokenReq"></a>

### VerifyTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |





 

 

 


<a name="auth-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DeliverTokenByRPC | [DeliverTokenReq](#auth-DeliverTokenReq) | [DeliveryResp](#auth-DeliveryResp) |  |
| VerifyTokenByRPC | [VerifyTokenReq](#auth-VerifyTokenReq) | [VerifyResp](#auth-VerifyResp) |  |
| DeleteTokenByRPC | [DeleteTokenReq](#auth-DeleteTokenReq) | [.common.Empty](#common-Empty) |  |
| DeleteTokenListByRPC | [DeleteTokenListReq](#auth-DeleteTokenListReq) | [.common.Empty](#common-Empty) |  |

 



<a name="cart-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cart.proto



<a name="cart-AddItemReq"></a>

### AddItemReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| item | [CartItem](#cart-CartItem) |  |  |






<a name="cart-AddItemResp"></a>

### AddItemResp







<a name="cart-Cart"></a>

### Cart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| items | [CartItem](#cart-CartItem) | repeated |  |






<a name="cart-CartItem"></a>

### CartItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| quantity | [uint32](#uint32) |  |  |






<a name="cart-EmptyCartReq"></a>

### EmptyCartReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="cart-EmptyCartResp"></a>

### EmptyCartResp







<a name="cart-GetCartReq"></a>

### GetCartReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="cart-GetCartResp"></a>

### GetCartResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cart | [Cart](#cart-Cart) |  |  |





 

 

 


<a name="cart-CartService"></a>

### CartService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddItem | [AddItemReq](#cart-AddItemReq) | [AddItemResp](#cart-AddItemResp) |  |
| GetCart | [GetCartReq](#cart-GetCartReq) | [GetCartResp](#cart-GetCartResp) |  |
| EmptyCart | [EmptyCartReq](#cart-EmptyCartReq) | [EmptyCartResp](#cart-EmptyCartResp) |  |

 



<a name="checkout-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## checkout.proto



<a name="checkout-Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street_address | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| country | [string](#string) |  |  |
| zip_code | [int32](#int32) |  |  |






<a name="checkout-CheckoutReq"></a>

### CheckoutReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| firstname | [string](#string) |  |  |
| lastname | [string](#string) |  |  |
| email | [string](#string) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#checkout-Address) |  |  |
| credit_card | [payment.CreditCardInfo](#payment-CreditCardInfo) |  |  |






<a name="checkout-CheckoutResp"></a>

### CheckoutResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  |  |
| transaction_id | [string](#string) |  |  |





 

 

 


<a name="checkout-CheckoutService"></a>

### CheckoutService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Checkout | [CheckoutReq](#checkout-CheckoutReq) | [CheckoutResp](#checkout-CheckoutResp) |  |

 



<a name="common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="common-Empty"></a>

### Empty






 

 

 

 



<a name="order-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order.proto



<a name="order-Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street_address | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| country | [string](#string) |  |  |
| zip_code | [int32](#int32) |  |  |






<a name="order-ListOrderReq"></a>

### ListOrderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="order-ListOrderResp"></a>

### ListOrderResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#order-Order) | repeated |  |






<a name="order-MarkOrderPaidReq"></a>

### MarkOrderPaidReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| order_id | [string](#string) |  |  |






<a name="order-MarkOrderPaidResp"></a>

### MarkOrderPaidResp







<a name="order-Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_items | [OrderItem](#order-OrderItem) | repeated |  |
| order_id | [string](#string) |  |  |
| user_id | [uint32](#uint32) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#order-Address) |  |  |
| email | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| marked_paid | [bool](#bool) |  | 标记是否支付 |






<a name="order-OrderItem"></a>

### OrderItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [cart.CartItem](#cart-CartItem) |  |  |
| cost | [float](#float) |  |  |






<a name="order-OrderResult"></a>

### OrderResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  |  |






<a name="order-PlaceOrderReq"></a>

### PlaceOrderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#order-Address) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| order_items | [OrderItem](#order-OrderItem) | repeated |  |






<a name="order-PlaceOrderResp"></a>

### PlaceOrderResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [OrderResult](#order-OrderResult) |  |  |





 

 

 


<a name="order-OrderService"></a>

### OrderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PlaceOrder | [PlaceOrderReq](#order-PlaceOrderReq) | [PlaceOrderResp](#order-PlaceOrderResp) |  |
| ListOrder | [ListOrderReq](#order-ListOrderReq) | [ListOrderResp](#order-ListOrderResp) |  |
| MarkOrderPaid | [MarkOrderPaidReq](#order-MarkOrderPaidReq) | [MarkOrderPaidResp](#order-MarkOrderPaidResp) |  |

 



<a name="payment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payment.proto



<a name="payment-ChargeReq"></a>

### ChargeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [float](#float) |  |  |
| credit_card | [CreditCardInfo](#payment-CreditCardInfo) |  |  |
| order_id | [string](#string) |  |  |
| user_id | [uint32](#uint32) |  |  |






<a name="payment-ChargeResp"></a>

### ChargeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  |  |






<a name="payment-CreditCardInfo"></a>

### CreditCardInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_card_number | [string](#string) |  |  |
| credit_card_cvv | [int32](#int32) |  |  |
| credit_card_expiration_year | [int32](#int32) |  |  |
| credit_card_expiration_month | [int32](#int32) |  |  |





 

 

 


<a name="payment-PaymentService"></a>

### PaymentService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Charge | [ChargeReq](#payment-ChargeReq) | [ChargeResp](#payment-ChargeResp) |  |

 



<a name="product-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product.proto



<a name="product-CreateCategoryReq"></a>

### CreateCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="product-CreateCategoryResp"></a>

### CreateCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [uint32](#uint32) |  |  |






<a name="product-CreateProductReq"></a>

### CreateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="product-CreateProductResp"></a>

### CreateProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="product-DeleteProductReq"></a>

### DeleteProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="product-DeleteProductResp"></a>

### DeleteProductResp







<a name="product-GetProductReq"></a>

### GetProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="product-GetProductResp"></a>

### GetProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [Product](#product-Product) |  |  |






<a name="product-ListProductByIdsReq"></a>

### ListProductByIdsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_ids | [uint32](#uint32) | repeated |  |






<a name="product-ListProductByIdsResp"></a>

### ListProductByIdsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#product-Product) | repeated |  |






<a name="product-ListProductsReq"></a>

### ListProductsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| pageSize | [int64](#int64) |  |  |
| categoryName | [string](#string) |  |  |






<a name="product-ListProductsResp"></a>

### ListProductsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#product-Product) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="product-Product"></a>

### Product



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="product-SearchProductsReq"></a>

### SearchProductsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="product-SearchProductsResp"></a>

### SearchProductsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#product-Product) | repeated |  |






<a name="product-UpdateProductReq"></a>

### UpdateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="product-UpdateProductResp"></a>

### UpdateProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |





 

 

 


<a name="product-ProductCatalogService"></a>

### ProductCatalogService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProducts | [ListProductsReq](#product-ListProductsReq) | [ListProductsResp](#product-ListProductsResp) |  |
| GetProduct | [GetProductReq](#product-GetProductReq) | [GetProductResp](#product-GetProductResp) |  |
| SearchProducts | [SearchProductsReq](#product-SearchProductsReq) | [SearchProductsResp](#product-SearchProductsResp) |  |
| ListProductByIds | [ListProductByIdsReq](#product-ListProductByIdsReq) | [ListProductByIdsResp](#product-ListProductByIdsResp) |  |
| CreateProduct | [CreateProductReq](#product-CreateProductReq) | [CreateProductResp](#product-CreateProductResp) |  |
| DeleteProduct | [DeleteProductReq](#product-DeleteProductReq) | [DeleteProductResp](#product-DeleteProductResp) |  |
| UpdateProduct | [UpdateProductReq](#product-UpdateProductReq) | [UpdateProductResp](#product-UpdateProductResp) |  |
| CreateCategory | [CreateCategoryReq](#product-CreateCategoryReq) | [CreateCategoryResp](#product-CreateCategoryResp) |  |

 



<a name="user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user-DeleteUserReq"></a>

### DeleteUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| confirmation | [string](#string) |  |  |






<a name="user-GetUserInfoReq"></a>

### GetUserInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| email | [string](#string) |  |  |






<a name="user-GetUserInfoResp"></a>

### GetUserInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |






<a name="user-LoginReq"></a>

### LoginReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="user-LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |
| token | [string](#string) |  |  |






<a name="user-RegisterReq"></a>

### RegisterReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| confirm_password | [string](#string) |  |  |
| username | [string](#string) |  |  |






<a name="user-RegisterResp"></a>

### RegisterResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |
| token | [string](#string) |  |  |






<a name="user-UpdateUserReq"></a>

### UpdateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| username | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |






<a name="user-UpdateUserResp"></a>

### UpdateUserResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |






<a name="user-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| email | [string](#string) |  |  |
| username | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |





 

 

 


<a name="user-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterReq](#user-RegisterReq) | [RegisterResp](#user-RegisterResp) |  |
| Login | [LoginReq](#user-LoginReq) | [LoginResp](#user-LoginResp) |  |
| GetUserInfo | [GetUserInfoReq](#user-GetUserInfoReq) | [GetUserInfoResp](#user-GetUserInfoResp) |  |
| Logout | [.common.Empty](#common-Empty) | [.common.Empty](#common-Empty) |  |
| DeleteUser | [DeleteUserReq](#user-DeleteUserReq) | [.common.Empty](#common-Empty) |  |
| UpdateUser | [UpdateUserReq](#user-UpdateUserReq) | [UpdateUserResp](#user-UpdateUserResp) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

