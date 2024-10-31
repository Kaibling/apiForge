package error

import "net/http"

type HTTPError interface {
	Error() string
	HTTPStatus() int
}

type ApiError struct {
	msg        string
	StatusCode int
}

// func new(err error, status int) ApiError {
// 	return ApiError{msg: err.Error(), StatusCode: status}
// }

// func NewGeneric(err error) ApiError {
// 	return ApiError{msg: err.Error(), StatusCode: http.StatusInternalServerError}
// }

func (e ApiError) Error() string {
	return e.msg
}

func (e ApiError) HTTPStatus() int {
	return e.StatusCode
}

var Forbidden = ApiError{
	msg:        "permission denied",
	StatusCode: http.StatusForbidden,
}

var ServerError = ApiError{
	msg:        "internal server error",
	StatusCode: http.StatusInternalServerError,
}

var NotFound = ApiError{
	msg:        "path not found",
	StatusCode: http.StatusNotFound,
}

var DataNotFound = ApiError{
	msg:        "not found",
	StatusCode: http.StatusNotFound,
}
