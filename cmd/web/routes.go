package main

import (
	"net/http"

	"github.com/aleph-design/hello/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// func routes(app *config.AppConfig) http.Handler {
func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)				// CSRF protection POST request
	mux.Use(SessionLoad)	// use the session

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// any thing not a template goes in ./static
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
