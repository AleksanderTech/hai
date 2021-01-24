package errors

const (
	MessageNotFound             string = "MessageNotfound"
	DeleteNonexistentMessage    string = "DeleteNonexistentMessage"
	DatabaseActionError         string = "DatabaseActionError"
	InvalidCreateMessageRequest string = "InvalidCreateMessageRequest"
	InvalidBodyRequest          string = "InvalidBodyRequest"
	InvalidQueryParam           string = "InvalidQueryParam"
	InternalServerError         string = "InternalServerError"
)
