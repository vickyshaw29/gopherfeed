package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vickyshaw29/gopherfeed/cmd/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type dbConfig struct {
	addr         string
	maxOpenCons  int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started to : %s", app.config.addr)

	return srv.ListenAndServe()
}
