package router

import (
	"github.com/gin-gonic/gin"
	// "log" // Import the log package
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	log.Println("Ping route hit") // Use log.Println instead of t.Log
	// 	c.String(200, "pong")
	// })
	SetCors(router)
	SetAccountsRoutes(router) // set authentication
	// router.Use(middleware.VerifyToken)
	return router
}
