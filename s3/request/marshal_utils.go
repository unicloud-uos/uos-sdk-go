package request

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"io"
	"strconv"
	"strings"
	"time"
)

type BoolValue bool

// MarshalValue formats the value into a string for encoding.
func (v BoolValue) MarshalValue() (string, error) {
	return strconv.FormatBool(bool(v)), nil
}

type StringValue string

// MarshalValue formats the value into a string for encoding.
func (v StringValue) MarshalValue() (string, error) {
	return string(v), nil
}

type Int64Value int64

// MarshalValue formats the value into a string for encoding.
func (v Int64Value) MarshalValue() (string, error) {
	return strconv.FormatInt(int64(v), 10), nil
}

type TimeValue struct {
	T      time.Time
	Format string
}

func (v TimeValue) MarshalValue() (string, error) {
	return helper.FormatTime(v.Format, v.T)
}

type ReadSeekerStream struct {
	V io.ReadSeeker
}

// MarshalStream returns the wrapped io.ReadSeeker for encoding.
func (v ReadSeekerStream) MarshalStream() (io.ReadSeeker, error) {
	return v.V, nil
}

// A StringStream aliases a string to be used as a StreamMarshaler.
type StringStream string

// MarshalStream marshals a string into an io.ReadSeeker for encoding.
func (v StringStream) MarshalStream() (io.ReadSeeker, error) {
	return strings.NewReader(string(v)), nil
}
