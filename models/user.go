package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"column:username;type:varchar(255);UNIQUE;NOT NULL"`
	Password  string     `json:"password" gorm:"column:password;type:varchar(255)"`
	UserRoles []UserRole `gorm:"foreignkey:UserID"`
}
