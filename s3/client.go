package s3

import (
	"net/http"

	. "github.com/uos-sdk-go/s3/client"
	"github.com/uos-sdk-go/s3/credential"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/log"
	"github.com/uos-sdk-go/s3/request"
	v4 "github.com/uos-sdk-go/s3/signature/v4"
)

type Client struct {
	Metadata    Metadata
	Config      helper.Config
	Credentials *credential.Credentials
	Handlers    request.Handlers
	Logger      *log.Logger
	HTTPClient  *http.Client
}

// NewClient creates a new client for request.
func NewClient(config helper.Config) *Client {
	client := &Client{
		Metadata: Metadata{
			ServiceName:   helper.ServiceName,
			ServiceID:     helper.ServiceID,
			APIVersion:    helper.APIVersion,
			SigningName:   helper.ServiceNameForSign,
			SigningRegion: helper.SDKSigningRegion,
		},
		Config:      config,
		Credentials: config.Credentials,
		Handlers:    request.NewHandlers(),
		Logger:      config.Logger,
		HTTPClient:  NewHttpClient(),
	}

	client.Handlers.Sign.PushBackHandlerItem(v4.SignV4Handler)
	client.Handlers.Sign.ForStopHandlers = request.StopHandlerListOnErr
	client.Handlers.Marshal.PushBackHandlerItem(request.MarshalRequestHandler)
	client.Handlers.Marshal.ForStopHandlers = request.StopHandlerListOnErr
	client.Handlers.Unmarshal.PushBackHandlerItem(request.UnmarshalRequestHandler)
	client.Handlers.Unmarshal.ForStopHandlers = request.StopHandlerListOnErr

	return client
}

// newRequest creates a new request for a client operation and runs any
// custom request initialization.
func (c *Client) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := request.NewRequest(c.Config, c.Metadata, c.Handlers, c.HTTPClient, op, params, data)

	return req
}
