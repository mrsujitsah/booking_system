package main

import (
	"net/http"

	handler "github.com/mrsujitsah/bookings/pkg/Handler"
	"github.com/mrsujitsah/bookings/pkg/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//Pat for route
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	//chi for route
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole) //own middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad) //access to session

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}
