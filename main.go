package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	pgxConn, err := NewPGX()
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}
	defer pgxConn.Close()
	pgxConn.GetEvents()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/users", usersResource{}.Routes())
	r.Mount("/events", eventsResource{}.Routes())

	http.ListenAndServe(":3333", r)
}
