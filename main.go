package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-api/internal/api/domain/user"
	md "golang-gin-api/internal/api/router/middleware"
	userHandle "golang-gin-api/internal/api/service/user"
	"golang-gin-api/pkg/db"
)

func main() {
	// 初始化db
	dbErr := db.InitMySQLCon()
	if dbErr != nil {
		panic(dbErr)
	}
	user.InitModel()

	// 初始化Gin实例
	router := gin.Default()
	v1 := router.Group("/V1/internal")
	{
		v1.POST("/customer/register", userHandle.RegisterUser)
		v1.POST("/customer/token", userHandle.Login)
	}

	// secure v1
	sv1 := router.Group("/V1/internal/auth")
	sv1.Use(md.JWTAuth())
	{
		sv1.GET("/validate", userHandle.GetValidate)
	}
	router.Run(":8080")
}
