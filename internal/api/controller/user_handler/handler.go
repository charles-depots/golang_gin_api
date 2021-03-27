package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"golang-gin-api/internal/api/service/user_service"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	GetUserInfo(c *gin.Context)
}

type handler struct {
	logger      *zap.Logger
	redis       *redis.Client
	userService user_service.UserService
}

func New(logger *zap.Logger, redis *redis.Client) Handler {
	return &handler{
		logger:      logger,
		redis:       redis,
		userService: user_service.NewUserService(logger),
	}
}
