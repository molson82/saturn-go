package controller

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/molson82/saturn-go/config"
)

func GetRedisRoutes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/twitch-status", getCurrentTwitchStatus(c))

	return router
}

func getCurrentTwitchStatus(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())

		val, err := c.Redis.Get(context.Background(), "twitch-status").Result()
		if err != nil {
			logUtil.Info().Msg("error reading from redis server")
			logUtil.Err(err)
			render.JSON(w, r, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, struct {
			TwitchStatus string `json:"twitchStatus"`
		}{val})
	}
}
