package main

import (
	"test-go/config"
	"test-go/database"
	"test-go/router"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server of go server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9002
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description Type "Bearer" followed by a space and JWT token.

func main() {
	config.EnvSetup()
	db := database.InitDb()
	router := router.InitRouter()
	// router.Run()
	router.Run(":9002")

	print(db)
}
