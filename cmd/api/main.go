package main

import (
	"log"

	env "github.com/vickyshaw29/gopherfeed/cmd/internal"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
