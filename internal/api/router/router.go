package router

import (
	"github.com/gin-gonic/gin"
	userHandle "golang-gin-api/internal/api/controller/user_handler"
	"golang-gin-api/internal/api/model/user"
	md "golang-gin-api/internal/api/router/middleware"
)

func InitHttpServer() {
	// Initialize go instance
	router := gin.Default()
	user.InitModel()
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
