package router

import (
	// "zurich-pos-gin/apis"
	// "zurich-pos-gin/middleware"

	"test-go/controllers"

	"github.com/gin-gonic/gin"
)

func SetAccountsRoutes(e *gin.Engine) {
	// accountsRepo := apis.AccountsMigration(db)
	controllers.AccountsAutoMigrate()
	accounts := e.Group("/accounts")
	{
		accounts.POST("/login", controllers.Login)
		accounts.GET("", controllers.GetAccounts)
		accounts.GET("/:id", controllers.GetAccountById)
		accounts.POST("", controllers.CreateAccount)
		accounts.PUT("/:id", controllers.UpdateAccount)
		accounts.DELETE("/:id", controllers.DeleteAccount)
	}
}
