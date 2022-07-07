package common

import "fmt"

type DBError struct {
	Err interface{}
}

func (dbError DBError) Error() string {
	return fmt.Sprintf("Database Error (%+v)", dbError.Err)
}

type ServiceError struct {
	Err interface{}
	API string
}

func (s *ServiceError) Error() string {
	return fmt.Sprintf("Service Error (%+v)", s.Err)
}

type HttpError struct {
	ErrorCode    int
	ResponseBody string
	URL          string
}

func (httpError HttpError) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ResponseBody: %s, URL: %s", httpError.ErrorCode, httpError.ResponseBody, httpError.URL)
}
