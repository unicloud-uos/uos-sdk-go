package error

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	ErrorCode() string
	Description() string
}

func SprintError(code, description, extra string, error error) string {
	msg := fmt.Sprintf("%s: %s", code, description)
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}
	if error != nil {
		msg = fmt.Sprintf("%s\ncaused by: %v", msg, error)
	}
	return msg
}

type BaseError struct {
	errCode     string
	description string
	error       error
}

func NewBaseError(code, message string, err error) *BaseError {
	b := &BaseError{
		errCode:     code,
		description: message,
		error:       err,
	}
	return b
}

func (b BaseError) Error() string {
	return SprintError(b.errCode, b.description, "", b.error)
}

func (b BaseError) String() string {
	return b.Error()
}

func (b BaseError) ErrorCode() string {
	return b.errCode
}

func (b BaseError) Description() string {
	return b.description
}

type SDKError struct {
	ErrorCode   string
	Description string
}

type ErrCode int

const (
	ErrMissingEndpoint ErrCode = iota
	ErrMissingRegion
)

var ErrorCodeStruct = map[ErrCode]SDKError{
	ErrMissingEndpoint: {
		ErrorCode:   "MissingEndpoint",
		Description: "'Endpoint' configuration is required for this service",
	},
	ErrMissingRegion: {
		ErrorCode:   "MissingRegion",
		Description: "could not find region configuration",
	},
}

func (e ErrCode) Error() string {
	err, ok := ErrorCodeStruct[e]
	if !ok {
		return "InternalError"
	}
	return fmt.Sprintf(err.ErrorCode, err.Description)
}

type ResponseError struct {
	Response *http.Response
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("HTTP response error, %s, %d",
		e.Response.Status, e.Response.StatusCode)
}

type apiError struct {
	Err        Error
	StatusCode int
	RequestID  string
}

func NewApiError(err Error, statusCode int, id string) *apiError {
	return &apiError{
		Err:        err,
		StatusCode: statusCode,
		RequestID:  id,
	}
}

func (a apiError) Error() string {
	extra := fmt.Sprintf("status code: %d, requestID: %s", a.StatusCode, a.RequestID)
	return SprintError(a.Err.ErrorCode(), a.Err.Description(), extra, nil)
}
