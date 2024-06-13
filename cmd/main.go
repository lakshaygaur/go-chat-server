package main

import (
	"chatserver/config"
	"chatserver/server"
	"chatserver/storage"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	// "github.com/stvp/go-toml-config"
)

// var (
// 	country            = config.String("country", "Unknown")
// 	atlantaEnabled     = config.Bool("atlanta.enabled", false)
// 	alantaPopulation   = config.Int("atlanta.population", 0)
// 	atlantaTemperature = config.Float("atlanta.temperature", 0)
// )

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Starting server...")
	storage.InitDatabase()
	server.StartServer()
	config.SetVars()
	// config.parse("/home/lakshay/go-playground/chat-server/config/config.conf")

	// os.Exit(storage.dbHandler)
	defer storage.HandleDBclose()

}
