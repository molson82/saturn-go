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
	router.Post("/twitch-offline", notifyTwitchOffline(c))

	return router
}

func notifyTwitchOnline(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())
		logUtil.Info().Msg("Twitch user went live. Sending notification.")

		var vobj model.TwitchEvent
		err := render.DecodeJSON(r.Body, &vobj)
		if err != nil {
			logUtil.Info().Msg("Error reading in request body")
			render.JSON(w, r, http.StatusInternalServerError)
			return
		}

		//if _, err := model.VerifySig(c, r, vobj); err != nil {
		//logUtil.Info().Msg("Verify Sig failed. Return 403")
		//render.JSON(w, r, http.StatusForbidden)
		//return
		//}

		logUtil.Info().Msg(fmt.Sprintf("Request Body: %v", vobj))

		logUtil.Info().Msg("Verify Sig success. Respond to callback")
		w.Write([]byte(vobj.Challenge))
	}
}

func notifyTwitchOffline(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())

		logUtil.Info().Msg("Twitch user went offline. Sending notification.")
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
