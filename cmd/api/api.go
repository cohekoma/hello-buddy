package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	appConfig appConfig
}

type appConfig struct {
	addr string
}

func (app *app) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *app) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.appConfig.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running at %s", app.appConfig.addr)

	return srv.ListenAndServe()
}
