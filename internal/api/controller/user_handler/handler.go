package user_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"golang-gin-api/internal/api/repository/user"
	"golang-gin-api/internal/api/service/user_service"
	"net/http"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	RegisterUser(ctx *gin.Context, req *http.Request)
	Login(ctx *gin.Context, req *http.Request)
	GetUserInfo(ctx *gin.Context)
}

type handler struct {
	logger      *zap.Logger
	redis       *redis.Client
	userService user_service.UserServiceServer
}

func New(logger *zap.Logger, redis *redis.Client) Handler {
	return &handler{
		logger:      logger,
		redis:       redis,
		userService: user_service.NewUserService(user.NewUserDomain()),
	}
}
