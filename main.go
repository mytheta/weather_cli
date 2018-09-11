package main

import (
	"log"
	"os"

	"github.com/mytheta/weather_cli/controller"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = controller.Weather

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
