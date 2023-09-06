package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "okay",
		"env":    app.config.addr,
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Print(err.Error())
	}
}
