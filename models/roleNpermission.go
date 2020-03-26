package models

import "github.com/jinzhu/gorm"

type (
	// Role struct
	Role struct {
		gorm.Model
		Name            string           `json:"name" gorm:"column:name;type:varchar(255);UNIQUE;NOT NULL"`
		UserRoles       []UserRole       `gorm:"foreignkey:RoleID"`
		RolePermissions []RolePermission `gorm:"foreignkey:RoleID"`
	}
	// UserRole struct
	UserRole struct {
		gorm.Model
		UserID uint `json:"user_id" gorm:"column:user_id"`
		RoleID uint `json:"role_id" gorm:"column:role_id"`
	}
	// Permission struct
	Permission struct {
		gorm.Model
		Method          string           `json:"method" gorm:"column:method;type:varchar(255)"`
		Path            string           `json:"path" gorm:"column:path;type:varchar(255)"`
		Name            string           `json:"name" gorm:"column:name;type:varchar(255)"`
		RolePermissions []RolePermission `gorm:"foreignkey:PermissionID"`
	}
	// RolePermission struct
	RolePermission struct {
		gorm.Model
		RoleID       uint `json:"role_id" gorm:"column:role_id"`
		PermissionID uint `json:"permission_id" gorm:"column:permission_id"`
	}
)
