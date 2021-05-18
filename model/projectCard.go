package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/molson82/saturn-go/config"
)

type ProjectCard struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes    `json:"attributes"`
	Relationships `json:"relationships"`
	Meta          `json:"meta"`
}

type Attributes struct {
	Title       string      `json:"title"`
	Status      string      `json:"status"`
	PublishedAt time.Time   `json:"published-at"`
	ExpiresAt   interface{} `json:"expires-at"`
	Fields      struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"fields"`
}

type Relationships struct {
	App struct {
		Links struct {
			Self    string `json:"self"`
			Related string `json:"related"`
		} `json:"links"`
	} `json:"app"`
	ContentType struct {
		Links struct {
			Self    string `json:"self"`
			Related string `json:"related"`
		} `json:"links"`
	} `json:"content-type"`
}

type Meta struct {
	UpdatedAt string      `json:"updated_at"`
	CreatedAt string      `json:"created_at"`
	UpdatedBy interface{} `json:"updated_by"`
	CreatedBy string      `json:"created_by"`
	Version   int         `json:"version"`
}

type APIResp struct {
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

	var apiResp APIResp
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return []ProjectCard{}, err
	}

	return apiResp.Data, nil
}
