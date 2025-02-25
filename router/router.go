package router

import (
	// "net/http"
	// "log/slog"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// log := slog.Default()
	SetCors(router)

	SetAccountsRoutes(router) // set authentication
	// router.Use(middleware.VerifyToken)
	return router
}
