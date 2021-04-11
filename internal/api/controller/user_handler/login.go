package user_handler

import (
	_ "fmt"
	"golang-gin-api/internal/api/model/user/dbinit"
	md "golang-gin-api/internal/api/router/middleware"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginResult struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}

func (h *handler) Login(ctx *gin.Context, req *http.Request) {
	var protoReq dbinit.User
	isPass, user, err := h.userService.LoginCheck(ctx, &protoReq)
	if isPass {
		generateToken(h, ctx, user)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "verification failed: " + err.Error(),
			"data":   nil,
		})
	}
}

// Token generator
func generateToken(h *handler, c *gin.Context, user dbinit.User) {
	// Construct SignKey: A value is required for signing and de-signing
	j := md.NewJWT()

	// Construct user claims information (load)
	claims := md.CustomClaims{
		user.Name,
		user.Email,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // Signature effective time
			ExpiresAt: int64(time.Now().Unix() + 3600), // Signature expiration time
			Issuer:    "charles.101",                   // Signature issuer
		},
	}
	// Generate token objects based on claims
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
	h.redis.Set(user.Email, token, 10*time.Second)

	// Get user related data
	data := LoginResult{
		Name:  user.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "Successful landing",
		"data":   data,
	})
	return
}
