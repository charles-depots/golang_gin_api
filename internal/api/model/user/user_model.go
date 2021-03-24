package user

import (
	mysql "golang-gin-api/pkg/db"
	"time"
)

// Construct user table
type User struct {
	Id        int32  `gorm:"AUTO_INCREMENT"`
	Name      string `json:"name"`
	Pwd       string `json:"password"`
	Phone     string `json:"phone`
	Email     string `gorm:"type:varchar(35);unique_index;"`
	CreatedAt *time.Time
	UpdateTAt *time.Time
}

// Initialize the structure and definition of the user table
func InitModel() {
	mysql.DB.AutoMigrate(&User{})
}

func NewModel() *User {
	return new(User)
}
