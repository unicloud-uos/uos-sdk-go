package request

import (
	"bytes"
	"fmt"
)

const (
	InvalidParameterErrCode = "InvalidParameter"
	ParamRequiredErrCode    = "ParamRequiredError"
	ParamMinValueErrCode    = "ParamMinValueError"
	ParamMinLenErrCode      = "ParamMinLenError"
	ParamMaxLenErrCode      = "ParamMaxLenError"
	ParamFormatErrCode      = "ParamFormatInvalidError"
)

type Validator interface {
	Validate() error
}

type ErrInvalidParams struct {
	// op struct name
	Context string
	errs    []ErrInvalidParam
}

func (e ErrInvalidParams) Len() int {
	return len(e.errs)
}

func (e ErrInvalidParams) ErrorCode() string {
	return InvalidParameterErrCode
}

func (e ErrInvalidParams) Description() string {
	return fmt.Sprintf("%d validation error(s) found.", len(e.errs))
}

func (e ErrInvalidParams) Error() string {
	w := &bytes.Buffer{}
	fmt.Fprintf(w, "%s: %s\n", e.ErrorCode(), e.Description())

	for _, err := range e.errs {
		fmt.Fprintf(w, "- %s\n", err.Message())
	}

	return w.String()
}

func (e ErrInvalidParams) Errs() []error {
	errs := make([]error, len(e.errs))
	for i := 0; i < len(errs); i++ {
		errs[i] = e.errs[i]
	}

	return errs
}

// Add adds a new invalid parameter error to the collection of invalid parameters.
func (e *ErrInvalidParams) Add(err ErrInvalidParam) {
	err.SetContext(e.Context)
	e.errs = append(e.errs, err)
}

// AddNested adds the invalid parameter errors from another ErrInvalidParams value into this collection.
func (e *ErrInvalidParams) AddNested(nestedCtx string, nested ErrInvalidParams) {
	for _, err := range nested.errs {
		err.SetContext(e.Context)
		err.AddNestedContext(nestedCtx)
		e.errs = append(e.errs, err)
	}
}

// An ErrInvalidParam represents an invalid parameter error type.
type ErrInvalidParam interface {
	Error() string

	Code() string
	Message() string

	// Field name the error occurred on.
	Field() string

	// SetContext updates the context of the error.
	SetContext(string)

	// AddNestedContext updates the error's context to include a nested level.
	AddNestedContext(string)
}

type errInvalidParam struct {
	context       string
	nestedContext string
	field         string
	code          string
	msg           string
}

// Code returns the error code for the type of invalid parameter.
func (e *errInvalidParam) Code() string {
	return e.code
}

// Message returns the reason the parameter was invalid, and its context.
func (e *errInvalidParam) Message() string {
	return fmt.Sprintf("%s, %s.", e.msg, e.Field())
}

// Error returns the string version of the invalid parameter error.
func (e *errInvalidParam) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.Message())
}

// Field Returns the field and context the error occurred.
func (e *errInvalidParam) Field() string {
	field := e.context
	if len(field) > 0 {
		field += "."
	}
	if len(e.nestedContext) > 0 {
		field += fmt.Sprintf("%s.", e.nestedContext)
	}
	field += e.field

	return field
}

// SetContext updates the base context of the error.
func (e *errInvalidParam) SetContext(ctx string) {
	e.context = ctx
}

// AddNestedContext prepends a context to the field's path.
func (e *errInvalidParam) AddNestedContext(ctx string) {
	if len(e.nestedContext) == 0 {
		e.nestedContext = ctx
	} else {
		e.nestedContext = fmt.Sprintf("%s.%s", ctx, e.nestedContext)
	}

}

// An ErrParamRequired represents an required parameter error.
type ErrParamRequired struct {
	errInvalidParam
}

// NewErrParamRequired creates a new required parameter error.
func NewErrParamRequired(field string) *ErrParamRequired {
	return &ErrParamRequired{
		errInvalidParam{
			code:  ParamRequiredErrCode,
			field: field,
			msg:   fmt.Sprintf("missing required field"),
		},
	}
}

// An ErrParamMinLen represents a minimum length parameter error.
type ErrParamMinLen struct {
	errInvalidParam
	min int
}

// NewErrParamMinLen creates a new minimum length parameter error.
func NewErrParamMinLen(field string, min int) *ErrParamMinLen {
	return &ErrParamMinLen{
		errInvalidParam: errInvalidParam{
			code:  ParamMinLenErrCode,
			field: field,
			msg:   fmt.Sprintf("minimum field size of %v", min),
		},
		min: min,
	}
}
