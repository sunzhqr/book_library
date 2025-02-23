package model

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
