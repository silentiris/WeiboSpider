package model

type Result struct {
	Content string   `json:"content"`
	Urls    []string `json:"picUrls"`
}
