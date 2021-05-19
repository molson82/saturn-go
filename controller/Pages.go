package controller

import (
	"html/template"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/molson82/saturn-go/config"
	"github.com/molson82/saturn-go/model"
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

		cards, err := model.GetAllProjectCards(c)
		logUtil.Err(err)
		handleError(w, err)

		aboutMe, err := model.GetAllAboutMeContent(c)
		logUtil.Err(err)
		handleError(w, err)

		timelineCards, err := model.GetAlltimelineCards(c)
		logUtil.Err(err)
		handleError(w, err)

		err = tpl.ExecuteTemplate(w, "index.html", struct {
			ProjectCards  interface{}
			AboutMe       interface{}
			TimeLineCards interface{}
		}{cards, aboutMe, timelineCards})

		logUtil.Err(err)
		handleError(w, err)
	}
}
