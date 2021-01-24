package errors

import "fmt"

const (
	InvalidMessageContent           string = "Message 'content' cannot be empty."
	InvalidMessageTitle             string = "Message 'title' cannot be empty."
	InvalidEmailFormat              string = "Invalid email format."
	EmptyEmail                      string = "Field 'email' cannot be empty."
	MessageNotFoundDetails          string = "The searching message does not exists."
	InternalServerErrorDetails      string = "Internal server error. Try again later."
	DeleteNonexistentMessageDetails string = "Target message does not exist."
	InvalidJson                     string = "Invalid passed json format."
)

func EmptyFieldOf(field string, entity string) string {
	return fmt.Sprintf("Field '%v' of %v cannot be empty.", field, entity)
}

func RequiredFieldOf(field string, entity string) string {
	return fmt.Sprintf("Field '%v' of %v is required.", field, entity)
}
