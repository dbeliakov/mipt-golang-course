package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/dbeliakov/mipt-golang-course/tasks/02/urlshortener"
)

const addr = "localhost:8080"

func main() {
	srv := urlshortener.NewShortener("http://" + addr)

	r := chi.NewMux()
	r.Put("/save", srv.HandleSave)
	r.Get("/{key}", srv.HandleExpand)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
