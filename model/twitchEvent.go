package model

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/molson82/saturn-go/config"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

type TwitchEvent struct {
	Subscription struct {
		ID        string `json:"id,omitempty"`
		Type      string `json:"type,omitempty"`
		Version   string `json:"version,omitempty"`
		Status    string `json:"status,omitempty"`
		Cost      int    `json:"cost,omitempty"`
		Condition struct {
			BroadcasterUserID string `json:"broadcaster_user_id,omitempty"`
		} `json:"condition,omitempty"`
		Transport struct {
			Method   string `json:"method,omitempty"`
			Callback string `json:"callback,omitempty"`
		} `json:"transport,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	} `json:"subscription,omitempty"`
	Event struct {
		ID                   string    `json:"id,omitempty"`
		BroadcasterUserID    string    `json:"broadcaster_user_id,omitempty"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login,omitempty"`
		BroadcasterUserName  string    `json:"broadcaster_user_name,omitempty"`
		Type                 string    `json:"type,omitempty"`
		StartedAt            time.Time `json:"started_at,omitempty"`
	} `json:"event,omitempty"`
}

func GetOAuthAccessToken(c *config.Config) (string, error) {
	oauth2Config := &clientcredentials.Config{
		ClientID:     c.Constants.TwitchClientId,
		ClientSecret: c.Constants.TwitchClientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func VerifySig(c *config.Config, r *http.Request, e string) (bool, error) {
	hmacMessage := r.Header.Get("Twitch-Eventsub-Message-Id") + r.Header.Get("Twitch-Eventsub-Message-Timestamp") + e
	signature := hmac.New(sha256.New, []byte(c.Constants.TwitchClientSecret))
	signature.Write([]byte(hmacMessage))

	sha := hex.EncodeToString(signature.Sum(nil))

	log.Printf("msgId: %v | msgTimestamp: %v", r.Header.Get("Twitch-Eventsub-Message-Id"), r.Header.Get("Twitch-Eventsub-Message-Timestamp"))
	log.Printf("hmacMessage: %v", hmacMessage)
	log.Printf("sha: %v | header: %v", sha, r.Header.Get("Twitch-Eventsub-Message-Signature"))

	if r.Header.Get("Twitch-Eventsub-Message-Signature") == fmt.Sprintf("sha256=%v", sha) {
		return true, nil
	}

	return false, errors.New("403 forbidden")
}
