package rest_err

import (
	"net/http"

	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type RestErr struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"err"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func ConvertToRestErr(internalError *internal_error.InternalError) *RestErr {
	switch internalError.Err {
	case "bad_request":
		return NewBadRequestError(internalError.Error())
	case "not_found":
		return NewNotFoundError(internalError.Error())
	default:
		return NewInternalServerError(internalError.Error())
	}
}

func NewBadRequestValidationError(message string, causes ...Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
		Causes:  causes,
	}
}

func NewBadRequestError(message string, causes ...Cause) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "internal_server_error",
		Causes:  nil,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "not_found",
		Causes:  nil,
	}
}
