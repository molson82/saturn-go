package main

import (
	"log"
	"net/http"
	"strings"

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
		r.Mount(newrelic.WrapHandle(c.NewRelicApp, "/callback", controller.CallbackRoutes(c)))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view/index", http.StatusSeeOther)
	})

	r.Mount(newrelic.WrapHandle(c.NewRelicApp, "/static/", http.StripPrefix(strings.TrimRight("/static/", "/"), http.FileServer(http.Dir("./static")))))

	return r
}

func main() {
	config := config.New()
	router := routes(config)

	//token, _ := model.GetOAuthAccessToken(config)
	//log.Printf("Token: %v", token)

	port := config.Constants.Port
	log.Printf("PORT: %v\n", port)
	if port == "" {
		log.Fatal("Port must be set")
	}

	log.Printf("Golang API and Web App running...\n")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
