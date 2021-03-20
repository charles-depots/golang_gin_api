package main

import (
	"fmt"
	"go.uber.org/zap"
	"golang-gin-api/config"
	"golang-gin-api/internal/api/router"
	"golang-gin-api/pkg/db"
	"golang-gin-api/pkg/logger"
)

func init() {

}

func main() {
	// Initialize zap logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s", config.ProjectName())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(config.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()

	// Initialize the database connection
	dbErr := db.InitMySQLCon()
	if dbErr != nil {
		loggers.Fatal("new db err", zap.Error(dbErr))
	}

	router.InitHttpServer(loggers)
}
