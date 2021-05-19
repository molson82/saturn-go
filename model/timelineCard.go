package model

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/molson82/saturn-go/config"
)

type TimelineCard struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	TimelineCardAttributes `json:"attributes"`
	Relationships          `json:"relationships"`
	Meta                   `json:"meta"`
}

type TimelineCardAttributes struct {
	Title              string      `json:"title"`
	Status             string      `json:"status"`
	PublishedAt        time.Time   `json:"published-at"`
	ExpiresAt          interface{} `json:"expires-at"`
	TimelineCardFields `json:"fields"`
}

type TimelineCardFields struct {
	Title           string `json:"title"`
	Body            string `json:"body"`
	BackgroundColor string `json:"backgroundcolor"`
	Date            string `json:"date"`
	Order           int    `json:"order"`
	Width           string `json:"width"`
	CardIcon        struct {
		Title       string `json:"title"`
		ContentType string `json:"content-type"`
		FileSize    int    `json:"file-size"`
		Url         string `json:"url"`
	} `json:"cardicon"`
}

type TimelineCardAPIResp struct {
	Data []TimelineCard `json:"data"`
}

func GetAlltimelineCards(c *config.Config) ([]TimelineCard, error) {
	url := c.Constants.ElegantCMSUrl + "?filter%5Btype%5D=timeline-card"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []TimelineCard{}, err
	}

	req.Header.Add("Authorization", "Token token="+c.Constants.ElegantCMSToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []TimelineCard{}, err
	}
	defer res.Body.Close()

	var apiResp TimelineCardAPIResp
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return []TimelineCard{}, err
	}
	sort.Slice(apiResp.Data, func(p, q int) bool {
		return apiResp.Data[p].TimelineCardAttributes.TimelineCardFields.Order < apiResp.Data[q].TimelineCardAttributes.TimelineCardFields.Order
	})

	return apiResp.Data, nil
}
