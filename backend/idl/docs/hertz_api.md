# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [frontend/api.proto](#frontend_api-proto)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
    - [File-level Extensions](#frontend_api-proto-extensions)
  
- [frontend/cart_page.proto](#frontend_cart_page-proto)
    - [AddItemReq](#frontend-cart-AddItemReq)
    - [AddItemResp](#frontend-cart-AddItemResp)
    - [Cart](#frontend-cart-Cart)
    - [CartItem](#frontend-cart-CartItem)
    - [EmptyCartReq](#frontend-cart-EmptyCartReq)
    - [EmptyCartResp](#frontend-cart-EmptyCartResp)
    - [GetCartReq](#frontend-cart-GetCartReq)
    - [GetCartResp](#frontend-cart-GetCartResp)
  
    - [CartService](#frontend-cart-CartService)
  
- [frontend/checkout_page.proto](#frontend_checkout_page-proto)
    - [Address](#frontend-checkout-Address)
    - [CheckoutReq](#frontend-checkout-CheckoutReq)
    - [CheckoutResp](#frontend-checkout-CheckoutResp)
  
    - [CheckoutService](#frontend-checkout-CheckoutService)
  
- [frontend/common_hertz.proto](#frontend_common_hertz-proto)
    - [Empty](#common_hertz-Empty)
  
- [frontend/order_page.proto](#frontend_order_page-proto)
    - [Address](#frontend-order-Address)
    - [ListOrderReq](#frontend-order-ListOrderReq)
    - [ListOrderResp](#frontend-order-ListOrderResp)
    - [MarkOrderPaidReq](#frontend-order-MarkOrderPaidReq)
    - [MarkOrderPaidResp](#frontend-order-MarkOrderPaidResp)
    - [Order](#frontend-order-Order)
    - [OrderItem](#frontend-order-OrderItem)
    - [OrderResult](#frontend-order-OrderResult)
    - [PlaceOrderReq](#frontend-order-PlaceOrderReq)
    - [PlaceOrderResp](#frontend-order-PlaceOrderResp)
  
    - [OrderService](#frontend-order-OrderService)
  
- [frontend/payment_page.proto](#frontend_payment_page-proto)
    - [ChargeReq](#frontend-payment-ChargeReq)
    - [ChargeResp](#frontend-payment-ChargeResp)
    - [CreditCardInfo](#frontend-payment-CreditCardInfo)
  
    - [PaymentService](#frontend-payment-PaymentService)
  
- [frontend/product_page.proto](#frontend_product_page-proto)
    - [CreateCategoryReq](#frontend-product-CreateCategoryReq)
    - [CreateCategoryResp](#frontend-product-CreateCategoryResp)
    - [CreateProductReq](#frontend-product-CreateProductReq)
    - [CreateProductResp](#frontend-product-CreateProductResp)
    - [DeleteProductReq](#frontend-product-DeleteProductReq)
    - [DeleteProductResp](#frontend-product-DeleteProductResp)
    - [GetProductReq](#frontend-product-GetProductReq)
    - [GetProductResp](#frontend-product-GetProductResp)
    - [ListProductsReq](#frontend-product-ListProductsReq)
    - [ListProductsResp](#frontend-product-ListProductsResp)
    - [Product](#frontend-product-Product)
    - [SearchProductsReq](#frontend-product-SearchProductsReq)
    - [SearchProductsResp](#frontend-product-SearchProductsResp)
    - [UpdateProductReq](#frontend-product-UpdateProductReq)
    - [UpdateProductResp](#frontend-product-UpdateProductResp)
  
    - [ProductCatalogService](#frontend-product-ProductCatalogService)
  
- [frontend/user_page.proto](#frontend_user_page-proto)
    - [DeleteUserReq](#frontend-user-DeleteUserReq)
    - [GetUserInfoReq](#frontend-user-GetUserInfoReq)
    - [GetUserInfoResp](#frontend-user-GetUserInfoResp)
    - [LoginReq](#frontend-user-LoginReq)
    - [LoginResp](#frontend-user-LoginResp)
    - [LogoutReq](#frontend-user-LogoutReq)
    - [RegisterReq](#frontend-user-RegisterReq)
    - [RegisterResp](#frontend-user-RegisterResp)
    - [UpdateUserReq](#frontend-user-UpdateUserReq)
    - [UpdateUserResp](#frontend-user-UpdateUserResp)
    - [User](#frontend-user-User)
  
    - [UserService](#frontend-user-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="frontend_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/api.proto
idl/api.proto; 注解拓展

 

 


<a name="frontend_api-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| http_code | int32 | .google.protobuf.EnumValueOptions | 50401 |  |
| body | string | .google.protobuf.FieldOptions | 50105 |  |
| cookie | string | .google.protobuf.FieldOptions | 50104 |  |
| file_name | string | .google.protobuf.FieldOptions | 50110 |  |
| file_name_compatible | string | .google.protobuf.FieldOptions | 50133 |  |
| form | string | .google.protobuf.FieldOptions | 50108 |  |
| form_compatible | string | .google.protobuf.FieldOptions | 50131 | 50131~50160 used to extend field option by hz |
| go_tag | string | .google.protobuf.FieldOptions | 51001 |  |
| header | string | .google.protobuf.FieldOptions | 50103 |  |
| js_conv | string | .google.protobuf.FieldOptions | 50109 |  |
| js_conv_compatible | string | .google.protobuf.FieldOptions | 50132 |  |
| none | string | .google.protobuf.FieldOptions | 50111 |  |
| none_compatible | string | .google.protobuf.FieldOptions | 50134 | 50135 is reserved to vt_compatible optional FieldRules vt_compatible = 50135; |
| path | string | .google.protobuf.FieldOptions | 50106 |  |
| query | string | .google.protobuf.FieldOptions | 50102 |  |
| raw_body | string | .google.protobuf.FieldOptions | 50101 |  |
| vd | string | .google.protobuf.FieldOptions | 50107 |  |
| reserve | string | .google.protobuf.MessageOptions | 50830 | 550831 is reserved to msg_vt_compatible optional FieldRules msg_vt_compatible = 50831; |
| any | string | .google.protobuf.MethodOptions | 50208 |  |
| baseurl | string | .google.protobuf.MethodOptions | 50308 | Baseurl used in ttnet routing |
| delete | string | .google.protobuf.MethodOptions | 50204 |  |
| gen_path | string | .google.protobuf.MethodOptions | 50301 | The path specified by the user when the client code is generated, with a higher priority than api_version |
| get | string | .google.protobuf.MethodOptions | 50201 |  |
| handler_path | string | .google.protobuf.MethodOptions | 50309 | handler_path specifies the path to generate the method |
| handler_path_compatible | string | .google.protobuf.MethodOptions | 50331 | 50331~50360 used to extend method option by hz

handler_path specifies the path to generate the method |
| head | string | .google.protobuf.MethodOptions | 50207 |  |
| name | string | .google.protobuf.MethodOptions | 50304 | Name of rpc |
| options | string | .google.protobuf.MethodOptions | 50206 |  |
| param | string | .google.protobuf.MethodOptions | 50307 | Whether client requests take public parameters |
| patch | string | .google.protobuf.MethodOptions | 50205 |  |
| post | string | .google.protobuf.MethodOptions | 50202 |  |
| put | string | .google.protobuf.MethodOptions | 50203 |  |
| serializer | string | .google.protobuf.MethodOptions | 50306 | Serialization method |
| tag | string | .google.protobuf.MethodOptions | 50303 | rpc tag, can be multiple, separated by commas |
| base_domain | string | .google.protobuf.ServiceOptions | 50402 |  |
| base_domain_compatible | string | .google.protobuf.ServiceOptions | 50731 | 50731~50760 used to extend service option by hz |
| api_level | string | .google.protobuf.MethodOptions | 50305 | Interface Level |
| api_version | string | .google.protobuf.MethodOptions | 50302 | Specify the value of the :version variable in path when the client code is generated |

 

 



<a name="frontend_cart_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/cart_page.proto



<a name="frontend-cart-AddItemReq"></a>

### AddItemReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| item | [CartItem](#frontend-cart-CartItem) |  |  |






<a name="frontend-cart-AddItemResp"></a>

### AddItemResp







<a name="frontend-cart-Cart"></a>

### Cart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| items | [CartItem](#frontend-cart-CartItem) | repeated |  |






<a name="frontend-cart-CartItem"></a>

### CartItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| quantity | [uint32](#uint32) |  |  |






<a name="frontend-cart-EmptyCartReq"></a>

### EmptyCartReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="frontend-cart-EmptyCartResp"></a>

### EmptyCartResp







<a name="frontend-cart-GetCartReq"></a>

### GetCartReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="frontend-cart-GetCartResp"></a>

### GetCartResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cart | [Cart](#frontend-cart-Cart) |  |  |





 

 

 


<a name="frontend-cart-CartService"></a>

### CartService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddItem | [AddItemReq](#frontend-cart-AddItemReq) | [AddItemResp](#frontend-cart-AddItemResp) |  |
| GetCart | [GetCartReq](#frontend-cart-GetCartReq) | [GetCartResp](#frontend-cart-GetCartResp) |  |
| EmptyCart | [EmptyCartReq](#frontend-cart-EmptyCartReq) | [EmptyCartResp](#frontend-cart-EmptyCartResp) |  |

 



<a name="frontend_checkout_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/checkout_page.proto



<a name="frontend-checkout-Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street_address | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| country | [string](#string) |  |  |
| zip_code | [int32](#int32) |  |  |






<a name="frontend-checkout-CheckoutReq"></a>

### CheckoutReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| firstname | [string](#string) |  |  |
| lastname | [string](#string) |  |  |
| email | [string](#string) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#frontend-checkout-Address) |  |  |
| credit_card | [frontend.payment.CreditCardInfo](#frontend-payment-CreditCardInfo) |  |  |






<a name="frontend-checkout-CheckoutResp"></a>

### CheckoutResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  |  |
| transaction_id | [string](#string) |  |  |





 

 

 


<a name="frontend-checkout-CheckoutService"></a>

### CheckoutService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Checkout | [CheckoutReq](#frontend-checkout-CheckoutReq) | [CheckoutResp](#frontend-checkout-CheckoutResp) |  |

 



<a name="frontend_common_hertz-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/common_hertz.proto



<a name="common_hertz-Empty"></a>

### Empty






 

 

 

 



<a name="frontend_order_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/order_page.proto



<a name="frontend-order-Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street_address | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| country | [string](#string) |  |  |
| zip_code | [int32](#int32) |  |  |






<a name="frontend-order-ListOrderReq"></a>

### ListOrderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="frontend-order-ListOrderResp"></a>

### ListOrderResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#frontend-order-Order) | repeated |  |






<a name="frontend-order-MarkOrderPaidReq"></a>

### MarkOrderPaidReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| order_id | [string](#string) |  |  |






<a name="frontend-order-MarkOrderPaidResp"></a>

### MarkOrderPaidResp







<a name="frontend-order-Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_items | [OrderItem](#frontend-order-OrderItem) | repeated |  |
| order_id | [string](#string) |  |  |
| user_id | [uint32](#uint32) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#frontend-order-Address) |  |  |
| email | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| marked_paid | [bool](#bool) |  | 标记是否支付 |






<a name="frontend-order-OrderItem"></a>

### OrderItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [frontend.cart.CartItem](#frontend-cart-CartItem) |  |  |
| cost | [float](#float) |  |  |






<a name="frontend-order-OrderResult"></a>

### OrderResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  |  |






<a name="frontend-order-PlaceOrderReq"></a>

### PlaceOrderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| user_currency | [string](#string) |  |  |
| address | [Address](#frontend-order-Address) |  |  |
| email | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| order_items | [OrderItem](#frontend-order-OrderItem) | repeated |  |






<a name="frontend-order-PlaceOrderResp"></a>

### PlaceOrderResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [OrderResult](#frontend-order-OrderResult) |  |  |





 

 

 


<a name="frontend-order-OrderService"></a>

### OrderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PlaceOrder | [PlaceOrderReq](#frontend-order-PlaceOrderReq) | [PlaceOrderResp](#frontend-order-PlaceOrderResp) |  |
| ListOrder | [ListOrderReq](#frontend-order-ListOrderReq) | [ListOrderResp](#frontend-order-ListOrderResp) |  |
| MarkOrderPaid | [MarkOrderPaidReq](#frontend-order-MarkOrderPaidReq) | [MarkOrderPaidResp](#frontend-order-MarkOrderPaidResp) |  |

 



<a name="frontend_payment_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/payment_page.proto



<a name="frontend-payment-ChargeReq"></a>

### ChargeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [float](#float) |  |  |
| credit_card | [CreditCardInfo](#frontend-payment-CreditCardInfo) |  |  |
| order_id | [string](#string) |  |  |
| user_id | [uint32](#uint32) |  |  |






<a name="frontend-payment-ChargeResp"></a>

### ChargeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  |  |






<a name="frontend-payment-CreditCardInfo"></a>

### CreditCardInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credit_card_number | [string](#string) |  |  |
| credit_card_cvv | [int32](#int32) |  |  |
| credit_card_expiration_year | [int32](#int32) |  |  |
| credit_card_expiration_month | [int32](#int32) |  |  |





 

 

 


<a name="frontend-payment-PaymentService"></a>

### PaymentService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Charge | [ChargeReq](#frontend-payment-ChargeReq) | [ChargeResp](#frontend-payment-ChargeResp) |  |

 



<a name="frontend_product_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/product_page.proto



<a name="frontend-product-CreateCategoryReq"></a>

### CreateCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="frontend-product-CreateCategoryResp"></a>

### CreateCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [uint32](#uint32) |  |  |






<a name="frontend-product-CreateProductReq"></a>

### CreateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="frontend-product-CreateProductResp"></a>

### CreateProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="frontend-product-DeleteProductReq"></a>

### DeleteProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="frontend-product-DeleteProductResp"></a>

### DeleteProductResp







<a name="frontend-product-GetProductReq"></a>

### GetProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |






<a name="frontend-product-GetProductResp"></a>

### GetProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [Product](#frontend-product-Product) |  |  |






<a name="frontend-product-ListProductsReq"></a>

### ListProductsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| pageSize | [int64](#int64) |  |  |
| categoryName | [string](#string) |  |  |






<a name="frontend-product-ListProductsResp"></a>

### ListProductsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#frontend-product-Product) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="frontend-product-Product"></a>

### Product



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="frontend-product-SearchProductsReq"></a>

### SearchProductsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="frontend-product-SearchProductsResp"></a>

### SearchProductsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Products | [Product](#frontend-product-Product) | repeated |  |






<a name="frontend-product-UpdateProductReq"></a>

### UpdateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| picture | [string](#string) |  |  |
| price | [float](#float) |  |  |
| categories | [string](#string) | repeated |  |






<a name="frontend-product-UpdateProductResp"></a>

### UpdateProductResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [uint32](#uint32) |  |  |





 

 

 


<a name="frontend-product-ProductCatalogService"></a>

### ProductCatalogService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProducts | [ListProductsReq](#frontend-product-ListProductsReq) | [ListProductsResp](#frontend-product-ListProductsResp) |  |
| GetProduct | [GetProductReq](#frontend-product-GetProductReq) | [GetProductResp](#frontend-product-GetProductResp) |  |
| SearchProducts | [SearchProductsReq](#frontend-product-SearchProductsReq) | [SearchProductsResp](#frontend-product-SearchProductsResp) |  |
| CreateProduct | [CreateProductReq](#frontend-product-CreateProductReq) | [CreateProductResp](#frontend-product-CreateProductResp) |  |
| DeleteProduct | [DeleteProductReq](#frontend-product-DeleteProductReq) | [DeleteProductResp](#frontend-product-DeleteProductResp) |  |
| UpdateProduct | [UpdateProductReq](#frontend-product-UpdateProductReq) | [UpdateProductResp](#frontend-product-UpdateProductResp) |  |
| CreateCategory | [CreateCategoryReq](#frontend-product-CreateCategoryReq) | [CreateCategoryResp](#frontend-product-CreateCategoryResp) |  |

 



<a name="frontend_user_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## frontend/user_page.proto



<a name="frontend-user-DeleteUserReq"></a>

### DeleteUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |






<a name="frontend-user-GetUserInfoReq"></a>

### GetUserInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| email | [string](#string) |  |  |






<a name="frontend-user-GetUserInfoResp"></a>

### GetUserInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#frontend-user-User) |  |  |






<a name="frontend-user-LoginReq"></a>

### LoginReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="frontend-user-LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#frontend-user-User) |  |  |
| access_token | [string](#string) |  |  |






<a name="frontend-user-LogoutReq"></a>

### LogoutReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |






<a name="frontend-user-RegisterReq"></a>

### RegisterReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| confirm_password | [string](#string) |  |  |
| username | [string](#string) |  |  |






<a name="frontend-user-RegisterResp"></a>

### RegisterResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#frontend-user-User) |  |  |
| access_token | [string](#string) |  |  |






<a name="frontend-user-UpdateUserReq"></a>

### UpdateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| username | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |






<a name="frontend-user-UpdateUserResp"></a>

### UpdateUserResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#frontend-user-User) |  |  |






<a name="frontend-user-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint32](#uint32) |  |  |
| email | [string](#string) |  |  |
| username | [string](#string) |  |  |
| avatar_url | [string](#string) |  |  |





 

 

 


<a name="frontend-user-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterReq](#frontend-user-RegisterReq) | [RegisterResp](#frontend-user-RegisterResp) |  |
| Login | [LoginReq](#frontend-user-LoginReq) | [LoginResp](#frontend-user-LoginResp) |  |
| GetUserInfo | [GetUserInfoReq](#frontend-user-GetUserInfoReq) | [GetUserInfoResp](#frontend-user-GetUserInfoResp) |  |
| Logout | [LogoutReq](#frontend-user-LogoutReq) | [.common_hertz.Empty](#common_hertz-Empty) |  |
| DeleteUser | [DeleteUserReq](#frontend-user-DeleteUserReq) | [.common_hertz.Empty](#common_hertz-Empty) |  |
| UpdateUser | [UpdateUserReq](#frontend-user-UpdateUserReq) | [UpdateUserResp](#frontend-user-UpdateUserResp) |  |

 



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

