package model

import "time"

// Trend represents the structure of a
// trendy query and its related books.
type Trend struct {
	Query     string    `json:"query"`
	Books     []Book    `json:"books"`
	CreatedAt time.Time `json:"created_at"`
}
