package router

import (
	"net/http"

	"github.com/cyan-store/hook/handlers"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/checkout", handlers.Checkout)
	mux.NotFound(handlers.NotFound)
	mux.MethodNotAllowed(handlers.MethodNotAllowed)

	return mux
}
