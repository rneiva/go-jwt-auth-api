# go-jwt-auth-api

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple Rest API in Golang with JWT Authentication, Fiber Web Framework, PostgreSQL DB with GORM.

### Prerequisites:

- [Golang](https://go.dev/doc/install)
  - [Installation Guide](https://golangdocs.com/install-go-windows)
- [Go-Fiber](https://docs.gofiber.io/)
- [Postgres](https://www.postgresql.org/download/)
- [GORM](https://gorm.io/index.html)

### 🛠️ Installing Packages:

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

### 💻 How to run

```
go run main.go
```
