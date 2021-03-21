package user_service

import "go.uber.org/zap"

var _ UserService = (*UserSer)(nil)

type UserService interface {
	Create(userRes *RegisterInfo) error
}

type UserSer struct {
	logger *zap.Logger
}

func NewUserService(logger *zap.Logger) UserService {
	return &UserSer{}
}
