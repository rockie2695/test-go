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

func AccountsAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Account{})
}

// @Security     BearerAuth
// @Summary      Get all accounts
// @Description  Get all accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.AccountsResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /accounts [get]
func GetAccounts(c *gin.Context) {
	accounts, err := models.GetAccounts(database.Db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	var accountsResponse []models.AccountWithoutPassword
	for _, account := range accounts {
		accountsResponse = append(accountsResponse, models.AccountWithoutPassword{
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

// @Security     BearerAuth
// @Summary      Get account by id
// @Description  Get account by id
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.AccountResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /accounts/{id} [get]
func GetAccountById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	account, err := models.GetAccountById(database.Db, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	accountResponse := models.AccountWithoutPassword{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusOK, gin.H{"account": accountResponse})
}

// @Security     BearerAuth
// @Summary      Create account
// @Description  Create account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body      models.Account  true  "Account Data"
// @Success      201       {object}  models.AccountResponse
// @Failure      400       {object}  models.HTTPError
// @Failure      500       {object}  models.HTTPError
// @Router       /accounts [post]
func CreateAccount(c *gin.Context) {
	var account models.Account
	var err error

	if err := c.ShouldBindJSON(&account); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if account.Username == "" || account.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect parameters"})
		return
	}
	account.Password, err = middleware.HashPassword(account.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	account, err = models.CreateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	accountResponse := models.AccountWithoutPassword{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusCreated, gin.H{"account": accountResponse})
}

// @Security     BearerAuth
// @Summary      Update account
// @Description  Update account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "account id"
// @Param        account  body      models.AccountWithoutPassword  true  "account data"
// @Success      200  {object}  models.AccountWithoutPassword
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /accounts/{id} [put]
func UpdateAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	account, err := models.GetAccountById(database.Db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var newAccount models.Account
	c.ShouldBindJSON(&newAccount)
	account.Username = newAccount.Username
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Set("username", account.Username)
	accountResponse := models.AccountWithoutPassword{
		ID:        account.ID,
		Token:     account.Token,
		Username:  account.Username,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
	c.JSON(http.StatusOK, gin.H{"account": accountResponse})
}

// @Security     BearerAuth
// @Summary      Delete account
// @Description  Delete account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "account id"
// @Success      200  {object}  models.HTTPResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = models.DeleteAccount(database.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account deleted"})
}

// @Summary      Login
// @Description  Login
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        loginData  body      models.AccountLoginData  true  "credentials"
// @Success      200  {object}  models.AccountLoginResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /accounts/login [post]
func Login(c *gin.Context) {
	var loginData models.AccountLoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "incorrect parameters",
		})
		return
	}
	account, err := models.GetAccountByUsername(database.Db, loginData.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("account %s not found", loginData.Username),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !middleware.CheckPasswordHash(loginData.Password, account.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "incorrect password",
		})
		return
	}
	expiresAt := time.Now().Add(12 * time.Hour)
	token, err := middleware.GenerateToken(account, expiresAt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// update token to account
	account.Token = token
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	accountResponse := models.AccountWithoutPassword{
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

// @Security     BearerAuth
// @Summary      Logout
// @Description  Logout account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.HTTPResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Security     BearerAuth
// @Router       /accounts/logout [post]
func Logout(c *gin.Context) {
	id, _ := c.Get("id")

	account, err := models.GetAccountById(database.Db, id.(uint64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	account.Token = ""
	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

// @Security     BearerAuth
// @Summary      Update account password
// @Description  Update account password
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "account id"
// @Param        accountChangePassword  body      models.AccountChangePassword  true  "credentials"
// @Success      200  {object}  models.AccountWithoutPassword
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Security     BearerAuth
// @Router       /accounts/{id}/password [put]
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
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if !middleware.CheckPasswordHash(accountChangePassword.OldPassword, account.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "incorrect password",
		})
		return
	}

	newHash, newHashErr := middleware.HashPassword(accountChangePassword.NewPassword)
	if newHashErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": newHashErr})
		return
	}

	account.Password = newHash

	account, err = models.UpdateAccount(database.Db, account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, account)
}
