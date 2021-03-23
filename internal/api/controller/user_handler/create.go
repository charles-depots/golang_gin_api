package user_handler

import (
	"github.com/gin-gonic/gin"
	"golang-gin-api/internal/api/code"
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
func (h *handler) RegisterUser(c *gin.Context) {
	var registerInfo = new(RegisterInfo)
	bindErr := c.BindJSON(&registerInfo)
	if bindErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code.UserParamErrorCode,
			"message": code.GetText(code.UserParamErrorCode) + bindErr.Error(),
			"data":    nil,
		})
	}
	createUserData := new(user_service.RegisterInfo)
	createUserData.Name = registerInfo.Name
	createUserData.Pwd = registerInfo.Pwd
	createUserData.Phone = registerInfo.Phone
	createUserData.Email = registerInfo.Email
	errUserCreate := h.userService.Create(createUserData)

	if errUserCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": code.UserCreateErrorCode,
			"msg":    code.GetText(code.UserCreateErrorCode) + errUserCreate.Error(),
			"data":   nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": code.UserCreateCode,
			"msg":    code.GetText(code.UserCreateCode),
			"data":   nil,
		})
	}
}
