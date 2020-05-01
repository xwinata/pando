package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pando/custom_modules/seed"
	"pando/server"
	"pando/worker"
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
			server.Start()
			fmt.Println(args[1])
			os.Exit(0)
			break
		case "worker":
			worker.Start()
			fmt.Println(args[1])
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
	[app_name] run server [queue_name]
	`

	log.Print(usagestring)
}
