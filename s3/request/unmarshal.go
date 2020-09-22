package request

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	. "github.com/uos-sdk-go/s3/error"
	"github.com/uos-sdk-go/s3/helper"
)

var UnmarshalRequestHandler = HandlerItem{
	Name: "sdk.request.unmarshal.request",
	Fn: func(request *Request) {
		if u, ok := request.Data.(UnmarshalerForOut); ok {
			request.RequestID = request.HTTPResponse.Header.Get(helper.UOSRequestID)
			err := UnmarshalHeader(request)
			if err != nil {
				request.Error = NewBaseError("UnmarshalHeaderErr", "Unmarshal response header failed!", nil)
				return
			}
			if request.HTTPResponse.StatusCode >= 300 {
				request.Error = &ResponseError{Response: request.HTTPResponse}
			}

			if request.Error != nil {
				// Unmarshal errorResponse
				UnmarshalError(request)
				return
			}

			err = u.UnmarshalBody(request)
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
	UnmarshalBody(r *Request) error
}

func UnmarshalHeader(r *Request) error {
	r.RequestID = r.HTTPResponse.Header.Get(helper.UOSRequestID)
	v := reflect.Indirect(reflect.ValueOf(r.Data))
	for i := 0; i < v.NumField(); i++ {
		m, field := v.Field(i), v.Type().Field(i)
		if n := field.Name; n[0:1] == strings.ToLower(n[0:1]) {
			continue
		}

		if m.IsValid() {
			name := field.Tag.Get("locationName")
			if name == "" {
				name = field.Name
			}

			switch field.Tag.Get("location") {
			case "statusCode":
				unmarshalStatusCode(m, r.HTTPResponse.StatusCode)
			case "header":
				err := unmarshalHeader(m, r.HTTPResponse.Header, name, field.Tag)
				if err != nil {
					r.Error = NewBaseError("SerializationError", "failed to decode REST response", err)
					break
				}
			}
		}
		if r.Error != nil {
			return r.Error
		}
	}
	return nil
}

type responseError struct {
	XMLName   xml.Name `xml:"Error"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
	RequestId string   `xml:"RequestId"`
}

func UnmarshalError(r *Request) {
	defer r.HTTPResponse.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Config.Logger.Debug("Failed to read from query HTTP response body", err)
		r.Error = NewBaseError("UnmarshalErrorErr", "failed to read from query HTTP response body", err)
	}
	resp := responseError{}
	decodeErr := xml.Unmarshal(bodyBytes, &resp)
	if decodeErr == nil {
		r.Error = NewApiError(
			NewBaseError(resp.Code, resp.Message, nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
	} else {
		r.Config.Logger.Debug("Failed to decode query XML error response", err)
		r.Error = NewApiError(
			NewBaseError("UnmarshalError",
				r.HTTPResponse.Status, decodeErr),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
	}
}

func unmarshalStatusCode(v reflect.Value, statusCode int) {
	if !v.IsValid() {
		return
	}

	switch v.Interface().(type) {
	case *int64:
		s := int64(statusCode)
		v.Set(reflect.ValueOf(&s))
	case int64:
		s := int64(statusCode)
		v.Set(reflect.ValueOf(s))
	}
}

func unmarshalHeader(v reflect.Value, headers http.Header, name string, tag reflect.StructTag) error {
	if v.Kind() == reflect.String {
		if len(headers.Get(name)) > 0 {
			v.SetString(headers.Get(name))
		}
		return nil
	} else if !v.IsValid() || headers.Get(name) == "" {
		return nil
	}

	switch v.Interface().(type) {
	case *string:
		header := headers.Get(name)
		v.Set(reflect.ValueOf(&header))
	case string:
		v.Set(reflect.ValueOf(headers.Get(name)))
	case []byte:
		b, err := base64.StdEncoding.DecodeString(headers.Get(name))
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&b))
	case *bool:
		b, err := strconv.ParseBool(headers.Get(name))
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&b))
	case bool:
		b, err := strconv.ParseBool(headers.Get(name))
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(b))
	case *int64:
		i, err := strconv.ParseInt(headers.Get(name), 10, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&i))
	case int64:
		i, err := strconv.ParseInt(headers.Get(name), 10, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	case *float64:
		f, err := strconv.ParseFloat(headers.Get(name), 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&f))
	case float64:
		f, err := strconv.ParseFloat(headers.Get(name), 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(f))
	case *time.Time:
		format := tag.Get("timestampFormat")
		if len(format) == 0 {
			format = helper.RFC822TimeFormatName
			if tag.Get("location") == "querystring" {
				format = helper.ISO8601TimeFormatName
			}
		}
		t, err := helper.ParseTime(format, headers.Get(name))
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&t))
	case time.Time:
		format := tag.Get("timestampFormat")
		if len(format) == 0 {
			format = helper.RFC822TimeFormatName
			if tag.Get("location") == "querystring" {
				format = helper.ISO8601TimeFormatName
			}
		}
		t, err := helper.ParseTime(format, headers.Get(name))
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(t))
	case map[string]string: // we only support string map value types
		out := map[string]string{}
		for k, v := range headers {
			k = http.CanonicalHeaderKey(k)
			if strings.HasPrefix(strings.ToLower(k), strings.ToLower(name)) {
				out[k[len(name):]] = v[0]
			}
		}
		v.Set(reflect.ValueOf(out))
	default:
		err := fmt.Errorf("unsupported value for param %v (%s)", v.Interface(), v.Type())
		return err
	}
	return nil
}
