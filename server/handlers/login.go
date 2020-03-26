package handlers

import (
	"fmt"
	"net/http"
	"os"
	"pando/db"
	"pando/models"
	"pando/server/helper"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type (
	// LoginPayload request body model
	LoginPayload struct {
		Username string `json:"username" validate:"required" example:"user123"`
		Password string `json:"password" validate:"required" example:"pass123"`
	}
	// LoginResponse response model
	LoginResponse struct {
		helper.EchoResp
		Details LoginResponseDetails
	}
	// LoginResponseDetails detail of response
	LoginResponseDetails struct {
		ExpiresIn int    `json:"expires_in" example:"3600"`
		Token     string `json:"token" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJle..."`
	}
)

// Login user using username and password
// @Summary Login user
// @Description login returns token and expire time in seconds
// @Tags Client
// @Accept  json
// @Produce  json
// @Security BasicAuth
// @Param RequestBody body LoginPayload true "JSON request body"
// @Router /client/login [post]
// @Success 200 {object} LoginResponse
// @Failure 404 {object} helper.EchoResp
// @Failure 500 {object} helper.EchoResp
func Login(c echo.Context) error {
	defer c.Request().Body.Close()

	body := new(LoginPayload)

	if err := c.Bind(body); err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0002", "Request body bind error", map[string]interface{}{
			"error": fmt.Sprintf("%+v", err),
		})
	}
	if err := c.Validate(body); err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0001", "Request body validation error", map[string]interface{}{
			"error": fmt.Sprintf("%+v", err),
		})
	}

	user := models.User{}
	err := db.DB.Model(&models.User{}).Find(&user, models.User{Username: body.Username}).Error
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusNotFound, "0001", "Username not found", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0001", "Wrong password", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	expiresIn, err := strconv.ParseInt(os.Getenv("PANDO_JWT_EXPIRES"), 10, 64)
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0002", "Parsing int error", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	token, err := helper.GenerateJWTtoken(fmt.Sprint(user.ID))
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0002", "Generate token error", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	return helper.ReturnJSONresp(c, http.StatusOK, "0000", "Success", map[string]interface{}{
		"expires_in": expiresIn,
		"token":      token,
	})
}
