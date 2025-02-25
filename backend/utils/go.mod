module byte_go/backend/utils

go 1.22.5

replace byte_go/backend/constants => ../constants

require (
	byte_go/backend/constants v0.0.0-00010101000000-000000000000
	github.com/bytedance/gopkg v0.1.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/redis/go-redis/v9 v9.7.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
