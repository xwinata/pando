package helper

import (
	"log"
	"net/http"
	"os"
	"pando/db"
	"pando/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// GenerateJWTtoken generates jwt token
func GenerateJWTtoken(id string) (string, error) {
	expiresIn, err := strconv.Atoi(os.Getenv("PANDO_JWT_EXPIRES"))
	if err != nil {
		return "", err
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
		Id:        id,
	})

	token, err := rawToken.SignedString([]byte(os.Getenv("PANDO_JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateClient check client from client table in database
func ValidateClient(key, secret string, c echo.Context) (bool, error) {
	client := models.Client{Key: key}
	err := db.DB.First(&client, client).Error
	if err != nil {
		log.Println(err)
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret))
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

// ValidateUserPermission check user permission by matching routes name and user's role permission
func ValidateUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken := c.Get("user").(*jwt.Token)
		tokenClaims := userToken.Claims.(jwt.MapClaims)
		userID, _ := strconv.ParseUint(tokenClaims["jti"].(string), 10, 64)

		rolePermission := models.RolePermission{}
		err := db.DB.Table("role_permissions").
			Joins("INNER JOIN permissions ON permissions.id = role_permissions.permission_id").
			Joins("INNER JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
			Where("permissions.path = ?", c.Path()).
			Where("user_roles.user_id = ?", userID).
			First(&rolePermission).
			Error
		if err != nil {
			return ReturnJSONresp(c, http.StatusForbidden, "0003", "User don't have permission access", nil)
		}
		return next(c)
	}
}
