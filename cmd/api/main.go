package main

import (
	"fmt"
	"log"

	"github.com/cohekoma/hello-buddy/internal/config"
	"github.com/cohekoma/hello-buddy/internal/storage"
)

func init() {
	fmt.Println("We just started!")
}

func main() {
	appCfg := appConfig{
		addr: config.GetString("ADDR", ":8080"),
	}

	appStorage := storage.NewStorage(nil)

	app := &app{
		appConfig: appCfg,
		storage:   *appStorage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
