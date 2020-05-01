package server

import (
	"os"
	"pando/server/router"

	"github.com/labstack/echo/v4/middleware"
)

// Start starts worker
func Start() {
	e := router.NewRouter()
	if os.Getenv("PANDO_CORS") == "true" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"*"},
		}))
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PANDO_PORT")))
}
