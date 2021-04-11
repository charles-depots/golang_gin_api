package user

import (
	"context"
	"golang-gin-api/internal/api/model/user/dbinit"
	"golang-gin-api/pkg/db/mysql"
)

type userDomain struct{}

func NewUserDomain() *userDomain {
	return &userDomain{}
}

func (d *userDomain) UserDb(ctx context.Context) dbinit.UserDbInterface {
	return &userDb{mysql.GetDB(ctx)}
}
