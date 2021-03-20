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

func (u *userSer) Create(userRes *RegisterInfo) error {
	model := user.NewModel()
	model.Name = userRes.Name
	model.Phone = userRes.Phone
	model.Pwd = userRes.Pwd
	model.Email = userRes.Email

	err := model.Insert()

	return err
}
