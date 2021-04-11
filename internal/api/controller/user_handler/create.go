package user_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-gin-api/internal/api/model/user/dbinit"
	"golang-gin-api/internal/api/status"
	"net/http"
)

type RegisterInfo struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Email string `json:"email"`
}

// User registration interface
func (h *handler) RegisterUser(ctx *gin.Context, req *http.Request) {
	var protoReq dbinit.User

	if err := req.ParseForm(); err != nil {
		fmt.Println(err)
	}
	errUserCreate := h.userService.CreateUser(ctx, &protoReq)

	if errUserCreate != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": status.UserCreateErrorCode,
			"msg":    status.GetStatusMessage(status.UserCreateErrorCode) + errUserCreate.Error(),
			"data":   nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": status.UserCreateCode,
			"msg":    status.GetStatusMessage(status.UserCreateCode),
			"data":   nil,
		})
	}
}
