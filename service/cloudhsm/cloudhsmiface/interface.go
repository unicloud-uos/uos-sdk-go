// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudhsmiface provides an interface to enable mocking the Amazon CloudHSM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudhsmiface

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/cloudhsm"
)

// CloudHSMAPI provides an interface to enable mocking the
// cloudhsm.CloudHSM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudHSM.
//    func myFunc(svc cloudhsmiface.CloudHSMAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudhsm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudHSMClient struct {
//        cloudhsmiface.CloudHSMAPI
//    }
//    func (m *mockCloudHSMClient) AddTagsToResource(input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudHSMClient{}
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
type CloudHSMAPI interface {
	AddTagsToResource(*cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(aws.Context, *cloudhsm.AddTagsToResourceInput, ...request.Option) (*cloudhsm.AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*cloudhsm.AddTagsToResourceInput) (*request.Request, *cloudhsm.AddTagsToResourceOutput)

	CreateHapg(*cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)
	CreateHapgWithContext(aws.Context, *cloudhsm.CreateHapgInput, ...request.Option) (*cloudhsm.CreateHapgOutput, error)
	CreateHapgRequest(*cloudhsm.CreateHapgInput) (*request.Request, *cloudhsm.CreateHapgOutput)

	CreateHsm(*cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)
	CreateHsmWithContext(aws.Context, *cloudhsm.CreateHsmInput, ...request.Option) (*cloudhsm.CreateHsmOutput, error)
	CreateHsmRequest(*cloudhsm.CreateHsmInput) (*request.Request, *cloudhsm.CreateHsmOutput)

	CreateLunaClient(*cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)
	CreateLunaClientWithContext(aws.Context, *cloudhsm.CreateLunaClientInput, ...request.Option) (*cloudhsm.CreateLunaClientOutput, error)
	CreateLunaClientRequest(*cloudhsm.CreateLunaClientInput) (*request.Request, *cloudhsm.CreateLunaClientOutput)

	DeleteHapg(*cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)
	DeleteHapgWithContext(aws.Context, *cloudhsm.DeleteHapgInput, ...request.Option) (*cloudhsm.DeleteHapgOutput, error)
	DeleteHapgRequest(*cloudhsm.DeleteHapgInput) (*request.Request, *cloudhsm.DeleteHapgOutput)

	DeleteHsm(*cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)
	DeleteHsmWithContext(aws.Context, *cloudhsm.DeleteHsmInput, ...request.Option) (*cloudhsm.DeleteHsmOutput, error)
	DeleteHsmRequest(*cloudhsm.DeleteHsmInput) (*request.Request, *cloudhsm.DeleteHsmOutput)

	DeleteLunaClient(*cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)
	DeleteLunaClientWithContext(aws.Context, *cloudhsm.DeleteLunaClientInput, ...request.Option) (*cloudhsm.DeleteLunaClientOutput, error)
	DeleteLunaClientRequest(*cloudhsm.DeleteLunaClientInput) (*request.Request, *cloudhsm.DeleteLunaClientOutput)

	DescribeHapg(*cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)
	DescribeHapgWithContext(aws.Context, *cloudhsm.DescribeHapgInput, ...request.Option) (*cloudhsm.DescribeHapgOutput, error)
	DescribeHapgRequest(*cloudhsm.DescribeHapgInput) (*request.Request, *cloudhsm.DescribeHapgOutput)

	DescribeHsm(*cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)
	DescribeHsmWithContext(aws.Context, *cloudhsm.DescribeHsmInput, ...request.Option) (*cloudhsm.DescribeHsmOutput, error)
	DescribeHsmRequest(*cloudhsm.DescribeHsmInput) (*request.Request, *cloudhsm.DescribeHsmOutput)

	DescribeLunaClient(*cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)
	DescribeLunaClientWithContext(aws.Context, *cloudhsm.DescribeLunaClientInput, ...request.Option) (*cloudhsm.DescribeLunaClientOutput, error)
	DescribeLunaClientRequest(*cloudhsm.DescribeLunaClientInput) (*request.Request, *cloudhsm.DescribeLunaClientOutput)

	GetConfig(*cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)
	GetConfigWithContext(aws.Context, *cloudhsm.GetConfigInput, ...request.Option) (*cloudhsm.GetConfigOutput, error)
	GetConfigRequest(*cloudhsm.GetConfigInput) (*request.Request, *cloudhsm.GetConfigOutput)

	ListAvailableZones(*cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)
	ListAvailableZonesWithContext(aws.Context, *cloudhsm.ListAvailableZonesInput, ...request.Option) (*cloudhsm.ListAvailableZonesOutput, error)
	ListAvailableZonesRequest(*cloudhsm.ListAvailableZonesInput) (*request.Request, *cloudhsm.ListAvailableZonesOutput)

	ListHapgs(*cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)
	ListHapgsWithContext(aws.Context, *cloudhsm.ListHapgsInput, ...request.Option) (*cloudhsm.ListHapgsOutput, error)
	ListHapgsRequest(*cloudhsm.ListHapgsInput) (*request.Request, *cloudhsm.ListHapgsOutput)

	ListHsms(*cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)
	ListHsmsWithContext(aws.Context, *cloudhsm.ListHsmsInput, ...request.Option) (*cloudhsm.ListHsmsOutput, error)
	ListHsmsRequest(*cloudhsm.ListHsmsInput) (*request.Request, *cloudhsm.ListHsmsOutput)

	ListLunaClients(*cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)
	ListLunaClientsWithContext(aws.Context, *cloudhsm.ListLunaClientsInput, ...request.Option) (*cloudhsm.ListLunaClientsOutput, error)
	ListLunaClientsRequest(*cloudhsm.ListLunaClientsInput) (*request.Request, *cloudhsm.ListLunaClientsOutput)

	ListTagsForResource(*cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cloudhsm.ListTagsForResourceInput, ...request.Option) (*cloudhsm.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cloudhsm.ListTagsForResourceInput) (*request.Request, *cloudhsm.ListTagsForResourceOutput)

	ModifyHapg(*cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)
	ModifyHapgWithContext(aws.Context, *cloudhsm.ModifyHapgInput, ...request.Option) (*cloudhsm.ModifyHapgOutput, error)
	ModifyHapgRequest(*cloudhsm.ModifyHapgInput) (*request.Request, *cloudhsm.ModifyHapgOutput)

	ModifyHsm(*cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)
	ModifyHsmWithContext(aws.Context, *cloudhsm.ModifyHsmInput, ...request.Option) (*cloudhsm.ModifyHsmOutput, error)
	ModifyHsmRequest(*cloudhsm.ModifyHsmInput) (*request.Request, *cloudhsm.ModifyHsmOutput)

	ModifyLunaClient(*cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
	ModifyLunaClientWithContext(aws.Context, *cloudhsm.ModifyLunaClientInput, ...request.Option) (*cloudhsm.ModifyLunaClientOutput, error)
	ModifyLunaClientRequest(*cloudhsm.ModifyLunaClientInput) (*request.Request, *cloudhsm.ModifyLunaClientOutput)

	RemoveTagsFromResource(*cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(aws.Context, *cloudhsm.RemoveTagsFromResourceInput, ...request.Option) (*cloudhsm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*cloudhsm.RemoveTagsFromResourceInput) (*request.Request, *cloudhsm.RemoveTagsFromResourceOutput)
}

var _ CloudHSMAPI = (*cloudhsm.CloudHSM)(nil)
