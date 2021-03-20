package user_handler

import (
	"github.com/gin-gonic/gin"
	"golang-gin-api/internal/api/model/user"
	"net/http"
)

// User registration information
// Note: Registration information can use Gin's internal verification tool
type RegisterInfo struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Email string `json:"email"`
}

// User registration interface
func RegisterUser(c *gin.Context) {
	var registerInfo RegisterInfo
	bindErr := c.BindJSON(&registerInfo)
	if bindErr == nil {
		// User registration
		err := user.Register(registerInfo.Name, registerInfo.Pwd, registerInfo.Phone, registerInfo.Email)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "success ",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "Registration failed: " + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "Failed to resolve user registration data: " + bindErr.Error(),
			"data":   nil,
		})
	}
}
