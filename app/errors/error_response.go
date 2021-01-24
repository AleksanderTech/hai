package errors

type ErrorResponse struct {
	ErrorCodes   []string `json:"errorCodes"`
	ErrorDetails []string `json:"errorDetails"`
}

func NewErrResponse(errorCodes []string, errorDetails []string) ErrorResponse {
	return ErrorResponse{ErrorCodes: errorCodes, ErrorDetails: errorDetails}
}
