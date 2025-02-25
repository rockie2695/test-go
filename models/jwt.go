package models

import "github.com/golang-jwt/jwt/v5"

var JwtKey = []byte("secretkey")

type AuthCustomClaims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"userId"`
}
