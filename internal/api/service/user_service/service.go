package user_service

import (
	"go.uber.org/zap"
	"golang-gin-api/internal/api/model/user"
)

var _ UserService = (*UserSer)(nil)

type UserService interface {
	Create(userRes *RegisterInfo) error
	LoginCheck(userRes *LoginReq) (bool, user.User, error)
}

type UserSer struct {
	logger *zap.Logger
}

func NewUserService(logger *zap.Logger) UserService {
	return &UserSer{}
}
