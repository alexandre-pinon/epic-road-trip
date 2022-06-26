package model

type Hotel struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results          []ActivityResult `json:"results"`
	Status string `json:"status"`
}