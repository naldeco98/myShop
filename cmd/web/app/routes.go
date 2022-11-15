package app

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	// basic url middlewares
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	// .-
	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Get("/payment-succeeded", app.PaymentSucceeded)
	return mux
}
