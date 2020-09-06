module github.com/skwongg/jackt

go 1.14

require (
	github.com/badoux/checkmail v1.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/skwongg/jackt/api/models v0.0.0-00010101000000-000000000000
	github.com/skwongg/jackt/api/utils/formaterror v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)

replace github.com/skwongg/jackt/api/utils/formaterror => ./api/formaterror

replace github.com/skwongg/jackt/api/models => ./api/models
