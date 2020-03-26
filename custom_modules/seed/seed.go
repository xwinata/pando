package seed

import (
	"pando/db"
	"pando/models"

	"golang.org/x/crypto/bcrypt"
)

// Seed inserts dummy datas
func Seed() {
	clientsecret, _ := bcrypt.GenerateFromPassword([]byte("clientsecret"), bcrypt.DefaultCost)
	client := models.Client{
		Key:    "clientkey",
		Secret: string(clientsecret),
	}
	db.DB.Create(&client)

	adminpass, _ := bcrypt.GenerateFromPassword([]byte("rahasia123"), bcrypt.DefaultCost)
	admin := models.User{
		Username: "admin",
		Password: string(adminpass),
	}
	db.DB.Create(&admin)

	adminRole := models.Role{
		Name: "Admin",
	}
	db.DB.Create(&adminRole)

	adminUserRole := models.UserRole{
		UserID: admin.ID,
		RoleID: adminRole.ID,
	}
	db.DB.Create(&adminUserRole)

	permission := models.Permission{
		Method: "ALL",
		Path:   "all",
		Name:   "all",
	}
	db.DB.Create(&permission)

	rolePermission := models.RolePermission{
		RoleID:       adminRole.ID,
		PermissionID: permission.ID,
	}
	db.DB.Create(&rolePermission)
}

// Unseed removes dummy datas
func Unseed() {
	seedUser := models.User{}
	seedUserRoles := []models.UserRole{}
	db.DB.Model(&models.User{}).Find(&seedUser, "username = ?", "admin")
	db.DB.Model(&seedUser).Related(&seedUserRoles)

	for _, v := range seedUserRoles {
		db.DB.Unscoped().Delete(models.RolePermission{}, "role_id = ?", v.RoleID)
		db.DB.Unscoped().Delete(v)
	}
	db.DB.Unscoped().Delete(models.Permission{}, "value = ?", "all")
	db.DB.Unscoped().Delete(models.Role{}, "name = ?", "Admin")
	db.DB.Unscoped().Delete(models.Client{}, "key = ?", "clientkey")
	db.DB.Unscoped().Delete(models.User{}, "username = ?", "admin")
}
