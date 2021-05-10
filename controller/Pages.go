package controller

import (
	"html/template"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/molson82/saturn-go/config"
)

func getTemplate(fm template.FuncMap, name string) *template.Template {
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob("view/template/*.tmpl.html"))
	tpl.ParseFiles("view/" + name)

	return tpl
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetIndexPage(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())
		tpl := getTemplate(template.FuncMap{}, "index.html")
		err := tpl.ExecuteTemplate(w, "index.html", nil)

		logUtil.Err(err)
		handleError(w, err)
	}
}
