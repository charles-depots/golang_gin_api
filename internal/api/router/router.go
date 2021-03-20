package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	userHandle "golang-gin-api/internal/api/controller/user_handler"
	"golang-gin-api/internal/api/model/user"
	md "golang-gin-api/internal/api/router/middleware"
)

func InitHttpServer(logger *zap.Logger) {
	if logger == nil {
		return
	}

	// Initialize go instance
	router := gin.Default()
	user.InitModel()

	// JWT middle
	middlesJWT := md.JWTAuth(logger)

	v1 := router.Group("/V1/internal")
	{
		v1.POST("/customer/register", userHandle.RegisterUser)
		v1.POST("/customer/token", userHandle.Login)
	}

	// secure v1
	sv1 := router.Group("/V1/internal/auth", middlesJWT)
	{
		sv1.GET("/validate", userHandle.GetValidate)
	}
	router.Run(":8080")
}
