# go-jwt-auth-api

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![](https://img.shields.io/github/go-mod/go-version/rneiva/go-jwt-auth-api)

Simple Rest API in Golang with:

- JWT Authentication
- Fiber Web Framework
- PostgreSQL DB and GORM

### Prerequisites:

- Installations :

- [Golang](https://go.dev/doc/install)
- [Go-Fiber](https://docs.gofiber.io/)
- [Postgres](https://www.postgresql.org/download/)
- [GORM](https://gorm.io/index.html)

### Installing Packages:

```
go mod init <repository_name>
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-jwt/jwt
```

```
go mod tidy
```

### How to run 💻

```
go run main.go
```
