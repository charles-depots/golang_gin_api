package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	testHandler "golang-gin-api/internal/api/controller/test_handler"
	userHandler "golang-gin-api/internal/api/controller/user_handler"
	"golang-gin-api/internal/api/model/user"
	md "golang-gin-api/internal/api/router/middleware"
	"golang-gin-api/pkg/rabbitmq"
)

func InitHttpServer(logger *zap.Logger, redis *redis.Client, rabbitmq *rabbitmq.RabbitMQ) {
	if logger == nil {
		return
	}

	// Initialize go instance
	router := gin.Default()

	// Test controller
	middleRateLimit := md.RateLimitMiddleware()
	testHandler := testHandler.New()
	test := router.Group("/V1/internal/test", middleRateLimit)
	{
		test.GET("/rate-limit", testHandler.RateLimitTest) // Testing rate limit for request
	}

	// Init user model
	user.InitModel()

	// JWT middle
	middlesJWT := md.JWTAuth(logger)

	// User controller
	userHandler := userHandler.New(logger, redis)
	customer := router.Group("/V1/internal/customer")
	{
		customer.POST("/register", userHandler.RegisterUser)
		customer.POST("/token", userHandler.Login)
	}

	// secure v1
	sv1 := router.Group("/V1/internal/auth", middlesJWT)
	{
		sv1.GET("/userinfo", userHandler.GetUserInfo)
	}

	router.Run(":8080")
}
