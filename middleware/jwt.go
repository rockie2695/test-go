package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"test-go/database"
	"test-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(account models.Account, expiresAt time.Time) (string, error) {
	// expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, models.AuthCustomClaims{
		UserID: account.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   account.Username,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "test-go",
		},
	})

	tokenString, err := token.SignedString(models.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authentication failed.",
		})
		return
	}
	id, username, err := validateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authentication failed.",
		})
		return
	}

	account, err := models.GetAccountById(database.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authentication failed.",
		})
		return
	}
	c.Set("id", id)
	c.Set("username", username)
	c.Set("account", account)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}

func validateToken(tokenString string) (uint64, string, error) {
	var claims models.AuthCustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return models.JwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	id := claims.UserID
	username := claims.Subject
	return id, username, nil
}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	a, b := strings.ToLower(authType), strings.ToLower("Bearer")
	if a != b {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
