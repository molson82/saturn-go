package model

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
