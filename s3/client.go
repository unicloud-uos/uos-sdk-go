package s3

import (
	"github.com/uos-sdk-go/s3/log"
	"net/http"

	"github.com/uos-sdk-go/s3/auxiliary"
	. "github.com/uos-sdk-go/s3/client"
	"github.com/uos-sdk-go/s3/credential"
	"github.com/uos-sdk-go/s3/handler"
	"github.com/uos-sdk-go/s3/request"
)

type Client struct {
	Metadata    Metadata
	Config      auxiliary.Config
	Credentials *credential.Credentials
	Handlers    handler.Handlers
	Logger      *log.Logger

	HTTPClient *http.Client
}

// NewClient creates a new client for request.
func NewClient(config auxiliary.Config) *Client {
	client := &Client{
		Metadata: Metadata{
			ServiceName:   auxiliary.ServiceName,
			ServiceID:     auxiliary.ServiceID,
			APIVersion:    auxiliary.APIVersion,
			SigningName:   auxiliary.ServiceNameForSign,
			SigningRegion: auxiliary.SDKSigningRegion,
		},
		Config:      config,
		Credentials: config.Credentials,
		Handlers: handler.
			Logger:      config.Logger,
	}

	return client
}

// newRequest creates a new request for a client operation and runs any
// custom request initialization.
func (c *Client) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := request.NewRequest(c.Config, c.Metadata, op, params, data)

	return req
}
