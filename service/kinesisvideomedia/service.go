// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideomedia

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client/metadata"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/signer/v4"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/private/protocol/restjson"
)

// KinesisVideoMedia provides the API operation methods for making requests to
// Amazon Kinesis Video Streams Media. See this package's package overview docs
// for details on the service.
//
// KinesisVideoMedia methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type KinesisVideoMedia struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "kinesisvideo"        // Name of service.
	EndpointsID = ServiceName           // ID to lookup a service endpoint with.
	ServiceID   = "Kinesis Video Media" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the KinesisVideoMedia client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a KinesisVideoMedia client from just a session.
//     svc := kinesisvideomedia.New(mySession)
//
//     // Create a KinesisVideoMedia client with additional configuration
//     svc := kinesisvideomedia.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *KinesisVideoMedia {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *KinesisVideoMedia {
	svc := &KinesisVideoMedia{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2017-09-30",
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

// newRequest creates a new request for a KinesisVideoMedia operation and runs any
// custom request initialization.
func (c *KinesisVideoMedia) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
