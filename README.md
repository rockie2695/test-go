# test-go

follow [this link](https://blog.logrocket.com/rest-api-golang-gin-gorm/)

``` go
go mod init test-go
go get github.com/gin-gonic/gin gorm.io/gorm
go get github.com/go-sql-driver/mysql
go get github.com/go-sql-driver/postgres
go get github.com/gin-contrib/cors
go get golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/swag
swag init --parseDependency --parseInternal
go run .
```

## format swagger

```go
swag fmt
```
