package main

import (
	"log"
	"net/http"

	"github.com/RichDom2185/go-simple-api-demo/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HandleGetHello)
	r.Post("/", handlers.HandlePostHello)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
