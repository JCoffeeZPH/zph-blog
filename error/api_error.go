package error

import "net/http"

type ApiError struct {
	HttpStatusCode int    `json:"-"`
	ErrorCode      string `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`
	params         []interface{}
	Data           interface{} `json:"data"`
}

func NewApiStatusError(httpStatusCode int, errCode string, errorMsg string) *ApiError {
	return &ApiError{
		HttpStatusCode: httpStatusCode,
		ErrorCode:      errCode,
		ErrorMsg:       errorMsg,
	}
}

func NewApiError(errorCode string, errorMsg string) *ApiError {
	return &ApiError{
		HttpStatusCode: 491,
		ErrorCode:      errorCode,
		ErrorMsg:       errorMsg,
	}
}

var (
	ParamError    = NewApiStatusError(http.StatusBadRequest, "param_error", "Param error.")
	Unauthorized  = NewApiStatusError(http.StatusUnauthorized, "user_is_unauthorized", "User is unauthorized.")
	AlreadyExists = NewApiError("already_exists", "")
	NotFound      = NewApiStatusError(http.StatusNotFound, "Resource_not_found", "Resource not found")
	CacheError    = NewApiStatusError(http.StatusInternalServerError, "cache_error", "cache error")
)
