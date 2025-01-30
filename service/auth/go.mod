module github.com/Genvekt/kudos-vault/service/auth

replace github.com/Genvekt/kudos-vault/library/api => ../../library/api

replace github.com/Genvekt/kudos-vault/library/pg_client => ../../library/pg_client

go 1.22.5

require (
	github.com/Genvekt/kudos-vault/library/api v0.0.0-00010101000000-000000000000
	github.com/Genvekt/kudos-vault/library/pg_client v0.0.0-00010101000000-000000000000
	github.com/Masterminds/squirrel v1.5.4
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.6.0
	golang.org/x/crypto v0.30.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.3
)

require (
	github.com/georgysavva/scany v1.2.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
)
