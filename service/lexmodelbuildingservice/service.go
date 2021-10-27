// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/client"
	"github.com/unicloud-uos/uos-sdk-go/aws/client/metadata"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/aws/signer/v4"
	"github.com/unicloud-uos/uos-sdk-go/private/protocol/restjson"
)

// LexModelBuildingService provides the API operation methods for making requests to
// Amazon Lex Model Building Service. See this package's package overview docs
// for details on the service.
//
// LexModelBuildingService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type LexModelBuildingService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "models.lex"                 // Name of service.
	EndpointsID = ServiceName                  // ID to lookup a service endpoint with.
	ServiceID   = "Lex Model Building Service" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the LexModelBuildingService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a LexModelBuildingService client from just a session.
//     svc := lexmodelbuildingservice.New(mySession)
//
//     // Create a LexModelBuildingService client with additional configuration
//     svc := lexmodelbuildingservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *LexModelBuildingService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "lex"
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *LexModelBuildingService {
	svc := &LexModelBuildingService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2017-04-19",
				JSONVersion:   "1.1",
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

// newRequest creates a new request for a LexModelBuildingService operation and runs any
// custom request initialization.
func (c *LexModelBuildingService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
