package main

import (
	// "encoding/json"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	// "strings"
	"test-go/config"
	"test-go/database"
	"test-go/models"
	"test-go/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountsRoute(t *testing.T) {
	config.EnvSetup()
	db := database.InitDb()
	router := router.InitRouter(db)

	w := httptest.NewRecorder()
	//login first
	loginData := models.AccountLoginData{
		Username: "a",
		Password: "b",
	}
	jsonLoginData, _ := json.Marshal(loginData)
	reqLogin, _ := http.NewRequest("POST", "/api/v1/accounts/login", strings.NewReader(string(jsonLoginData)))
	router.ServeHTTP(w, reqLogin)

	//put token to auth
	var loginResponse models.AccountLoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)

	token := loginResponse.Token
	req, _ := http.NewRequest("GET", "/api/v1/accounts", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	//make return to string
	resBody := w.Body.String()
	//return body is model.AccountsResponse
	var accountsResponse models.AccountsResponse
	err = json.Unmarshal([]byte(resBody), &accountsResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, accountsResponse.Accounts)
	assert.NotEmpty(t, accountsResponse.Count)
}
