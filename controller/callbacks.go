package controller

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/molson82/saturn-go/config"
	"github.com/molson82/saturn-go/model"
)

func CallbackRoutes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/twitch-online", notifyTwitchOnline(c))

	return router
}

func notifyTwitchOnline(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())

		logUtil.Info().Msg("Twitch user went live. Sending notification.")
		var tevt model.TwitchEvent
		err := render.DecodeJSON(r.Body, &tevt)
		if err != nil {
			logUtil.Error().Msg("Error reading body from callback")
			logUtil.Err(err)
			return
		}

		logUtil.Info().Msg(fmt.Sprintf("Response: %v\n", tevt))
	}
}
