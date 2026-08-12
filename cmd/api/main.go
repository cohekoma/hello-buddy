package main

import (
	"fmt"
	"log"

	"github.com/cohekoma/hello-buddy/internal/config"
)

func init() {
	fmt.Println("We just started!")
}

func main() {
	appCfg := appConfig{
		addr: config.GetString("ADDR", ":8080"),
	}
	app := &app{
		appConfig: appCfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
