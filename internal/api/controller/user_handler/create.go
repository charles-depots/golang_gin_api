package user_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-api/internal/api/service/user_service"
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
func (h *handler) RegisterUser(c *gin.Context)  {
	var registerInfo = new(RegisterInfo)
	bindErr := c.BindJSON(&registerInfo)
	if bindErr == nil {
		createUserData := new(user_service.RegisterInfo)
		createUserData.Name = registerInfo.Name
		createUserData.Pwd = registerInfo.Pwd
		createUserData.Phone = registerInfo.Phone
		createUserData.Email = registerInfo.Email
		var x = &h.userService
		fmt.Printf("%v",x)

		//if errNow == nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"status": 0,
		//		"msg":    "success ",
		//		"data":   nil,
		//	})
		//}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户注册解析数据失败" + bindErr.Error(),
			"data":   nil,
		})
	}
}
