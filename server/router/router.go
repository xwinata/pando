package router

import (
	"flag"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"pando/db"
	"pando/models"
	"pando/server/handlers"

	echoSwagger "github.com/swaggo/echo-swagger"

	// import swagger api documentation drive
	_ "pando/docs"
)

// CustomValidator type for json body validation
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates request body
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewRouter swagger doc
// @title Pando API
// @version 1.0
// @description This is a sample server Petstore server.
// @contact.name richstain
// @contact.email richstain2u@gmail.com
// @BasePath /
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/testTaskDispatch", handlers.TestQueueAtask)

	// url rewrite
	url := strings.Split(os.Getenv("PANDO_URL_REWRITE"), ",")
	x := make(map[string]string)
	for _, v := range url {
		if len(v) > 0 {
			x[v] = "/$1"
		}
	}
	if len(x) > 0 {
		e.Pre(middleware.Rewrite(x))
	}

	e.Validator = &CustomValidator{validator: validator.New()}

	ClientGroup(e)
	UserGroup(e)

	generateRoutePermissions(e)

	return e
}

func generateRoutePermissions(e *echo.Echo) {
	if flag.Lookup("test.v") != nil {
		return
	}

	routes := e.Routes()
	err := db.DB.Exec("TRUNCATE TABLE permissions RESTART IDENTITY CASCADE").Error
	if err != nil {
		panic(err)
	}

	allGrant := models.Permission{
		Method: "ALL",
		Path:   "all",
		Name:   "all",
	}
	db.DB.Create(&allGrant)

	for _, route := range routes {
		permission := models.Permission{
			Method: route.Method,
			Path:   route.Path,
			Name:   route.Name,
		}
		db.DB.Create(&permission)
	}
}
