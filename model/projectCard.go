package model

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/molson82/saturn-go/config"
)

type ProjectCard struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ProjectCardAttributes `json:"attributes"`
	Relationships         `json:"relationships"`
	Meta                  `json:"meta"`
}

type ProjectCardAttributes struct {
	Title             string      `json:"title"`
	Status            string      `json:"status"`
	PublishedAt       time.Time   `json:"published-at"`
	ExpiresAt         interface{} `json:"expires-at"`
	ProjectCardFields `json:"fields"`
}

type ProjectCardFields struct {
	Title           string `json:"title"`
	Body            string `json:"body"`
	Order           int    `json:"order"`
	BackgroundColor string `json:"backgroundcolor"`
	ButtonText      string `json:"buttontext"`
	ButtonLink      string `json:"buttonlink"`
	CardIcon        struct {
		Title       string `json:"title"`
		ContentType string `json:"content-type"`
		FileSize    int    `json:"file-size"`
		Url         string `json:"url"`
	} `json:"cardicon"`
}

type ProjectCardAPIResp struct {
	Data []ProjectCard `json:"data"`
}

func GetAllProjectCards(c *config.Config) ([]ProjectCard, error) {
	url := c.Constants.ElegantCMSUrl + "?filter%5Btype%5D=project-card"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []ProjectCard{}, err
	}

	req.Header.Add("Authorization", "Token token="+c.Constants.ElegantCMSToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []ProjectCard{}, err
	}
	defer res.Body.Close()

	var apiResp ProjectCardAPIResp
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return []ProjectCard{}, err
	}
	sort.Slice(apiResp.Data, func(p, q int) bool {
		return apiResp.Data[p].ProjectCardAttributes.ProjectCardFields.Order < apiResp.Data[q].ProjectCardAttributes.ProjectCardFields.Order
	})

	return apiResp.Data, nil
}
