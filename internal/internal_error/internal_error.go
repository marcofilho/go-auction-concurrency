package internal_error

type InternalError struct {
	Message string `json:"message"`
	Err     string `json:"err"`
}

func (i *InternalError) Error() string {
	return i.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server_error",
	}
}
