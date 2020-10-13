package request

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	. "github.com/unicloud-uos/uos-sdk-go/s3/error"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
)

var MarshalRequestHandler = HandlerItem{
	Name: "sdk.request.marshal.parameters",
	Fn: func(request *Request) {
		if m, ok := request.Params.(MarshalerForPut); ok {
			encoder := NewEncoder(request.HTTPRequest)
			err := m.MarshalForPut(encoder)
			if err != nil {
				request.Error = NewBaseError("MarshalFieldErr", "Marshal field for put error!", err)
				return
			}
			var body io.ReadSeeker
			request.HTTPRequest, body, err = encoder.Encode()
			if err != nil {
				request.Error = NewBaseError("SerializationError", "failed to encode rest XML request", err)
				return
			}
			if body != nil {
				request.SetReaderBody(body)
			}
		} else {
			request.Error = NewBaseError("MarshalRequestErr", "Request Params type unasserted", nil)
			return
		}
	},
}

type MarshalerForPut interface {
	MarshalForPut(*EncoderForPut) error
}

type EncoderForPut struct {
	method  string
	request *http.Request
	path    PathReplace
	query   url.Values
	header  http.Header
	payload io.ReadSeeker
	err     error
}

func NewEncoder(req *http.Request) *EncoderForPut {
	encoder := &EncoderForPut{
		method:  req.Method,
		request: req,
		path:    NewPathReplace(req.URL.Path),
		query:   req.URL.Query(),
		header:  req.Header,
	}
	return encoder
}

func (e *EncoderForPut) SetValue(t, k string, v ValueMarshaler) {
	if e.err != nil {
		return
	}

	var str string
	str, e.err = v.MarshalValue()
	if e.err != nil {
		return
	}

	switch t {
	case helper.HeaderTarget:
		k = strings.TrimSpace(k)
		str = strings.TrimSpace(str)
		e.header.Set(k, str)
	case helper.PathTarget:
		e.path.ReplaceElement(k, str)
	case helper.QueryTarget:
		e.query.Set(k, str)
	case helper.BodyTarget:
		if e.method != "GET" {
			e.err = fmt.Errorf("body target not supported for rest non-GET methods %s, %s", t, k)
			return
		}
		e.query.Set(k, str)
	default:
		e.err = fmt.Errorf("unknown SetValue rest encode target, %s, %s", t, k)
	}

}

func (e *EncoderForPut) SetStream(t, k string, v StreamMarshaler) {
	if e.err != nil {
		return
	}

	switch t {
	case helper.PayloadTarget:
		e.payload, e.err = v.MarshalStream()
	default:
		e.err = fmt.Errorf("unknown SetStream rest encode target, %s, %s", t, k)
	}

}

func (e *EncoderForPut) SetStruct(t, k string, v interface{}) {
	if e.err != nil {
		return
	}

	switch t {
	case helper.PayloadTarget:
		fallthrough
	case helper.BodyTarget:
		data, err := xml.Marshal(v)
		if err != nil {
			e.err = err
			return
		}

		fmt.Println(string(data))

		e.payload = bytes.NewReader(data)
	default:
		e.err = fmt.Errorf("unknown SetStruct rest encode target, %s, %s", t, k)
	}
}

func (e *EncoderForPut) Encode() (*http.Request, io.ReadSeeker, error) {
	if e.err != nil {
		return nil, nil, e.err
	}

	e.request.URL.Path, e.request.URL.RawPath = e.path.Encode()
	e.request.URL.RawQuery = e.query.Encode()
	e.request.Header = e.header

	return e.request, e.payload, nil
}

type ValueMarshaler interface {
	MarshalValue() (string, error)
}

type StreamMarshaler interface {
	MarshalStream() (io.ReadSeeker, error)
}
