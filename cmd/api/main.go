package main

import (
	"fmt"
	"log"
)

func init() {
	fmt.Println("We just started!")
}

func main() {
	cfg := config{
		addr: ":3004",
	}
	app := &app{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
