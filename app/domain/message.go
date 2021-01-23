package domain

type Message struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}
