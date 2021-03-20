package user_service

import "go.uber.org/zap"

var _ UserService = (*userSer)(nil)

type UserService interface {
	Create(userRes *RegisterInfo) error
}

type userSer struct {
}

func NewUserService(logger *zap.Logger) UserService {
	return &userSer{}
}
