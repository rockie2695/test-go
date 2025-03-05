package router

import (
	"test-go/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "test-go/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	SetCors(router)
	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{

		SetAccountsRoutes(v1, db) // set authentication
		router.Use(middleware.VerifyToken)
		SetCustomersRoutes(v1, db)
	}

	return router
}
