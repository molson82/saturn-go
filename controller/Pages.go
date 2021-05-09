package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/molson82/saturn-go/config"
)

func PageRoutes(c *config.Config, page string) *chi.Mux {
	router := chi.NewRouter()
	switch page {
	case "index.html":
		router.Get("/", getIndexPage(c))
	}

	return router
}

func getIndexPage(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.New("index").Funcs(template.FuncMap{}).ParseFiles("view/index.html"))
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
