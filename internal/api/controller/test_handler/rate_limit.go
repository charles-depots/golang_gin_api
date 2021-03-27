package test_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Test an interface that rate limit for request
func (h *handler) RateLimitTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success",
	})
}
