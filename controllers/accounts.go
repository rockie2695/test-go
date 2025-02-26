package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"test-go/database"
	"test-go/middleware"
	"test-go/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountsAutoMigrate() {
	database.Db.AutoMigrate(&models.Account{})
}

func GetAccounts(c *gin.Context) {
	accounts, err := models.GetAccounts(database.Db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var accountsResponse []models.AccountResponse
	for _, account := range accounts {
		accountsResponse = append(accountsResponse, models.AccountResponse{
			ID:        account.ID,
			Token:     account.Token,
			Username:  account.Username,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
			DeletedAt: account.DeletedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accountsResponse,
		"count":    len(accountsResponse),
	})
}
func GetAccountById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	account, err := models.GetAccountById(database.Db, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	accountResponse := models.AccountResponse{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusOK, gin.H{"account": accountResponse})
}
func CreateAccount(c *gin.Context) {
	var account models.Account
	var err error

	if err := c.ShouldBindJSON(&account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account.Username == "" || account.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		return
	}
	account.Password, err = middleware.HashPassword(account.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	account, err = models.CreateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountResponse := models.AccountResponse{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusCreated, gin.H{"account": accountResponse})
}
func UpdateAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	account, err := models.GetAccountById(database.Db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var newAccount models.Account
	c.ShouldBindJSON(&newAccount)
	account.Username = newAccount.Username
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("username", account.Username)
	accountResponse := models.AccountResponse{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusOK, gin.H{"account": accountResponse})
}
func DeleteAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteAccount(database.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account deleted"})
}
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	account, err := models.GetAccountByUsername(database.Db, loginData.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("account %s not found", loginData.Username),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !middleware.CheckPasswordHash(loginData.Password, account.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}
	expiresAt := time.Now().Add(12 * time.Hour)
	token, err := middleware.GenerateToken(account, expiresAt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update token to account
	account.Token = token
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	accountResponse := models.AccountResponse{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"account":    accountResponse,
		"expires_at": expiresAt.Format(time.RFC3339),
	})
}

func Logout(c *gin.Context) {
	id, _ := c.Get("id")

	account, err := models.GetAccountById(database.Db, id.(uint64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	account.Token = ""
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

func UpdateAccountPassword(c *gin.Context) {
	var accountChangePassword models.AccountChangePassword
	c.ShouldBindJSON(&accountChangePassword)

	if accountChangePassword.OldPassword == "" || accountChangePassword.NewPassword == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "incorrect parameters",
		})
		return
	}

	if accountChangePassword.OldPassword == accountChangePassword.NewPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "New password cannot be same as old password",
		})
		return
	}
	var account models.Account
	id, _ := strconv.Atoi(c.Param("id"))
	account, err := models.GetAccountById(database.Db, uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.ShouldBindJSON(&account)

	if !middleware.CheckPasswordHash(accountChangePassword.OldPassword, account.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	newHash, newHashErr := middleware.HashPassword(accountChangePassword.NewPassword)
	if newHashErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": newHashErr})
		return
	}

	account.Password = newHash

	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, account)
}
