package request

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/uos-sdk-go/s3/auxiliary"
	"github.com/uos-sdk-go/s3/client"
	. "github.com/uos-sdk-go/s3/error"
)

// Request is the service request to be made.
type Request struct {
	ID                 string
	Config             auxiliary.Config
	Metadata           client.Metadata
	Time               time.Time
	AttemptTime        time.Time
	ExpireTime         time.Duration
	Operation          *Operation
	HTTPRequest        *http.Request
	HTTPResponse       *http.Response
	SignedHeaderValues http.Header
	Body               io.ReadSeeker
	BodyStart          int64 // offset of Body that the request body starts
	Params             interface{}
	Error              error
	Data               interface{}
	build              bool
}

// An Operation is the service API operation to be made.
type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
}

// New returns a new Request pointer for the service API
// operation and parameters.
func NewRequest(cfg auxiliary.Config, metadata client.Metadata,
	operation *Operation, params interface{}, data interface{}) *Request {

	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	request, _ := http.NewRequest(method, "", nil)

	var err error
	request.URL, err = url.Parse(cfg.Endpoint + operation.HTTPPath)
	if err != nil {
		request.URL = &url.URL{}
		err = NewBaseError("InvalidEndpointURL", "invalid endpoint uri", err)
	}

	ResetHostForHeader(request)

	r := &Request{
		Config:      cfg,
		Metadata:    metadata,
		Time:        time.Now(),
		Operation:   operation,
		HTTPRequest: request,
		Params:      params,
		Data:        data,
	}

	return r
}

func (r *Request) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("context cannot be nil")
	}
	r.HTTPRequest = r.HTTPRequest.WithContext(ctx)
}

// ResetHostForHeader removes replace default port from host and updates request.Host
func ResetHostForHeader(r *http.Request) {
	host := getHost(r)
	port := getPort(host)
	if port != "" && isDefaultPort(r.URL.Scheme, port) {
		r.Host = stripPort(host)
	}
}

func getHost(r *http.Request) string {
	if r.Host != "" {
		return r.Host
	}

	return r.URL.Host
}

// stripPort return host without port
func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}

// getPort get port
func getPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return ""
	}
	if i := strings.Index(hostport, "]:"); i != -1 {
		return hostport[i+len("]:"):]
	}
	if strings.Contains(hostport, "]") {
		return ""
	}
	return hostport[colon+len(":"):]
}

func isDefaultPort(scheme, port string) bool {
	if port == "" {
		return true
	}

	if (scheme == "http" && port == "80") || (scheme == "https" && port == "443") {
		return true
	}

	return false
}

func (r *Request) Do() error {
	defer func() {
		// Regardless of success or failure of the request trigger the Complete
		// request handlers.
		r.Handlers.Complete.Run(r)
	}()

	if err := r.Error; err != nil {
		return err
	}

	for {
		r.Error = nil
		r.AttemptTime = time.Now()

		if err := r.Sign(); err != nil {
			reqDebugLog(r, "Sign Request", err)
			return err
		}

		if err := r.sendRequest(); err == nil {
			return nil
		} else if !shouldRetryCancel(r.Error) {
			return err
		} else {
			r.Handlers.Retry.Run(r)
			r.Handlers.AfterRetry.Run(r)

			if r.Error != nil || !aws.BoolValue(r.Retryable) {
				return r.Error
			}

			r.prepareRetry()
			continue
		}
	}
}

func (r *Request) Sign() error {
	r.Build()
	if r.Error != nil {
		r.Config.Logger.Debug()
		reqDebugLog(r, "Build Request", r.Error)
		return r.Error
	}

	r.Handlers.Sign.Run(r)
	return r.Error
}

// Build validate parameters and build request's object
func (r *Request) Build() error {
	if !r.build {
		r.Handlers.Validate.Run(r)
		if r.Error != nil {
			reqDebugLog(r, "Validate Request", r.Error)
			return r.Error
		}
		r.Handlers.Build.Run(r)
		if r.Error != nil {
			reqDebugLog(r, "Build Request", r.Error)
			return r.Error
		}
		r.build = true
	}

	return r.Error
}

func (r *Request) doRequest() (sendErr error) {
	defer r.Handlers.CompleteAttempt.Run(r)
	defer func() {
		if r.Error != nil {
			r.Handlers.ShouldRetry.Run(r)
		}
	}()

	r.Handlers.Send.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.UnmarshalMeta.Run(r)
	r.Handlers.ValidateResponse.Run(r)
	if r.Error != nil {
		r.Handlers.UnmarshalError.Run(r)
		return r.Error
	}

	r.Handlers.Unmarshal.Run(r)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func reqDebugLog(r *Request, stage string, err error) {
	r.Config.Logger.Debug("%s %s/%s failed, error %v",
		stage, r.Metadata.ServiceName, r.Operation.Name, err)
}

// check parameters are valid or not
// is no parameter or invalid, return false
func (r *Request) IsParamsValid() bool {
	return r.Params != nil && reflect.ValueOf(r.Params).Elem().IsValid()
}

func (r *Request) AddAgentInfo(s string) {
	r.HTTPRequest.Header.Set("AgentInfo", s)
}
