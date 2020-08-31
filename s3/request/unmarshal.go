package request

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	. "github.com/uos-sdk-go/s3/error"
	"github.com/uos-sdk-go/s3/helper"
)

var UnmarshalRequestHandler = HandlerItem{
	Name: "sdk.request.unmarshal.request",
	Fn: func(request *Request) {
		if u, ok := request.Data.(UnmarshalerForOut); ok {
			request.RequestID = request.HTTPResponse.Header.Get(helper.UOSRequestID)
			err := u.UnmarshalHeader(request)
			if err != nil {
				request.Error = NewBaseError("UnmarshalHeaderErr", "Unmarshal response header failed!", nil)
				return
			}
			if request.HTTPResponse.StatusCode >= 300 {
				request.Error = &ResponseError{Response: request.HTTPResponse}
				return
			}
			// Unmarshal errorResponse
			err = UnmarshalError(request)
			if err != nil {
				request.Error = NewBaseError("UnmarshalErrorErr", "failed to decode query XML error response", err)
				return
			}

			decoder := xml.NewDecoder(io.LimitReader(request.HTTPResponse.Body, request.HTTPResponse.ContentLength))
			err = u.UnmarshalBody(decoder)
			if err != nil {
				request.Error = NewBaseError("UnmarshalBodyErr", "Unmarshal response body failed!", err)
				return
			}
			return
		} else {
			request.Error = NewBaseError("UnmarshalRequestErr", "Request Data type unasserted", nil)
			return

		}
	},
}

type UnmarshalerForOut interface {
	UnmarshalBody(*xml.Decoder) error
	UnmarshalHeader(*Request) error
}

type responseError struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Code    string   `xml:"Error>Code"`
	Message string   `xml:"Error>Message"`
}

func UnmarshalError(r *Request) error {
	defer r.HTTPResponse.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Config.Logger.Debug("Failed to read from query HTTP response body", err)
		return err
	}
	resp := responseError{}
	decodeErr := xml.Unmarshal(bodyBytes, &resp)
	if decodeErr == nil {
		r.Error = NewApiError(
			NewBaseError(resp.Code, resp.Message, nil),
			r.HTTPResponse.StatusCode,
		)
		return nil
	}

	return decodeErr
}
