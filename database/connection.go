package database

import (
	"fmt"
	"os"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connect()
	return Db
}

func connect() *gorm.DB {
	var err error
	// dsn := getMysqlConnection() + "?parseTime=true&loc=Local"
	dsn := getPGConnection()
	fmt.Println("dsn : ", dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}

// func getMysqlConnection() string {
// 	network := os.Getenv("network")
// 	host := os.Getenv("host")
// 	port := os.Getenv("port")
// 	username := os.Getenv("username")
// 	password := os.Getenv("password")
// 	dbName := os.Getenv("db")

// 	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, host, port, dbName)

// 	return connection
// }

func getPGConnection() string {
	username := os.Getenv("pg_username")
	password := os.Getenv("pg_password")
	host := os.Getenv("pg_host")
	dbName := os.Getenv("pg_db")

	// connection := "host=localhost user=gorm password=gorm dbname=gorm port=9920"
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s TimeZone=Asia/Shanghai", host, username, password, dbName)
	return connection
}
