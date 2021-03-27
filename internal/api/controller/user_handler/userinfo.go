package user_handler

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	md "golang-gin-api/internal/api/router/middleware"
	"net/http"
)

// Test an interface that requires certification
func (h *handler) GetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*md.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token is valid",
			"data":   claims,
		})
	}
}
