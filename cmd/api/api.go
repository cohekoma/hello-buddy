package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cohekoma/hello-buddy/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	appConfig appConfig
	storage   storage.Storage
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
		r.Post("/users", app.createUser)
	})

	return r
}

func (app *app) createUser(w http.ResponseWriter, r *http.Request) {
	app.storage.UsersStorage.Create()
	w.Write([]byte("User is created!"))
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
