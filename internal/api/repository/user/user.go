package user

import (
	"fmt"
	"golang-gin-api/internal/api/model/user/dbinit"
	"golang-gin-api/pkg/db/mysql"
	"gorm.io/gorm"
)

func init() {
	mysql.RegisterInjector(func(db *gorm.DB) {
		mysql.SetupTableModel(db, &dbinit.User{})
	})
}

type userDb struct {
	db *gorm.DB
}

func (u userDb) Create(username, pwd string, phone string, email string) error {
	fmt.Println(username, pwd, phone, email)
	dbResult := u.db.Where("name = ?", username).Find(dbinit.User{})
	if dbResult.Error != nil {
		return fmt.Errorf("%s", "used to already exist, please log in directly.")
	}
	user := dbinit.User{
		Name:  username,
		Pwd:   pwd,
		Phone: phone,
		Email: email,
	}
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userDb) Get(id string) (*dbinit.User, error) {
	var r dbinit.User
	err := u.db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (u userDb) List(query *dbinit.User, offset, limit int) ([]*dbinit.User, error) {
	var r []*dbinit.User

	db := mysql.WithOffsetLimit(u.db, offset, limit)

	err := db.Where(query).Find(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u userDb) Update(query *dbinit.User) (*dbinit.User, error) {
	err := u.db.Updates(query).Error
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (u userDb) Delete(query *dbinit.User) error {
	err := u.db.Where(query).Delete(&dbinit.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userDb) CheckUserByName(username string) (bool, *dbinit.User, error) {
	userData := dbinit.User{}
	userExist := false

	var user dbinit.User
	dbResult := u.db.Where("name = ?", username).Find(&user)
	if dbResult.Error != nil {
		return userExist, &userData, dbResult.Error
	} else {
		userExist = true
	}

	return userExist, &user, nil
}
