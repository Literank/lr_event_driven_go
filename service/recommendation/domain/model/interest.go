package model

// Interest represents a user's interest in a book.
type Interest struct {
	UserID string  `json:"user_id" bson:"user_id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Score  float32 `json:"score"`
}
