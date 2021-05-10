package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/molson82/saturn-go/config"
)

func getTemplate(fm template.FuncMap, name string) *template.Template {
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob("view/template/*.tmpl.html"))
	tpl.ParseFiles("view/" + name)

	return tpl
}

func GetIndexPage(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl := getTemplate(template.FuncMap{}, "index.html")
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
