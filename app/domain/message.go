package domain

type Message struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}
