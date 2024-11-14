package apierror

import (
	"net/http"
)

type HTTPError interface {
	Error() string
	Errors() []string
	HTTPStatus() int
}

type APIError struct {
	msg        string
	statusCode int
	// errors     []APIError
	errors []string
}

func New(err error, status int) APIError {
	return APIError{msg: err.Error(), statusCode: status}
}

func NewGeneric(err error) APIError {
	return APIError{msg: err.Error(), statusCode: http.StatusInternalServerError}
}

func NewMulti(err APIError, errors []string) APIError {
	return APIError{msg: err.msg, statusCode: err.statusCode, errors: errors}
}

func NewGenericMulti(msg string, statusCode int, errors []string) APIError {
	return APIError{msg, statusCode, errors}
}

func (e APIError) Error() string {
	return e.msg
}

func (e APIError) HTTPStatus() int {
	return e.statusCode
}

func (e APIError) Errors() []string {
	return e.errors
}

var ErrForbidden = APIError{
	msg:        "permission denied",
	statusCode: http.StatusForbidden,
}

// var ErrServerError = APIError{
// 	msg:        "internal server error",
// 	StatusCode: http.StatusInternalServerError,
// }

var ErrNotFound = APIError{
	msg:        "path not found",
	statusCode: http.StatusNotFound,
}

var ErrDataNotFound = APIError{
	msg:        "not found",
	statusCode: http.StatusNotFound,
}

var ErrContextMissingLogger = APIError{
	msg:        "logger missing in context",
	statusCode: http.StatusInternalServerError,
}

var ErrContextMissingRequestID = APIError{
	msg:        "request_id missing in context",
	statusCode: http.StatusInternalServerError,
}

var ErrContextMissingEnvelope = APIError{
	msg:        "envelope missing in context",
	statusCode: http.StatusInternalServerError,
}

var ErrContextMissing = APIError{
	msg:        "missing in context",
	statusCode: http.StatusInternalServerError,
}

var ErrRouteNotFound = APIError{
	msg:        "path not found",
	statusCode: http.StatusNotFound,
}

var ErrServerError = APIError{
	msg:        "internal server error",
	statusCode: http.StatusInternalServerError,
}
