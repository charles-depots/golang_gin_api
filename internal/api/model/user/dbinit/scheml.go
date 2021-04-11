package dbinit

import (
	"context"
	"time"
)

type UserDomainInterface interface {
	UserDb(ctx context.Context) UserDbInterface
}

type User struct {
	Id        int32  `gorm:"AUTO_INCREMENT"`
	Name      string `json:"name"`
	Pwd       string `json:"password"`
	Phone     string `json:"phone`
	Email     string `gorm:"type:varchar(35);unique_index;"`
	CreatedAt *time.Time
	UpdateTAt *time.Time
}

type UserDbInterface interface {
	Create(username, pwd string, phone string, email string) error
	Get(id string) (*User, error)
	List(query *User, offset, limit int) ([]*User, error)
	Update(query *User) (*User, error)
	Delete(query *User) error
	CheckUserByName(username string) (bool, *User, error)
}
