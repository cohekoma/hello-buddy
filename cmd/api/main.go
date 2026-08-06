package main

import "log"

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
