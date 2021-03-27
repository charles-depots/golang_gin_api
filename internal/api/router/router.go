package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	userHandler "golang-gin-api/internal/api/controller/user_handler"
	"golang-gin-api/internal/api/model/user"
	md "golang-gin-api/internal/api/router/middleware"
	"github.com/go-redis/redis/v7"
)

func InitHttpServer(logger *zap.Logger, redis *redis.Client) {
	if logger == nil {
		return
	}

	// Initialize go instance
	router := gin.Default()
	user.InitModel()

	// JWT middle
	middlesJWT := md.JWTAuth(logger)

	// User controller
	userHandler := userHandler.New(logger, redis)

	v1 := router.Group("/V1/internal")
	{
		v1.POST("/customer/register", userHandler.RegisterUser)
		v1.POST("/customer/token", userHandler.Login)
	}

	// secure v1
	sv1 := router.Group("/V1/internal/auth", middlesJWT)
	{
		sv1.GET("/userinfo", userHandler.GetUserInfo)
	}
	router.Run(":8080")
}
