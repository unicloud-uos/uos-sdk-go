package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	. "github.com/unicloud-uos/uos-sdk-go/s3/client"
	. "github.com/unicloud-uos/uos-sdk-go/s3/error"
	. "github.com/unicloud-uos/uos-sdk-go/s3/helper"
)

// Request is the service request to be made.
type Request struct {
	RequestID          string
	Config             Config
	Metadata           Metadata
	Handlers           Handlers
	Time               time.Time
	AttemptTime        time.Time
	ExpireTime         time.Duration
	Operation          *Operation
	HTTPRequest        *http.Request
	HTTPResponse       *http.Response
	HTTPClient         *http.Client
	SignedHeaderValues http.Header
	Body               io.ReadSeeker
	BodyStart          int64 // offset of Body that the request body starts
	Params             interface{}
	Error              error
	Data               interface{}
	build              bool
	LastSignedTime     time.Time
	// Need to persist an intermediate body between the input Body and HTTP
	// request body because the HTTP Client's transport can maintain a reference
	// to the HTTP request's body after the client has returned. This value is
	// safe to use concurrently and wrap the input Body for each HTTP request.
	safeBody *offsetReader
}

// An Operation is the service API operation to be made.
type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
}

// New returns a new Request pointer for the service API
// operation and parameters.
func NewRequest(cfg Config, metadata Metadata, handlers Handlers, client *http.Client,
	operation *Operation, params interface{}, data interface{}) *Request {

	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	request, _ := http.NewRequest(method, "", nil)

	var err error
	endpoint := parseEndpoint(cfg.Endpoint, cfg.DisableSSL)

	cfg.Logger.Debug("Endpoint in NewRequest: ", endpoint)
	request.URL, err = url.Parse(endpoint + operation.HTTPPath)
	if err != nil {
		request.URL = &url.URL{}
		err = NewBaseError("InvalidEndpointURL", "invalid endpoint uri", err)
	}

	cfg.Logger.Debug("URL in NewRequest: ", request.URL)
	ResetHostForHeader(request)
	if !cfg.RedirectEnabled {
		DisableHTTPRedirect(client)
	}

	r := &Request{
		Config:      cfg,
		Metadata:    metadata,
		Handlers:    handlers,
		Time:        time.Now(),
		Operation:   operation,
		HTTPRequest: request,
		HTTPClient:  client,
		Params:      params,
		Data:        data,
		Error:       err,
	}
	r.SetBufferBody([]byte{})

	return r
}

func (r *Request) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("context cannot be nil")
	}
	r.HTTPRequest = r.HTTPRequest.WithContext(ctx)
}

func (r *Request) SetBufferBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

// SetReaderBody will set the request's body reader.
func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.Body = reader
	var err error
	// Get the Bodies current offset so retries will start from the same
	// initial position.
	r.BodyStart, err = reader.Seek(0, io.SeekCurrent)
	if err != nil {
		r.Error = NewBaseError("SerializationError", "failed to determine start of request body", err)
		return
	}
	r.ResetBody()
}

// ResetBody rewinds the request body back to its starting position, and
// set's the HTTP Request body reference. When the body is read prior
// to being sent in the HTTP request it will need to be rewound.
//
// ResetBody will automatically be called by the SDK's build handler, but if
// the request is being used directly ResetBody must be called before the request
// is Sent.  SetStringBody, SetBufferBody, and SetReaderBody will automatically
// call ResetBody.
//
func (r *Request) ResetBody() {
	body, err := r.getNextRequestBody()
	if err != nil {
		r.Error = NewBaseError("SerializationError", "failed to reset request body", err)
		return
	}

	r.HTTPRequest.Body = body
	r.HTTPRequest.GetBody = r.getNextRequestBody
}

func (r *Request) getNextRequestBody() (body io.ReadCloser, err error) {
	if r.safeBody != nil {
		r.safeBody.Close()
	}

	r.safeBody, err = newOffsetReader(r.Body, r.BodyStart)
	if err != nil {
		return nil, NewBaseError("SerializationError", "failed to get request body error", err)
	}

	l, err := seekerLen(r.Body)
	if err != nil {
		return nil, NewBaseError("SerializationError", "failed to compute request body size", err)
	}

	if l == 0 {
		body = http.NoBody
	} else if l > 0 {
		body = r.safeBody
	} else {
		// Hack to prevent sending bodies for methods where the body
		// should be ignored by the server. Sending bodies on these
		// methods without an associated ContentLength will cause the
		// request to socket timeout because the server does not handle
		// Transfer-Encoding: chunked bodies for these methods.
		//
		// This would only happen if a ReaderSeekerCloser was used with
		// a io.Reader that was not also an io.Seeker, or did not
		// implement Len() method.
		switch r.Operation.HTTPMethod {
		case "GET", "HEAD", "DELETE":
			body = http.NoBody
		default:
			body = r.safeBody
		}
	}

	return body, nil
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
	if r.Error != nil {
		return r.Error
	}

	r.Error = nil
	r.AttemptTime = time.Now()

	if err := r.Sign(); err != nil {
		reqDebugLog(r, "Sign Request", err)
		return err
	}

	if err := r.doRequest(); err != nil {
		return err
	}
	return nil
}

func (r *Request) Sign() error {
	r.Build()
	if r.Error != nil {
		reqDebugLog(r, "Build Request", r.Error)
		return r.Error
	}

	r.Handlers.Sign.Run(r)
	return r.Error
}

// GetBody will return an io.ReadSeeker of the Request's underlying
// input body with a concurrency safe wrapper.
func (r *Request) GetBody() io.ReadSeeker {
	return r.safeBody
}

// Build validate parameters and build request's object
func (r *Request) Build() error {
	if !r.build {
		r.Handlers.Validate.Run(r)
		if r.Error != nil {
			reqDebugLog(r, "Validate Request", r.Error)
			return r.Error
		}
		r.Handlers.Marshal.Run(r)
		if r.Error != nil {
			reqDebugLog(r, "Marshal Request", r.Error)
			return r.Error
		}
		r.Handlers.Set.Run(r)
		if r.Error != nil {
			reqDebugLog(r, "Set Request", r.Error)
			return r.Error
		}
		r.build = true
	}

	return r.Error
}

func (r *Request) doRequest() (sendErr error) {
	r.Handlers.Send.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.Unmarshal.Run(r)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func reqDebugLog(r *Request, stage string, err error) {
	r.Config.Logger.Debug(fmt.Sprintf("%s %s/%s failed, error %v",
		stage, r.Metadata.ServiceName, r.Operation.Name, err))
}

// check parameters are valid or not
// is no parameter or invalid, return false
func (r *Request) IsParamsValid() bool {
	return r.Params != nil && reflect.ValueOf(r.Params).Elem().IsValid()
}

func (r *Request) AddAgentInfo(s string) {
	r.HTTPRequest.Header.Set("User-Agent", s)
}

func (r *Request) GetResponseHeader(s string) string {
	return r.HTTPResponse.Header.Get(s)
}

// Presign returns the request's signed URL. Error will be returned
// if the signing fails.
//
// It is invalid to create a presigned URL with a expire duration 0 or less. An
// error is returned if expire duration is 0 or less.
func (r *Request) Presign(expire time.Duration) (string, error) {
	if expire <= 0 {
		return "", NewBaseError(
			"InvalidPresignExpireError",
			"presigned URL requires an expire duration greater than 0",
			nil,
		)
	}

	r.ExpireTime = expire

	if err := r.Sign(); err != nil {
		return "", err
	}

	return r.HTTPRequest.URL.String(), nil
}

// offsetReader is a thread-safe io.ReadCloser to prevent racing
// with retrying requests
type offsetReader struct {
	buf    io.ReadSeeker
	lock   sync.Mutex
	closed bool
}

func newOffsetReader(buf io.ReadSeeker, offset int64) (*offsetReader, error) {
	reader := &offsetReader{}
	_, err := buf.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}
	reader.buf = buf
	return reader, nil
}

// Close will close the instance of the offset reader's access to
// the underlying io.ReadSeeker.
func (o *offsetReader) Close() error {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.closed = true
	return nil
}

// Read is a thread-safe read of the underlying io.ReadSeeker
func (o *offsetReader) Read(p []byte) (int, error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if o.closed {
		return 0, io.EOF
	}

	return o.buf.Read(p)
}

// Seek is a thread-safe seeking operation.
func (o *offsetReader) Seek(offset int64, whence int) (int64, error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	return o.buf.Seek(offset, whence)
}

// CloseAndCopy will return a new offsetReader with a copy of the old buffer
// and close the old buffer.
func (o *offsetReader) CloseAndCopy(offset int64) (*offsetReader, error) {
	if err := o.Close(); err != nil {
		return nil, err
	}
	return newOffsetReader(o.buf, offset)
}

func seekerLen(s io.Seeker) (int64, error) {
	curOffset, err := s.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	endOffset, err := s.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	_, err = s.Seek(curOffset, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return endOffset - curOffset, nil
}

func parseEndpoint(endpoint string, disableSSL bool) string {
	strs := strings.Split(endpoint, "://")

	if disableSSL {
		return "http://" + strs[len(strs)-1]
	} else if strs[0] == "http" || strs[0] == "https" {
		return endpoint
	} else {
		return "https://" + endpoint
	}
}
