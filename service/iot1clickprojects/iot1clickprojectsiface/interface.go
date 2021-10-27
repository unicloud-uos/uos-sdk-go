// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iot1clickprojectsiface provides an interface to enable mocking the AWS IoT 1-Click Projects Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iot1clickprojectsiface

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/iot1clickprojects"
)

// IoT1ClickProjectsAPI provides an interface to enable mocking the
// iot1clickprojects.IoT1ClickProjects service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT 1-Click Projects Service.
//    func myFunc(svc iot1clickprojectsiface.IoT1ClickProjectsAPI) bool {
//        // Make svc.AssociateDeviceWithPlacement request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iot1clickprojects.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoT1ClickProjectsClient struct {
//        iot1clickprojectsiface.IoT1ClickProjectsAPI
//    }
//    func (m *mockIoT1ClickProjectsClient) AssociateDeviceWithPlacement(input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoT1ClickProjectsClient{}
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
type IoT1ClickProjectsAPI interface {
	AssociateDeviceWithPlacement(*iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error)
	AssociateDeviceWithPlacementWithContext(aws.Context, *iot1clickprojects.AssociateDeviceWithPlacementInput, ...request.Option) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error)
	AssociateDeviceWithPlacementRequest(*iot1clickprojects.AssociateDeviceWithPlacementInput) (*request.Request, *iot1clickprojects.AssociateDeviceWithPlacementOutput)

	CreatePlacement(*iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error)
	CreatePlacementWithContext(aws.Context, *iot1clickprojects.CreatePlacementInput, ...request.Option) (*iot1clickprojects.CreatePlacementOutput, error)
	CreatePlacementRequest(*iot1clickprojects.CreatePlacementInput) (*request.Request, *iot1clickprojects.CreatePlacementOutput)

	CreateProject(*iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *iot1clickprojects.CreateProjectInput, ...request.Option) (*iot1clickprojects.CreateProjectOutput, error)
	CreateProjectRequest(*iot1clickprojects.CreateProjectInput) (*request.Request, *iot1clickprojects.CreateProjectOutput)

	DeletePlacement(*iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error)
	DeletePlacementWithContext(aws.Context, *iot1clickprojects.DeletePlacementInput, ...request.Option) (*iot1clickprojects.DeletePlacementOutput, error)
	DeletePlacementRequest(*iot1clickprojects.DeletePlacementInput) (*request.Request, *iot1clickprojects.DeletePlacementOutput)

	DeleteProject(*iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *iot1clickprojects.DeleteProjectInput, ...request.Option) (*iot1clickprojects.DeleteProjectOutput, error)
	DeleteProjectRequest(*iot1clickprojects.DeleteProjectInput) (*request.Request, *iot1clickprojects.DeleteProjectOutput)

	DescribePlacement(*iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error)
	DescribePlacementWithContext(aws.Context, *iot1clickprojects.DescribePlacementInput, ...request.Option) (*iot1clickprojects.DescribePlacementOutput, error)
	DescribePlacementRequest(*iot1clickprojects.DescribePlacementInput) (*request.Request, *iot1clickprojects.DescribePlacementOutput)

	DescribeProject(*iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error)
	DescribeProjectWithContext(aws.Context, *iot1clickprojects.DescribeProjectInput, ...request.Option) (*iot1clickprojects.DescribeProjectOutput, error)
	DescribeProjectRequest(*iot1clickprojects.DescribeProjectInput) (*request.Request, *iot1clickprojects.DescribeProjectOutput)

	DisassociateDeviceFromPlacement(*iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error)
	DisassociateDeviceFromPlacementWithContext(aws.Context, *iot1clickprojects.DisassociateDeviceFromPlacementInput, ...request.Option) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error)
	DisassociateDeviceFromPlacementRequest(*iot1clickprojects.DisassociateDeviceFromPlacementInput) (*request.Request, *iot1clickprojects.DisassociateDeviceFromPlacementOutput)

	GetDevicesInPlacement(*iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error)
	GetDevicesInPlacementWithContext(aws.Context, *iot1clickprojects.GetDevicesInPlacementInput, ...request.Option) (*iot1clickprojects.GetDevicesInPlacementOutput, error)
	GetDevicesInPlacementRequest(*iot1clickprojects.GetDevicesInPlacementInput) (*request.Request, *iot1clickprojects.GetDevicesInPlacementOutput)

	ListPlacements(*iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error)
	ListPlacementsWithContext(aws.Context, *iot1clickprojects.ListPlacementsInput, ...request.Option) (*iot1clickprojects.ListPlacementsOutput, error)
	ListPlacementsRequest(*iot1clickprojects.ListPlacementsInput) (*request.Request, *iot1clickprojects.ListPlacementsOutput)

	ListProjects(*iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *iot1clickprojects.ListProjectsInput, ...request.Option) (*iot1clickprojects.ListProjectsOutput, error)
	ListProjectsRequest(*iot1clickprojects.ListProjectsInput) (*request.Request, *iot1clickprojects.ListProjectsOutput)

	UpdatePlacement(*iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error)
	UpdatePlacementWithContext(aws.Context, *iot1clickprojects.UpdatePlacementInput, ...request.Option) (*iot1clickprojects.UpdatePlacementOutput, error)
	UpdatePlacementRequest(*iot1clickprojects.UpdatePlacementInput) (*request.Request, *iot1clickprojects.UpdatePlacementOutput)

	UpdateProject(*iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *iot1clickprojects.UpdateProjectInput, ...request.Option) (*iot1clickprojects.UpdateProjectOutput, error)
	UpdateProjectRequest(*iot1clickprojects.UpdateProjectInput) (*request.Request, *iot1clickprojects.UpdateProjectOutput)
}

var _ IoT1ClickProjectsAPI = (*iot1clickprojects.IoT1ClickProjects)(nil)
