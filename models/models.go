package models

type Validation struct {
	ContentType   string `json:"contentType"`
	Authorization string `json:"authorization"`
}

type Client struct {
	Username string `json:"symbol"`
	Faction  string `json:"faction"`
}

type ResponseFactions struct {
	Data []struct {
		Description  string `json:"description"`
		Headquarters string `json:"headquarters"`
		IsRecruiting bool   `json:"isRecruiting"`
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Traits       []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Symbol      string `json:"symbol"`
		} `json:"traits"`
	} `json:"data"`
	Meta struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
		Total int `json:"total"`
	} `json:"meta"`
}
