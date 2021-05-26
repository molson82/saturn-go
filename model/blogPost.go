package model

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/molson82/saturn-go/config"
)

type BlogPost struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	BlogPostAttributes `json:"attributes"`
	Relationships      `json:"relationships"`
	Meta               `json:"meta"`
}

type BlogPostAttributes struct {
	Title          string      `json:"title"`
	Status         string      `json:"status"`
	PublishedAt    time.Time   `json:"published-at"`
	ExpiresAt      interface{} `json:"expires-at"`
	BlogPostFields `json:"fields"`
}

type BlogPostFields struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Order      int    `json:"order"`
	Created    string `json:"created"`
	PostLink   string `json:"postlink"`
	CoverPhoto struct {
		Title       string `json:"title"`
		ContentType string `json:"content-type"`
		FileSize    int    `json:"file-size"`
		Url         string `json:"url"`
	} `json:"coverphoto"`
}

type BlogPostAPIResp struct {
	Data []BlogPost `json:"data"`
}

func GetAllBlogPosts(c *config.Config) ([]BlogPost, error) {
	url := c.Constants.ElegantCMSUrl + "?filter%5Btype%5D=blog-post"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []BlogPost{}, err
	}

	req.Header.Add("Authorization", "Token token="+c.Constants.ElegantCMSToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []BlogPost{}, err
	}
	defer res.Body.Close()

	var apiResp BlogPostAPIResp
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return []BlogPost{}, err
	}
	sort.Slice(apiResp.Data, func(p, q int) bool {
		return apiResp.Data[p].BlogPostAttributes.BlogPostFields.Order < apiResp.Data[q].BlogPostAttributes.BlogPostFields.Order
	})

	return apiResp.Data, nil
}
