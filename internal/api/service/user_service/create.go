package user_service

import (
	"golang-gin-api/internal/api/model/user"
)

// User registration information
type RegisterInfo struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Email string `json:"email"`
}

func (u *UserSer) Create(userRes *RegisterInfo) error {
	userModel := user.NewModel()
	userModel.Name = userRes.Name
	userModel.Phone = userRes.Phone
	userModel.Pwd = userRes.Pwd
	userModel.Email = userRes.Email

	err := userModel.Insert()
	if err != nil {
		return err
	}

	return nil
}
