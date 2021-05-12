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
	"github.com/newrelic/go-agent/v3/newrelic"
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
		r.Mount(newrelic.WrapHandle(c.NewRelicApp, "/index", controller.GetIndexPage(c)))
	})

	r.Route("/api", func(r chi.Router) {

	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view/index", http.StatusSeeOther)
	})

	r.Get("/view/tailwind.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/dist/tailwind.css")
	})

	r.Get("/view/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/css/styles.css")
	})

	r.Get("/view/rocket.mp4", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/video/rocket.mp4")
	})

	r.Get("/view/star_banner.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/images/star_banner.png")
	})

	r.Get("/view/togo_icon.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/images/togo_icon.png")
	})

	r.Get("/view/rocket.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/video/rocket.gif")
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
