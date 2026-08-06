package main

import (
	"log"
	"net/http"
	"time"
)

type app struct {
	config config
}

type config struct {
	addr string
}

func (app *app) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app *app) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running at %s", app.config.addr)

	return srv.ListenAndServe()
}
