package router

import (
	"test-go/controllers"
	// "test-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetCustomersRoutes(e *gin.RouterGroup) {
	// accountsRepo := apis.AccountsMigration(db)
	controllers.CustomersAutoMigrate()
	customers := e.Group("/customers")
	{
		customers.GET("/", controllers.GetCustomers)
		customers.GET("/:id", controllers.GetCustomerById)
		customers.POST("/", controllers.CreateCustomer)
		customers.PUT("/:id", controllers.UpdateCustomer)
		customers.DELETE("/:id", controllers.DeleteCustomer)
	}
}
