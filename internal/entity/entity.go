package entity

type CatResponse struct {
	CurrentPage int `json:"current_page"`
	Data        []struct {
		Breed   string `json:"breed"`
		Country string `json:"country"`
		Origin  string `json:"origin"`
		Coat    string `json:"coat"`
		Pattern string `json:"pattern"`
	} `json:"data"`
	FirstPageUrl string `json:"first_page_url"`
	From         int    `json:"from"`
	LastPage     int    `json:"last_page"`
	LastPageUrl  string `json:"last_page_url"`
	Links        []struct {
		Url    *string `json:"url"`
		Label  string  `json:"label"`
		Active bool    `json:"active"`
	} `json:"links"`
	NextPageUrl string      `json:"next_page_url"`
	Path        string      `json:"path"`
	PerPage     int         `json:"per_page"`
	PrevPageUrl interface{} `json:"prev_page_url"`
	To          int         `json:"to"`
	Total       int         `json:"total"`
}
