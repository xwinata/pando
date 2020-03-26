package main

import (
	"flag"
	"log"
	"os"
	"pando/custom_modules/seed"
	"pando/server/router"

	"github.com/labstack/echo/middleware"
)

var (
	flags = flag.NewFlagSet("pando", flag.ExitOnError)
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()

	switch args[0] {
	default:
		flags.Usage()
		break
	case "run":
		switch args[1] {
		default:
			flags.Usage()
			break
		case "server":
			e := router.NewRouter()
			if os.Getenv("PANDO_CORS") == "true" {
				e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
					AllowOrigins: []string{"*"},
					AllowMethods: []string{"*"},
					AllowHeaders: []string{"*"},
				}))
			}
			e.Logger.Fatal(e.Start(":" + os.Getenv("PANDO_PORT")))
			os.Exit(0)
			break
		case "worker":
			os.Exit(0)
			break
		}
		break
	case "seed":
		seed.Seed()
		os.Exit(0)
		break
	case "unseed":
		seed.Unseed()
		os.Exit(0)
		break
	}
}

func usage() {
	usagestring := `
to run the app :
	[app_name] run
	`

	log.Print(usagestring)
}
