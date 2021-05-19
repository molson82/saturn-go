package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/molson82/saturn-go/config"
)

type AboutMe struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	AboutMeAttributes `json:"attributes"`
	Relationships     `json:"relationships"`
	Meta              `json:"meta"`
}

type AboutMeAttributes struct {
	Title         string      `json:"title"`
	Status        string      `json:"status"`
	PublishedAt   time.Time   `json:"published-at"`
	ExpiresAt     interface{} `json:"expires-at"`
	AboutMeFields `json:"fields"`
}

type AboutMeFields struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type AboutMeAPIResp struct {
	Data []AboutMe `json:"data"`
}

func GetAllAboutMeContent(c *config.Config) ([]AboutMe, error) {
	url := c.Constants.ElegantCMSUrl + "?filter%5Btype%5D=about-me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []AboutMe{}, err
	}

	req.Header.Add("Authorization", "Token token="+c.Constants.ElegantCMSToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []AboutMe{}, err
	}
	defer res.Body.Close()

	var apiResp AboutMeAPIResp
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return []AboutMe{}, err
	}

	return apiResp.Data, nil
}
