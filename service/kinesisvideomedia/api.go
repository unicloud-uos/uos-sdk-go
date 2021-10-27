// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideomedia

import (
	"io"
	"time"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awsutil"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

const opGetMedia = "GetMedia"

// GetMediaRequest generates a "aws/request.Request" representing the
// client's request for the GetMedia operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetMedia for more information on using the GetMedia
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetMediaRequest method.
//    req, resp := client.GetMediaRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-media-2017-09-30/GetMedia
func (c *KinesisVideoMedia) GetMediaRequest(input *GetMediaInput) (req *request.Request, output *GetMediaOutput) {
	op := &request.Operation{
		Name:       opGetMedia,
		HTTPMethod: "POST",
		HTTPPath:   "/getMedia",
	}

	if input == nil {
		input = &GetMediaInput{}
	}

	output = &GetMediaOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetMedia API operation for Amazon Kinesis Video Streams Media.
//
// Use this API to retrieve media content from a Kinesis video stream. In the
// request, you identify the stream name or stream Amazon Resource Name (ARN),
// and the starting chunk. Kinesis Video Streams then returns a stream of chunks
// in order by fragment number.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send
// the GetMedia requests to this endpoint using the --endpoint-url parameter
// (https://docs.aws.amazon.com/cli/latest/reference/).
//
// When you put media data (fragments) on a stream, Kinesis Video Streams stores
// each incoming fragment and related metadata in what is called a "chunk."
// For more information, see . The GetMedia API returns a stream of these chunks
// starting from the chunk that you specify in the request.
//
// The following limits apply when using the GetMedia API:
//
//    * A client can call GetMedia up to five times per second per stream.
//
//    * Kinesis Video Streams sends media data at a rate of up to 25 megabytes
//    per second (or 200 megabits per second) during a GetMedia session.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Kinesis Video Streams Media's
// API operation GetMedia for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   Status Code: 404, The stream with the given name does not exist.
//
//   * ErrCodeNotAuthorizedException "NotAuthorizedException"
//   Status Code: 403, The caller is not authorized to perform an operation on
//   the given stream, or the token has expired.
//
//   * ErrCodeInvalidEndpointException "InvalidEndpointException"
//   Status Code: 400, Caller used wrong endpoint to write data to a stream. On
//   receiving such an exception, the user must call GetDataEndpoint with AccessMode
//   set to "READ" and use the endpoint Kinesis Video returns in the next GetMedia
//   call.
//
//   * ErrCodeClientLimitExceededException "ClientLimitExceededException"
//   Kinesis Video Streams has throttled the request because you have exceeded
//   the limit of allowed client calls. Try making the call later.
//
//   * ErrCodeConnectionLimitExceededException "ConnectionLimitExceededException"
//   Kinesis Video Streams has throttled the request because you have exceeded
//   the limit of allowed client connections.
//
//   * ErrCodeInvalidArgumentException "InvalidArgumentException"
//   The value for this input parameter is invalid.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-media-2017-09-30/GetMedia
func (c *KinesisVideoMedia) GetMedia(input *GetMediaInput) (*GetMediaOutput, error) {
	req, out := c.GetMediaRequest(input)
	return out, req.Send()
}

// GetMediaWithContext is the same as GetMedia with the addition of
// the ability to pass a context and additional request options.
//
// See GetMedia for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KinesisVideoMedia) GetMediaWithContext(ctx aws.Context, input *GetMediaInput, opts ...request.Option) (*GetMediaOutput, error) {
	req, out := c.GetMediaRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetMediaInput struct {
	_ struct{} `type:"structure"`

	// Identifies the starting chunk to get from the specified stream.
	//
	// StartSelector is a required field
	StartSelector *StartSelector `type:"structure" required:"true"`

	// The ARN of the stream from where you want to get the media content. If you
	// don't specify the streamARN, you must specify the streamName.
	StreamARN *string `min:"1" type:"string"`

	// The Kinesis video stream name from where you want to get the media content.
	// If you don't specify the streamName, you must specify the streamARN.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetMediaInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMediaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetMediaInput"}
	if s.StartSelector == nil {
		invalidParams.Add(request.NewErrParamRequired("StartSelector"))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("StreamName", 1))
	}
	if s.StartSelector != nil {
		if err := s.StartSelector.Validate(); err != nil {
			invalidParams.AddNested("StartSelector", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetStartSelector sets the StartSelector field's value.
func (s *GetMediaInput) SetStartSelector(v *StartSelector) *GetMediaInput {
	s.StartSelector = v
	return s
}

// SetStreamARN sets the StreamARN field's value.
func (s *GetMediaInput) SetStreamARN(v string) *GetMediaInput {
	s.StreamARN = &v
	return s
}

// SetStreamName sets the StreamName field's value.
func (s *GetMediaInput) SetStreamName(v string) *GetMediaInput {
	s.StreamName = &v
	return s
}

type GetMediaOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The content type of the requested media.
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// The payload Kinesis Video Streams returns is a sequence of chunks from the
	// specified stream. For information about the chunks, see . The chunks that
	// Kinesis Video Streams returns in the GetMedia call also include the following
	// additional Matroska (MKV) tags:
	//
	//    * AWS_KINESISVIDEO_CONTINUATION_TOKEN (UTF-8 string) - In the event your
	//    GetMedia call terminates, you can use this continuation token in your
	//    next request to get the next chunk where the last request terminated.
	//
	//    * AWS_KINESISVIDEO_MILLIS_BEHIND_NOW (UTF-8 string) - Client applications
	//    can use this tag value to determine how far behind the chunk returned
	//    in the response is from the latest chunk on the stream.
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	//    * AWS_KINESISVIDEO_SERVER_TIMESTAMP - Server timestamp of the fragment.
	//
	//    * AWS_KINESISVIDEO_PRODUCER_TIMESTAMP - Producer timestamp of the fragment.
	//
	// The following tags will be present if an error occurs:
	//
	//    * AWS_KINESISVIDEO_ERROR_CODE - String description of an error that caused
	//    GetMedia to stop.
	//
	//    * AWS_KINESISVIDEO_ERROR_ID: Integer code of the error.
	//
	// The error codes are as follows:
	//
	//    * 3002 - Error writing to the stream
	//
	//    * 4000 - Requested fragment is not found
	//
	//    * 4500 - Access denied for the stream's KMS key
	//
	//    * 4501 - Stream's KMS key is disabled
	//
	//    * 4502 - Validation error on the stream's KMS key
	//
	//    * 4503 - KMS key specified in the stream is unavailable
	//
	//    * 4504 - Invalid usage of the KMS key specified in the stream
	//
	//    * 4505 - Invalid state of the KMS key specified in the stream
	//
	//    * 4506 - Unable to find the KMS key specified in the stream
	//
	//    * 5000 - Internal error
	Payload io.ReadCloser `type:"blob"`
}

// String returns the string representation
func (s GetMediaOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaOutput) GoString() string {
	return s.String()
}

// SetContentType sets the ContentType field's value.
func (s *GetMediaOutput) SetContentType(v string) *GetMediaOutput {
	s.ContentType = &v
	return s
}

// SetPayload sets the Payload field's value.
func (s *GetMediaOutput) SetPayload(v io.ReadCloser) *GetMediaOutput {
	s.Payload = v
	return s
}

// Identifies the chunk on the Kinesis video stream where you want the GetMedia
// API to start returning media data. You have the following options to identify
// the starting chunk:
//
//    * Choose the latest (or oldest) chunk.
//
//    * Identify a specific chunk. You can identify a specific chunk either
//    by providing a fragment number or timestamp (server or producer).
//
//    * Each chunk's metadata includes a continuation token as a Matroska (MKV)
//    tag (AWS_KINESISVIDEO_CONTINUATION_TOKEN). If your previous GetMedia request
//    terminated, you can use this tag value in your next GetMedia request.
//    The API then starts returning chunks starting where the last API ended.
type StartSelector struct {
	_ struct{} `type:"structure"`

	// Specifies the fragment number from where you want the GetMedia API to start
	// returning the fragments.
	AfterFragmentNumber *string `min:"1" type:"string"`

	// Continuation token that Kinesis Video Streams returned in the previous GetMedia
	// response. The GetMedia API then starts with the chunk identified by the continuation
	// token.
	ContinuationToken *string `min:"1" type:"string"`

	// Identifies the fragment on the Kinesis video stream where you want to start
	// getting the data from.
	//
	//    * NOW - Start with the latest chunk on the stream.
	//
	//    * EARLIEST - Start with earliest available chunk on the stream.
	//
	//    * FRAGMENT_NUMBER - Start with the chunk containing the specific fragment.
	//    You must also specify the StartFragmentNumber.
	//
	//    * PRODUCER_TIMESTAMP or SERVER_TIMESTAMP - Start with the chunk containing
	//    a fragment with the specified producer or server timestamp. You specify
	//    the timestamp by adding StartTimestamp.
	//
	//    *  CONTINUATION_TOKEN - Read using the specified continuation token.
	//
	// If you choose the NOW, EARLIEST, or CONTINUATION_TOKEN as the startSelectorType,
	// you don't provide any additional information in the startSelector.
	//
	// StartSelectorType is a required field
	StartSelectorType *string `type:"string" required:"true" enum:"StartSelectorType"`

	// A timestamp value. This value is required if you choose the PRODUCER_TIMESTAMP
	// or the SERVER_TIMESTAMP as the startSelectorType. The GetMedia API then starts
	// with the chunk containing the fragment that has the specified timestamp.
	StartTimestamp *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s StartSelector) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSelector) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSelector) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartSelector"}
	if s.AfterFragmentNumber != nil && len(*s.AfterFragmentNumber) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("AfterFragmentNumber", 1))
	}
	if s.ContinuationToken != nil && len(*s.ContinuationToken) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ContinuationToken", 1))
	}
	if s.StartSelectorType == nil {
		invalidParams.Add(request.NewErrParamRequired("StartSelectorType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAfterFragmentNumber sets the AfterFragmentNumber field's value.
func (s *StartSelector) SetAfterFragmentNumber(v string) *StartSelector {
	s.AfterFragmentNumber = &v
	return s
}

// SetContinuationToken sets the ContinuationToken field's value.
func (s *StartSelector) SetContinuationToken(v string) *StartSelector {
	s.ContinuationToken = &v
	return s
}

// SetStartSelectorType sets the StartSelectorType field's value.
func (s *StartSelector) SetStartSelectorType(v string) *StartSelector {
	s.StartSelectorType = &v
	return s
}

// SetStartTimestamp sets the StartTimestamp field's value.
func (s *StartSelector) SetStartTimestamp(v time.Time) *StartSelector {
	s.StartTimestamp = &v
	return s
}

const (
	// StartSelectorTypeFragmentNumber is a StartSelectorType enum value
	StartSelectorTypeFragmentNumber = "FRAGMENT_NUMBER"

	// StartSelectorTypeServerTimestamp is a StartSelectorType enum value
	StartSelectorTypeServerTimestamp = "SERVER_TIMESTAMP"

	// StartSelectorTypeProducerTimestamp is a StartSelectorType enum value
	StartSelectorTypeProducerTimestamp = "PRODUCER_TIMESTAMP"

	// StartSelectorTypeNow is a StartSelectorType enum value
	StartSelectorTypeNow = "NOW"

	// StartSelectorTypeEarliest is a StartSelectorType enum value
	StartSelectorTypeEarliest = "EARLIEST"

	// StartSelectorTypeContinuationToken is a StartSelectorType enum value
	StartSelectorTypeContinuationToken = "CONTINUATION_TOKEN"
)
