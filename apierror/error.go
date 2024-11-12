package apierror

import (
	"net/http"
)

type HTTPError interface {
	Error() string
	HTTPStatus() int
}

type APIError struct {
	msg        string
	StatusCode int
}

func New(err error, status int) APIError {
	return APIError{msg: err.Error(), StatusCode: status}
}

func NewGeneric(err error) APIError {
	return APIError{msg: err.Error(), StatusCode: http.StatusInternalServerError}
}

func (e APIError) Error() string {
	return e.msg
}

func (e APIError) HTTPStatus() int {
	return e.StatusCode
}

var ErrForbidden = APIError{
	msg:        "permission denied",
	StatusCode: http.StatusForbidden,
}

var ErrServerError = APIError{
	msg:        "internal server error",
	StatusCode: http.StatusInternalServerError,
}

var ErrNotFound = APIError{
	msg:        "path not found",
	StatusCode: http.StatusNotFound,
}

var ErrDataNotFound = APIError{
	msg:        "not found",
	StatusCode: http.StatusNotFound,
}

var ErrContextMissingLogger = APIError{
	msg:        "logger missing in context",
	StatusCode: http.StatusInternalServerError,
}

var ErrContextMissingRequestID = APIError{
	msg:        "request_id missing in context",
	StatusCode: http.StatusInternalServerError,
}

var ErrContextMissingEnvelope = APIError{
	msg:        "envelope missing in context",
	StatusCode: http.StatusInternalServerError,
}

// var ErrContextMissingLogger = APIError{
// 	msg:        "logger missing in context",
// 	StatusCode: http.StatusInternalServerError,
// }
// var ErrContextMissingLogger = APIError{
// 	msg:        "logger missing in context",
// 	StatusCode: http.StatusInternalServerError,
// }
