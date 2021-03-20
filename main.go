package main

import (
	"golang-gin-api/internal/api/router"
	"golang-gin-api/pkg/db"
)

func init() {
	// Initialize the database connection
	dbErr := db.InitMySQLCon()
	if dbErr != nil {
		panic(dbErr)
	}
}

func main() {
	router.InitHttpServer()
}
