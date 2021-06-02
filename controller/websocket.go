package controller

import (
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
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logUtil.Error().Msg("Error setting up upgrader for ws")
			logUtil.Err(err)
			return
		}

		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				logUtil.Error().Msg("read error for ws")
				logUtil.Err(err)
				break
			}

			logUtil.Info().Msg(fmt.Sprintf("recv: %v\n", message))

			err = c.WriteMessage(mt, message)
			if err != nil {
				logUtil.Error().Msg("write error")
				logUtil.Err(err)
				break
			}
		}
	}
}
