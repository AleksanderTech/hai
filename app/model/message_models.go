package model

type DeleteMessageRequest struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type CreateMessageResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

type CreateMessageRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}
