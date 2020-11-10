package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var maria *MariaDB

func main() {
	userName := "root"
	password := "1234"
	host := "localhost"
	database := "test"

	maria = new(MariaDB)
	err := maria.Connect(userName, password, host, database)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}
	defer maria.Close()

	maria.Migrate(testMigrations)

	router = gin.Default()
	SetUpRouter(router)
	router.Run(":8080")
}
