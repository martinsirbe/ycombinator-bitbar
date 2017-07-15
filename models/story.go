package models

// Story used to map hacker-news.firebaseio.com story response.
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"-"` // Ignore kids >:)
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}