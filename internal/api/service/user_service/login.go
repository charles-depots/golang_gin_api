package user_service

import (
	"fmt"
	"golang-gin-api/internal/api/model/user"
)

// User login request parameters
type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}

// Login information verification
func (u *UserSer) LoginCheck(userRes *LoginReq) (bool, user.User, error) {
	userData := user.User{}
	userExist, userInfo, _ := user.CheckUserByName(userRes.Name)
	if !userExist {
		fmt.Println("Sorry, user is not exist")
	}

	if userRes.Name == userInfo.Name && userRes.Pwd == userInfo.Pwd {
		userExist = true
		userData.Name = userInfo.Name
		userData.Email = userInfo.Email
	}

	if !userExist {
		return userExist, userData, fmt.Errorf("%s", "The login information is wrong")
	}

	return userExist, userData, nil
}
