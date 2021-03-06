module github.com/nigoroku/amb-user/service

go 1.14

require (
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20200616145509-8d140a17f351 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	local.packages/models v0.0.0-00010101000000-000000000000

)

replace local.packages/models => ./../models
