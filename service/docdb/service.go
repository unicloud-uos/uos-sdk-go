// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/client"
	"github.com/unicloud-uos/uos-sdk-go/aws/client/metadata"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/aws/signer/v4"
	"github.com/unicloud-uos/uos-sdk-go/private/protocol/query"
)

// DocDB provides the API operation methods for making requests to
// Amazon DocumentDB with MongoDB compatibility. See this package's package overview docs
// for details on the service.
//
// DocDB methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DocDB struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "DocDB" // Name of service.
	EndpointsID = "rds"   // ID to lookup a service endpoint with.
	ServiceID   = "DocDB" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the DocDB client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a DocDB client from just a session.
//     svc := docdb.New(mySession)
//
//     // Create a DocDB client with additional configuration
//     svc := docdb.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DocDB {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "rds"
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DocDB {
	svc := &DocDB{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-10-31",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DocDB operation and runs any
// custom request initialization.
func (c *DocDB) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
