package router

import (
	// "zurich-pos-gin/apis"
	// "zurich-pos-gin/middleware"

	"test-go/controllers"
	"test-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetAccountsRoutes(e *gin.RouterGroup) {
	// accountsRepo := apis.AccountsMigration(db)
	controllers.AccountsAutoMigrate()
	accounts := e.Group("/accounts")
	{
		accounts.POST("/login", controllers.Login)

		accounts.Use(middleware.VerifyToken)
		// after login or authentication
		accounts.POST("/logout", controllers.Logout)
		accounts.GET("", controllers.GetAccounts)
		accounts.GET("/:id", controllers.GetAccountById)
		accounts.POST("", controllers.CreateAccount)
		accounts.PUT("/:id", controllers.UpdateAccount)
		accounts.DELETE("/:id", controllers.DeleteAccount)
		accounts.PUT("/:id/change_password", controllers.UpdateAccountPassword)
	}
}
