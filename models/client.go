package models

import (
	"github.com/jinzhu/gorm"
)

// Client struct
type Client struct {
	gorm.Model
	Key    string `json:"key" gorm:"column:key;type:varchar(255);UNIQUE;NOT NULL"`
	Secret string `json:"secret" gorm:"column:secret;type:varchar(255)"`
}
