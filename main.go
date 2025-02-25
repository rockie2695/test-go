package main

import (
	"test-go/config"
	"test-go/database"
	"test-go/router"
)

func main() {
	config.EnvSetup()
	db := database.InitDb()
	router := router.InitRouter()
	// router.Run()
	router.Run(":9002")

	print(db)
}
