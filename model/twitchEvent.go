package model

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/molson82/saturn-go/config"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

type TwitchEvent struct {
	Subscription struct {
		ID        string `json:"id"`
		Type      string `json:"type"`
		Version   string `json:"version"`
		Status    string `json:"status"`
		Cost      int    `json:"cost"`
		Condition struct {
			BroadcasterUserID string `json:"broadcaster_user_id"`
		} `json:"condition"`
		Transport struct {
			Method   string `json:"method"`
			Callback string `json:"callback"`
		} `json:"transport"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"subscription"`
	Event struct {
		ID                   string    `json:"id"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		Type                 string    `json:"type"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
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

func VerifySig(c *config.Config, r *http.Request) (bool, error) {
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	hmacMessage := r.Header.Get("Twitch-Eventsub-Message-Id") + r.Header.Get("Twitch-Eventsub-Message-Timestamp") + string(rBody)
	signature := hmac.New(sha256.New, []byte(c.Constants.TwitchClientSecret))
	signature.Write([]byte(hmacMessage))

	sha := hex.EncodeToString(signature.Sum(nil))

	log.Printf("hmacMessage: %v", hmacMessage)
	log.Printf("sha: %v | header: %v", sha, r.Header.Get("Twitch-Eventsub-Message-Signature"))

	if r.Header.Get("Twitch-Eventsub-Message-Signature") == fmt.Sprintf("sha256=%v", sha) {
		return true, nil
	}

	return false, errors.New("403 forbidden")
}
