package router

import (
	"os"
	"pando/server/handlers"
	"pando/server/helper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ClientGroup group of client's endpoints
func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	g.Use(middleware.BasicAuth(helper.ValidateClient))

	g.POST("/login", handlers.Login).Name = "client-login"
}

// UserGroup group of authenticated user's endpoints
func UserGroup(e *echo.Echo) {
	g := e.Group("/user")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("PANDO_JWT_SECRET")),
	}))

	g.Use(helper.ValidateUserPermission)

	g.GET("/user", handlers.ViewCurrentUser).Name = "user-user"
}
