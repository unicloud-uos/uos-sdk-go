// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/awsutil"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
)

const opInvokeEndpoint = "InvokeEndpoint"

// InvokeEndpointRequest generates a "aws/request.Request" representing the
// client's request for the InvokeEndpoint operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See InvokeEndpoint for more information on using the InvokeEndpoint
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the InvokeEndpointRequest method.
//    req, resp := client.InvokeEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *SageMakerRuntime) InvokeEndpointRequest(input *InvokeEndpointInput) (req *request.Request, output *InvokeEndpointOutput) {
	op := &request.Operation{
		Name:       opInvokeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/endpoints/{EndpointName}/invocations",
	}

	if input == nil {
		input = &InvokeEndpointInput{}
	}

	output = &InvokeEndpointOutput{}
	req = c.newRequest(op, input, output)
	return
}

// InvokeEndpoint API operation for Amazon SageMaker Runtime.
//
// After you deploy a model into production using Amazon SageMaker hosting services,
// your client applications use this API to get inferences from the model hosted
// at the specified endpoint.
//
// For an overview of Amazon SageMaker, see How It Works (http://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html).
//
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Cals to InvokeEndpoint are authenticated by using AWS Signature Version 4.
// For information, see Authenticating Requests (AWS Signature Version 4) (http://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference.
//
// Endpoints are scoped to an individual account, and are not public. The URL
// does not contain the account ID, but Amazon SageMaker determines the account
// ID from the authentication token that is supplied by the caller.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon SageMaker Runtime's
// API operation InvokeEndpoint for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalFailure "InternalFailure"
//   An internal failure occurred.
//
//   * ErrCodeServiceUnavailable "ServiceUnavailable"
//   The service is unavailable. Try your call again.
//
//   * ErrCodeValidationError "ValidationError"
//   Inspect your request and try again.
//
//   * ErrCodeModelError "ModelError"
//   Model (owned by the customer in the container) returned an error 500.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/runtime.sagemaker-2017-05-13/InvokeEndpoint
func (c *SageMakerRuntime) InvokeEndpoint(input *InvokeEndpointInput) (*InvokeEndpointOutput, error) {
	req, out := c.InvokeEndpointRequest(input)
	return out, req.Send()
}

// InvokeEndpointWithContext is the same as InvokeEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See InvokeEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SageMakerRuntime) InvokeEndpointWithContext(ctx aws.Context, input *InvokeEndpointInput, opts ...request.Option) (*InvokeEndpointOutput, error) {
	req, out := c.InvokeEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type InvokeEndpointInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The desired MIME type of the inference in the response.
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model.
	//
	// For information about the format of the request body, see Common Data Formats—Inference
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the input data in the request body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// The name of the endpoint that you specified when you created the endpoint
	// using the CreateEndpoint (http://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html)
	// API.
	//
	// EndpointName is a required field
	EndpointName *string `location:"uri" locationName:"EndpointName" type:"string" required:"true"`
}

// String returns the string representation
func (s InvokeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "InvokeEndpointInput"}
	if s.Body == nil {
		invalidParams.Add(request.NewErrParamRequired("Body"))
	}
	if s.EndpointName == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointName"))
	}
	if s.EndpointName != nil && len(*s.EndpointName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EndpointName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccept sets the Accept field's value.
func (s *InvokeEndpointInput) SetAccept(v string) *InvokeEndpointInput {
	s.Accept = &v
	return s
}

// SetBody sets the Body field's value.
func (s *InvokeEndpointInput) SetBody(v []byte) *InvokeEndpointInput {
	s.Body = v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *InvokeEndpointInput) SetContentType(v string) *InvokeEndpointInput {
	s.ContentType = &v
	return s
}

// SetCustomAttributes sets the CustomAttributes field's value.
func (s *InvokeEndpointInput) SetCustomAttributes(v string) *InvokeEndpointInput {
	s.CustomAttributes = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *InvokeEndpointInput) SetEndpointName(v string) *InvokeEndpointInput {
	s.EndpointName = &v
	return s
}

type InvokeEndpointOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Includes the inference provided by the model.
	//
	// For information about the format of the response body, see Common Data Formats—Inference
	// (http://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// Body is a required field
	Body []byte `type:"blob" required:"true" sensitive:"true"`

	// The MIME type of the inference returned in the response body.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	CustomAttributes *string `location:"header" locationName:"X-Amzn-SageMaker-Custom-Attributes" type:"string" sensitive:"true"`

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string `location:"header" locationName:"x-Amzn-Invoked-Production-Variant" type:"string"`
}

// String returns the string representation
func (s InvokeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeEndpointOutput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *InvokeEndpointOutput) SetBody(v []byte) *InvokeEndpointOutput {
	s.Body = v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *InvokeEndpointOutput) SetContentType(v string) *InvokeEndpointOutput {
	s.ContentType = &v
	return s
}

// SetCustomAttributes sets the CustomAttributes field's value.
func (s *InvokeEndpointOutput) SetCustomAttributes(v string) *InvokeEndpointOutput {
	s.CustomAttributes = &v
	return s
}

// SetInvokedProductionVariant sets the InvokedProductionVariant field's value.
func (s *InvokeEndpointOutput) SetInvokedProductionVariant(v string) *InvokeEndpointOutput {
	s.InvokedProductionVariant = &v
	return s
}
