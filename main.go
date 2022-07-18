package main

import (
	"fmt"
	"log"
	"maurodesouza/first-golang-api/configs"
	"maurodesouza/first-golang-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Get("/", handlers.List)
	router.Post("/", handlers.Create)
	router.Get("/{id}", handlers.Get)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)

	port := configs.GetServerPort()

	log.Printf(`server running in localhost:%s`, port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
