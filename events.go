package main

import (
	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type eventsResource struct{}

func (rs eventsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /events - Read a list of events.
	r.Post("/", rs.Create) // POST /events - Create a new event.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(PostCtx)
		r.Get("/", rs.Get) // GET /events/{id} - Read a single event by :id.
		// r.Put("/", rs.Update)    // PUT /events/{id} - Update a single event by :id.
		// r.Delete("/", rs.Delete) // DELETE /events/{id} - Delete a single event by :id.
	})

	return r
}

// Request Handler - GET /events - Read a list of events.
func (rs eventsResource) List(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	// GetEvents(conn)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Request Handler - POST /events - Create a new event.
func (rs eventsResource) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Request Handler - GET /events/{id} - Read a single event by :id.
func (rs eventsResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get events by id.."))
}
