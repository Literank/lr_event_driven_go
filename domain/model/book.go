package model

// Book represents the structure of a book
type Book struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
	Description string `json:"description"`
}
