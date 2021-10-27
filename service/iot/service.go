// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client/metadata"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/signer/v4"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/private/protocol/restjson"
)

// IoT provides the API operation methods for making requests to
// AWS IoT. See this package's package overview docs
// for details on the service.
//
// IoT methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type IoT struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "iot"       // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "IoT"       // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the IoT client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a IoT client from just a session.
//     svc := iot.New(mySession)
//
//     // Create a IoT client with additional configuration
//     svc := iot.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *IoT {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "execute-api"
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *IoT {
	svc := &IoT{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-05-28",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a IoT operation and runs any
// custom request initialization.
func (c *IoT) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
