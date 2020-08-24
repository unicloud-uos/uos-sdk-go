package error

import "fmt"

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

type baseError struct {
	errCode     string
	description string
	error       error
}

func NewBaseError(code, message string, err error) *baseError {
	b := &baseError{
		errCode:     code,
		description: message,
		error:       err,
	}
	return b
}

func (b baseError) Error() string {
	return SprintError(b.errCode, b.description, "", b.error)
}

func (b baseError) String() string {
	return b.Error()
}

func (b baseError) ErrorCode() string {
	return b.errCode
}

func (b baseError) Description() string {
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
