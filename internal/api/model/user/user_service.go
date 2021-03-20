package user

import (
	"fmt"
	mysql "golang-gin-api/pkg/db"
)

// User login request parameters
type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}

func NewModel() *User {
	return new(User)
}

// Insert data into the user table
func (user *User) Insert() error {
	return mysql.DB.Model(&User{}).Create(&user).Error
}

// Realization of user registration function
func Register(username, pwd string, phone string, email string) error {
	fmt.Println(username, pwd, phone, email)

	if CheckUser(username) {
		return fmt.Errorf("%s", "used to already exist, please log in directly.")
	}

	// Construct user registration information
	user := User{
		Name:  username,
		Pwd:   pwd,
		Phone: phone,
		Email: email,
	}
	insertErr := user.Insert()
	return insertErr

}

// Check user information
func CheckUser(username string) bool {

	result := false
	var user User

	dbResult := mysql.DB.Where("name = ?", username).Find(&user)
	if dbResult.Error != nil {
		fmt.Printf("Failed to obtain user information:%v\n", dbResult.Error)
	} else {
		result = true
	}
	return result
}

// Login information verification
func LoginCheck(login LoginReq) (bool, User, error) {
	userData := User{}
	userExist := false

	var user User
	dbErr := mysql.DB.Where("name = ?", login.Name).Find(&user).Error

	if dbErr != nil {
		return userExist, userData, dbErr
	}
	if login.Name == user.Name && login.Pwd == user.Pwd {
		userExist = true
		userData.Name = user.Name
		userData.Email = user.Email
	}

	if !userExist {
		return userExist, userData, fmt.Errorf("%s", "The login information is wrong")
	}
	return userExist, userData, nil
}
