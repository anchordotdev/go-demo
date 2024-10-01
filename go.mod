module go-demo

go 1.20

require (
	github.com/anchordotdev/anchor-go v0.0.0-20221027025216-20181e03b5a5
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.27.0
)

require (
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

replace anchor.dev/stolt45/localhost/pki-go => ./anchor.dev/stolt45/localhost/pki-go
