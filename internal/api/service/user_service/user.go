package user_service

import (
	"context"
	"fmt"
	"golang-gin-api/internal/api/model/user/dbinit"
)

var _ UserServiceServer = (*UserService)(nil)

type UserServiceServer interface {
	CreateUser(ctx context.Context, in *dbinit.User) error
	LoginCheck(ctx context.Context, in *dbinit.User) (bool, dbinit.User, error)
}

type UserService struct {
	userDomain dbinit.UserDomainInterface
}

func NewUserService(userDomain dbinit.UserDomainInterface) *UserService {
	return &UserService{
		userDomain: userDomain,
	}
}

func (s *UserService) CreateUser(ctx context.Context, in *dbinit.User) error {
	err := s.userDomain.UserDb(ctx).Create(in.Name, in.Pwd, in.Phone, in.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginCheck(ctx context.Context, in *dbinit.User) (bool, dbinit.User, error) {
	userData := dbinit.User{}
	userExist, userInfo, _ := s.userDomain.UserDb(ctx).CheckUserByName(in.Name)
	if !userExist {
		fmt.Println("Sorry, user is not exist")
		return userExist, userData, fmt.Errorf("%s", "The login information is wrong")
	}

	if in.Name == userInfo.Name && in.Pwd == userInfo.Pwd {
		userExist = true
		userData.Name = userInfo.Name
		userData.Email = userInfo.Email
	}

	return userExist, userData, nil
}
