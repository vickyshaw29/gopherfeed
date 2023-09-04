package main

import (
	"log"

	env "github.com/vickyshaw29/gopherfeed/cmd/internal"
	"github.com/vickyshaw29/gopherfeed/cmd/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
