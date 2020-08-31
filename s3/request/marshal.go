package request

import (
	"io"
	"net/http"
	"net/url"

	. "github.com/uos-sdk-go/s3/error"
)

var MarshalRequestHandler = HandlerItem{
	Name: "sdk.request.marshal.request",
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

func (e *EncoderForPut) Encode() (*http.Request, io.ReadSeeker, error) {
	if e.err != nil {
		return nil, nil, e.err
	}

	e.request.URL.Path, e.request.URL.RawPath = e.path.Encode()
	e.request.URL.RawQuery = e.query.Encode()
	e.request.Header = e.header

	return e.request, e.payload, nil
}
