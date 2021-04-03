package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	mqHandler "golang-gin-api/internal/api/controller/queue_handler"
	testHandler "golang-gin-api/internal/api/controller/test_handler"
	userHandler "golang-gin-api/internal/api/controller/user_handler"
	"golang-gin-api/internal/api/model/user"
	md "golang-gin-api/internal/api/router/middleware"
)

func InitHttpServer(logger *zap.Logger, redis *redis.Client) {
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

	// RabbitMQ controller
	mqCtl := mqHandler.MqHandler()
	mq := router.Group("/V1/internal/mq", middleRateLimit)
	{
		mq.POST("/exchange", func(c *gin.Context) {
			mqCtl.ExchangeHandler(c.Writer, c.Request)
		})
		mq.POST("/queue/bind", func(c *gin.Context) {
			mqCtl.QueueBindHandler(c.Writer, c.Request)
		})
		//consume queue
		mq.GET("/queue", func(c *gin.Context) {
			mqCtl.QueueHandler(c.Writer, c.Request)
		})
		//declare queue
		mq.POST("/queue", func(c *gin.Context) {
			mqCtl.QueueHandler(c.Writer, c.Request)
		})
		//delete queue
		mq.DELETE("/queue", func(c *gin.Context) {
			mqCtl.QueueHandler(c.Writer, c.Request)
		})
		mq.POST("/publish", func(c *gin.Context) {
			mqCtl.PublishHandler(c.Writer, c.Request)
		})
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
