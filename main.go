package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/molson82/saturn-go/config"
	"github.com/molson82/saturn-go/controller"
)

func routes(c *config.Config) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	logger := httplog.NewLogger("sci-logger", httplog.Options{
		JSON:     true,
		Concise:  true,
		LogLevel: "info",
	})

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RedirectSlashes,
		httplog.RequestLogger(logger),
		middleware.Recoverer)

	r.Route("/view", func(r chi.Router) {
		r.Mount("/index", controller.GetIndexPage(c))
	})

	r.Route("/api", func(r chi.Router) {

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view/index", http.StatusSeeOther)
	})

	r.Get("/view/tailwind.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dist/tailwind.css")
	})

	return r
}

func main() {
	config := config.New()
	router := routes(config)

	port := config.Constants.Port
	log.Printf("PORT: %v\n", port)
	if port == "" {
		log.Fatal("Port must be set")
	}

	log.Printf("Golang API and Web App running...\n")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
