module github.com/Genvekt/kudos-vault/service/auth

replace github.com/Genvekt/kudos-vault/library/api => ../../library/api

go 1.22.5

require (
	github.com/Genvekt/kudos-vault/library/api v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.6.0
	golang.org/x/crypto v0.30.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.3
)

require (
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
)
