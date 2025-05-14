.PONY: gen_frontend
gen_frontend:
	@cd app/frontend && cwgo server --type HTTP -I ../../idl --idl ../../idl/frontend/product_page.proto  --service frontend -module byte_go/backend/app/frontend

.PONY: gen_user
gen_user:
	@cd rpc_gen && cwgo client --type RPC --server_name user --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --server_name user --module byte_go/backend/app/user --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PONY: gen_auth
gen_auth:
	@cd rpc_gen && cwgo client --type RPC --server_name auth --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/auth.proto
	@cd app/auth && cwgo server --type RPC --server_name auth --module byte_go/backend/app/auth --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/auth.proto

.PONY: gen_product
gen_product:
	@cd rpc_gen && cwgo client --type RPC --server_name product --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --server_name product --module byte_go/backend/app/product --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PONY: gen_cart
gen_cart:
	@cd rpc_gen && cwgo client --type RPC --server_name cart --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --server_name cart --module byte_go/backend/app/cart --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PONY: gen_order
gen_order:
	@cd rpc_gen && cwgo client --type RPC --server_name order --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --server_name order --module byte_go/backend/app/order --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto

.PONY: gen_payment
gen_payment:
	@cd rpc_gen && cwgo client --type RPC --server_name payment --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --server_name payment --module byte_go/backend/app/payment --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto

.PONY: gen_checkout
gen_checkout:
	@cd rpc_gen && cwgo client --type RPC --server_name checkout --module byte_go/backend/rpc_gen -I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --server_name checkout --module byte_go/backend/app/checkout --pass "-use byte_go/backend/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto
