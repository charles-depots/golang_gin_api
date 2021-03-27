package test_handler

import (
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	RateLimitTest(c *gin.Context)
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
