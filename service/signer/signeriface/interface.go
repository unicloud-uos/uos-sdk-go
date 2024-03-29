// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package signeriface provides an interface to enable mocking the AWS Signer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package signeriface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/signer"
)

// SignerAPI provides an interface to enable mocking the
// signer.Signer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Signer.
//    func myFunc(svc signeriface.SignerAPI) bool {
//        // Make svc.CancelSigningProfile request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := signer.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSignerClient struct {
//        signeriface.SignerAPI
//    }
//    func (m *mockSignerClient) CancelSigningProfile(input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSignerClient{}
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
type SignerAPI interface {
	CancelSigningProfile(*signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error)
	CancelSigningProfileWithContext(aws.Context, *signer.CancelSigningProfileInput, ...request.Option) (*signer.CancelSigningProfileOutput, error)
	CancelSigningProfileRequest(*signer.CancelSigningProfileInput) (*request.Request, *signer.CancelSigningProfileOutput)

	DescribeSigningJob(*signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error)
	DescribeSigningJobWithContext(aws.Context, *signer.DescribeSigningJobInput, ...request.Option) (*signer.DescribeSigningJobOutput, error)
	DescribeSigningJobRequest(*signer.DescribeSigningJobInput) (*request.Request, *signer.DescribeSigningJobOutput)

	GetSigningPlatform(*signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error)
	GetSigningPlatformWithContext(aws.Context, *signer.GetSigningPlatformInput, ...request.Option) (*signer.GetSigningPlatformOutput, error)
	GetSigningPlatformRequest(*signer.GetSigningPlatformInput) (*request.Request, *signer.GetSigningPlatformOutput)

	GetSigningProfile(*signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error)
	GetSigningProfileWithContext(aws.Context, *signer.GetSigningProfileInput, ...request.Option) (*signer.GetSigningProfileOutput, error)
	GetSigningProfileRequest(*signer.GetSigningProfileInput) (*request.Request, *signer.GetSigningProfileOutput)

	ListSigningJobs(*signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error)
	ListSigningJobsWithContext(aws.Context, *signer.ListSigningJobsInput, ...request.Option) (*signer.ListSigningJobsOutput, error)
	ListSigningJobsRequest(*signer.ListSigningJobsInput) (*request.Request, *signer.ListSigningJobsOutput)

	ListSigningJobsPages(*signer.ListSigningJobsInput, func(*signer.ListSigningJobsOutput, bool) bool) error
	ListSigningJobsPagesWithContext(aws.Context, *signer.ListSigningJobsInput, func(*signer.ListSigningJobsOutput, bool) bool, ...request.Option) error

	ListSigningPlatforms(*signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error)
	ListSigningPlatformsWithContext(aws.Context, *signer.ListSigningPlatformsInput, ...request.Option) (*signer.ListSigningPlatformsOutput, error)
	ListSigningPlatformsRequest(*signer.ListSigningPlatformsInput) (*request.Request, *signer.ListSigningPlatformsOutput)

	ListSigningPlatformsPages(*signer.ListSigningPlatformsInput, func(*signer.ListSigningPlatformsOutput, bool) bool) error
	ListSigningPlatformsPagesWithContext(aws.Context, *signer.ListSigningPlatformsInput, func(*signer.ListSigningPlatformsOutput, bool) bool, ...request.Option) error

	ListSigningProfiles(*signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error)
	ListSigningProfilesWithContext(aws.Context, *signer.ListSigningProfilesInput, ...request.Option) (*signer.ListSigningProfilesOutput, error)
	ListSigningProfilesRequest(*signer.ListSigningProfilesInput) (*request.Request, *signer.ListSigningProfilesOutput)

	ListSigningProfilesPages(*signer.ListSigningProfilesInput, func(*signer.ListSigningProfilesOutput, bool) bool) error
	ListSigningProfilesPagesWithContext(aws.Context, *signer.ListSigningProfilesInput, func(*signer.ListSigningProfilesOutput, bool) bool, ...request.Option) error

	PutSigningProfile(*signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error)
	PutSigningProfileWithContext(aws.Context, *signer.PutSigningProfileInput, ...request.Option) (*signer.PutSigningProfileOutput, error)
	PutSigningProfileRequest(*signer.PutSigningProfileInput) (*request.Request, *signer.PutSigningProfileOutput)

	StartSigningJob(*signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error)
	StartSigningJobWithContext(aws.Context, *signer.StartSigningJobInput, ...request.Option) (*signer.StartSigningJobOutput, error)
	StartSigningJobRequest(*signer.StartSigningJobInput) (*request.Request, *signer.StartSigningJobOutput)

	WaitUntilSuccessfulSigningJob(*signer.DescribeSigningJobInput) error
	WaitUntilSuccessfulSigningJobWithContext(aws.Context, *signer.DescribeSigningJobInput, ...request.WaiterOption) error
}

var _ SignerAPI = (*signer.Signer)(nil)
