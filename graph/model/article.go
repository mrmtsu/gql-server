package model

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user"`
}
