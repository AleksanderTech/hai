package errors

import "fmt"

type HaiError struct {
	ErrorCode string
}

func New(errorCode string) HaiError {
	return HaiError{ErrorCode: errorCode}
}

func (e HaiError) Error() string {
	return fmt.Sprintf("Error occured, message: %v", e.ErrorCode)
}
