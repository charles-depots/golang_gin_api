package user_handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-gin-api/internal/api/service/user_service"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i 为了避免被其他包实现
	i()
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	GetValidate(c *gin.Context)
}

type handler struct {
	logger      *zap.Logger
	userService user_service.UserService
}

func New(logger *zap.Logger) Handler {
	return &handler{logger: logger}
}

func (h *handler) i() {}
