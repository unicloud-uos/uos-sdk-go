// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisvideoiface provides an interface to enable mocking the Amazon Kinesis Video Streams service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisvideoiface

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/kinesisvideo"
)

// KinesisVideoAPI provides an interface to enable mocking the
// kinesisvideo.KinesisVideo service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Video Streams.
//    func myFunc(svc kinesisvideoiface.KinesisVideoAPI) bool {
//        // Make svc.CreateStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kinesisvideo.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisVideoClient struct {
//        kinesisvideoiface.KinesisVideoAPI
//    }
//    func (m *mockKinesisVideoClient) CreateStream(input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisVideoClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KinesisVideoAPI interface {
	CreateStream(*kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error)
	CreateStreamWithContext(aws.Context, *kinesisvideo.CreateStreamInput, ...request.Option) (*kinesisvideo.CreateStreamOutput, error)
	CreateStreamRequest(*kinesisvideo.CreateStreamInput) (*request.Request, *kinesisvideo.CreateStreamOutput)

	DeleteStream(*kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error)
	DeleteStreamWithContext(aws.Context, *kinesisvideo.DeleteStreamInput, ...request.Option) (*kinesisvideo.DeleteStreamOutput, error)
	DeleteStreamRequest(*kinesisvideo.DeleteStreamInput) (*request.Request, *kinesisvideo.DeleteStreamOutput)

	DescribeStream(*kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error)
	DescribeStreamWithContext(aws.Context, *kinesisvideo.DescribeStreamInput, ...request.Option) (*kinesisvideo.DescribeStreamOutput, error)
	DescribeStreamRequest(*kinesisvideo.DescribeStreamInput) (*request.Request, *kinesisvideo.DescribeStreamOutput)

	GetDataEndpoint(*kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error)
	GetDataEndpointWithContext(aws.Context, *kinesisvideo.GetDataEndpointInput, ...request.Option) (*kinesisvideo.GetDataEndpointOutput, error)
	GetDataEndpointRequest(*kinesisvideo.GetDataEndpointInput) (*request.Request, *kinesisvideo.GetDataEndpointOutput)

	ListStreams(*kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error)
	ListStreamsWithContext(aws.Context, *kinesisvideo.ListStreamsInput, ...request.Option) (*kinesisvideo.ListStreamsOutput, error)
	ListStreamsRequest(*kinesisvideo.ListStreamsInput) (*request.Request, *kinesisvideo.ListStreamsOutput)

	ListTagsForStream(*kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error)
	ListTagsForStreamWithContext(aws.Context, *kinesisvideo.ListTagsForStreamInput, ...request.Option) (*kinesisvideo.ListTagsForStreamOutput, error)
	ListTagsForStreamRequest(*kinesisvideo.ListTagsForStreamInput) (*request.Request, *kinesisvideo.ListTagsForStreamOutput)

	TagStream(*kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error)
	TagStreamWithContext(aws.Context, *kinesisvideo.TagStreamInput, ...request.Option) (*kinesisvideo.TagStreamOutput, error)
	TagStreamRequest(*kinesisvideo.TagStreamInput) (*request.Request, *kinesisvideo.TagStreamOutput)

	UntagStream(*kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error)
	UntagStreamWithContext(aws.Context, *kinesisvideo.UntagStreamInput, ...request.Option) (*kinesisvideo.UntagStreamOutput, error)
	UntagStreamRequest(*kinesisvideo.UntagStreamInput) (*request.Request, *kinesisvideo.UntagStreamOutput)

	UpdateDataRetention(*kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error)
	UpdateDataRetentionWithContext(aws.Context, *kinesisvideo.UpdateDataRetentionInput, ...request.Option) (*kinesisvideo.UpdateDataRetentionOutput, error)
	UpdateDataRetentionRequest(*kinesisvideo.UpdateDataRetentionInput) (*request.Request, *kinesisvideo.UpdateDataRetentionOutput)

	UpdateStream(*kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error)
	UpdateStreamWithContext(aws.Context, *kinesisvideo.UpdateStreamInput, ...request.Option) (*kinesisvideo.UpdateStreamOutput, error)
	UpdateStreamRequest(*kinesisvideo.UpdateStreamInput) (*request.Request, *kinesisvideo.UpdateStreamOutput)
}

var _ KinesisVideoAPI = (*kinesisvideo.KinesisVideo)(nil)
