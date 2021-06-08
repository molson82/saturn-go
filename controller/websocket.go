package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/gorilla/websocket"
	"github.com/molson82/saturn-go/config"
)

func WebSocketRoutes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/twitch-status", getTwitchStatusWS(c))

	return router
}

func getTwitchStatusWS(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logUtil := httplog.LogEntry(r.Context())
		logUtil.Info().Msg("Twitch Status Websocket connection starting...")

		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logUtil.Error().Msg("Error setting up upgrader for ws")
			logUtil.Err(err)
			return
		}

		defer conn.Close()
		for {
			mt, _, err := conn.ReadMessage()
			if err != nil {
				logUtil.Error().Msg("read error for ws")
				logUtil.Err(err)
				break
			}

			val, err := c.Redis.Get(context.Background(), "twitch-status").Result()
			if err != nil {
				logUtil.Info().Msg("error reading from redis server")
				logUtil.Err(err)
				break
			}

			logUtil.Info().Msg(fmt.Sprintf("redis res: %v\n", val))

			status, _ := json.Marshal(struct {
				TwitchStatus string `json:"twitchStatus"`
			}{val})

			err = conn.WriteMessage(mt, status)
			if err != nil {
				logUtil.Error().Msg("write error")
				logUtil.Err(err)
				break
			}
		}
	}
}
