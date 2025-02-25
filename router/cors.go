package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCors(e *gin.Engine) {

	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:2096", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:2096"
		},
		MaxAge: 12 * time.Hour,
	}))

}
