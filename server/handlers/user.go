package handlers

import (
	"fmt"
	"net/http"
	"pando/db"
	"pando/models"
	"pando/server/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// ViewCurrentUserResponse struct response for function ViewCurrentUser
type ViewCurrentUserResponse struct {
	Username string `json:"username"`
}

// ViewCurrentUser shows current user's datas
// @Summary View current user datas
// @Tags User
// @Accept  json
// @Produce  json
// @Security OAuth2Password
// @Router /user/user [get]
// @Success 200 {object} ViewCurrentUserResponse
// @Failure 403 {object} helper.EchoResp
// @Failure 404 {object} helper.EchoResp
// @Failure 500 {object} helper.EchoResp
func ViewCurrentUser(c echo.Context) error {
	defer c.Request().Body.Close()

	userToken := c.Get("user").(*jwt.Token)
	tokenClaims := userToken.Claims.(jwt.MapClaims)
	userID, _ := strconv.ParseUint(tokenClaims["jti"].(string), 10, 64)

	user := models.User{}
	err := db.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusUnauthorized, "0002", "Error finding user. please re-login", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	return helper.ReturnJSONresp(c, http.StatusOK, "0000", "Success", ViewCurrentUserResponse{Username: user.Username})
}
