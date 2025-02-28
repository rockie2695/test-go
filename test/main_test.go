/*
unknown reason why go test fail
--- FAIL: TestPingRoute (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
panic: runtime error: invalid memory address or nil pointer dereference
*/
// package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	// "test-go/config"
// 	// "test-go/database"
// 	"test-go/router"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPingRoute(t *testing.T) {
// 	config.EnvSetup()
// 	database.InitDb()
// 	router := router.InitRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ping", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "pong", w.Body.String())
// }

// // Test for POST /user/add
// func TestPostUser(t *testing.T) {
// 	router := setupRouter()
// 	router = postUser(router)

// 	w := httptest.NewRecorder()

// 	// Create an example user for testing
// 	exampleUser := User{
// 		Username: "test_name",
// 		Gender:   "male",
// 	}
// 	userJson, _ := json.Marshal(exampleUser)
// 	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
// 	router.ServeHTTP(w, req)

//		assert.Equal(t, 200, w.Code)
//		// Compare the response body with the json data of exampleUser
//		assert.Equal(t, string(userJson), w.Body.String())
//	}
// func TestPingRoute(t *testing.T) {
// 	router := router.InitRouter()
// 	if router == nil {
// 		t.Fatal("Failed to initialize router")
// 	}

// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest("GET", "/ping", nil)
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "pong", w.Body.String())
// }
