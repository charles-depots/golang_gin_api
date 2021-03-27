package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-api/config"
	"golang-gin-api/internal/api/status"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// Create a rate limiter with a maximum number of operations to perform per second.
func RateLimitMiddleware() gin.HandlerFunc {
	rateCfg := config.GetConfig().Rate
	r := rate.Every(time.Duration(rateCfg.Limit) * time.Second)
	limit := rate.NewLimiter(r, rateCfg.Burst)
	return func(c *gin.Context) {
		if !limit.Allow() {
			fmt.Println("The current flow is too big, please try again later")
			c.JSON(http.StatusOK, gin.H{
				"status":  status.RequestRateLimit,
				"message": status.GetStatusMessage(status.RequestRateLimit),
			})
			c.Abort()
			return
		}
	}
}
