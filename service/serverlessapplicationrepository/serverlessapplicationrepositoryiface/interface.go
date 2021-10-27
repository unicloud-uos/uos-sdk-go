// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package serverlessapplicationrepositoryiface provides an interface to enable mocking the AWSServerlessApplicationRepository service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package serverlessapplicationrepositoryiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/serverlessapplicationrepository"
)

// ServerlessApplicationRepositoryAPI provides an interface to enable mocking the
// serverlessapplicationrepository.ServerlessApplicationRepository service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWSServerlessApplicationRepository.
//    func myFunc(svc serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI) bool {
//        // Make svc.CreateApplication request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := serverlessapplicationrepository.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServerlessApplicationRepositoryClient struct {
//        serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI
//    }
//    func (m *mockServerlessApplicationRepositoryClient) CreateApplication(input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServerlessApplicationRepositoryClient{}
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
type ServerlessApplicationRepositoryAPI interface {
	CreateApplication(*serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *serverlessapplicationrepository.CreateApplicationRequest, ...request.Option) (*serverlessapplicationrepository.CreateApplicationOutput, error)
	CreateApplicationRequest(*serverlessapplicationrepository.CreateApplicationRequest) (*request.Request, *serverlessapplicationrepository.CreateApplicationOutput)

	CreateApplicationVersion(*serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
	CreateApplicationVersionWithContext(aws.Context, *serverlessapplicationrepository.CreateApplicationVersionRequest, ...request.Option) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
	CreateApplicationVersionRequest(*serverlessapplicationrepository.CreateApplicationVersionRequest) (*request.Request, *serverlessapplicationrepository.CreateApplicationVersionOutput)

	CreateCloudFormationChangeSet(*serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
	CreateCloudFormationChangeSetWithContext(aws.Context, *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest, ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
	CreateCloudFormationChangeSetRequest(*serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*request.Request, *serverlessapplicationrepository.CreateCloudFormationChangeSetOutput)

	CreateCloudFormationTemplate(*serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
	CreateCloudFormationTemplateWithContext(aws.Context, *serverlessapplicationrepository.CreateCloudFormationTemplateInput, ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
	CreateCloudFormationTemplateRequest(*serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*request.Request, *serverlessapplicationrepository.CreateCloudFormationTemplateOutput)

	DeleteApplication(*serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *serverlessapplicationrepository.DeleteApplicationInput, ...request.Option) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*serverlessapplicationrepository.DeleteApplicationInput) (*request.Request, *serverlessapplicationrepository.DeleteApplicationOutput)

	GetApplication(*serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error)
	GetApplicationWithContext(aws.Context, *serverlessapplicationrepository.GetApplicationInput, ...request.Option) (*serverlessapplicationrepository.GetApplicationOutput, error)
	GetApplicationRequest(*serverlessapplicationrepository.GetApplicationInput) (*request.Request, *serverlessapplicationrepository.GetApplicationOutput)

	GetApplicationPolicy(*serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
	GetApplicationPolicyWithContext(aws.Context, *serverlessapplicationrepository.GetApplicationPolicyInput, ...request.Option) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
	GetApplicationPolicyRequest(*serverlessapplicationrepository.GetApplicationPolicyInput) (*request.Request, *serverlessapplicationrepository.GetApplicationPolicyOutput)

	GetCloudFormationTemplate(*serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
	GetCloudFormationTemplateWithContext(aws.Context, *serverlessapplicationrepository.GetCloudFormationTemplateInput, ...request.Option) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
	GetCloudFormationTemplateRequest(*serverlessapplicationrepository.GetCloudFormationTemplateInput) (*request.Request, *serverlessapplicationrepository.GetCloudFormationTemplateOutput)

	ListApplicationDependencies(*serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
	ListApplicationDependenciesWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationDependenciesInput, ...request.Option) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
	ListApplicationDependenciesRequest(*serverlessapplicationrepository.ListApplicationDependenciesInput) (*request.Request, *serverlessapplicationrepository.ListApplicationDependenciesOutput)

	ListApplicationDependenciesPages(*serverlessapplicationrepository.ListApplicationDependenciesInput, func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool) error
	ListApplicationDependenciesPagesWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationDependenciesInput, func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool, ...request.Option) error

	ListApplicationVersions(*serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
	ListApplicationVersionsWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationVersionsInput, ...request.Option) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
	ListApplicationVersionsRequest(*serverlessapplicationrepository.ListApplicationVersionsInput) (*request.Request, *serverlessapplicationrepository.ListApplicationVersionsOutput)

	ListApplicationVersionsPages(*serverlessapplicationrepository.ListApplicationVersionsInput, func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool) error
	ListApplicationVersionsPagesWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationVersionsInput, func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool, ...request.Option) error

	ListApplications(*serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationsInput, ...request.Option) (*serverlessapplicationrepository.ListApplicationsOutput, error)
	ListApplicationsRequest(*serverlessapplicationrepository.ListApplicationsInput) (*request.Request, *serverlessapplicationrepository.ListApplicationsOutput)

	ListApplicationsPages(*serverlessapplicationrepository.ListApplicationsInput, func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *serverlessapplicationrepository.ListApplicationsInput, func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool, ...request.Option) error

	PutApplicationPolicy(*serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
	PutApplicationPolicyWithContext(aws.Context, *serverlessapplicationrepository.PutApplicationPolicyInput, ...request.Option) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
	PutApplicationPolicyRequest(*serverlessapplicationrepository.PutApplicationPolicyInput) (*request.Request, *serverlessapplicationrepository.PutApplicationPolicyOutput)

	UpdateApplication(*serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *serverlessapplicationrepository.UpdateApplicationRequest, ...request.Option) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*serverlessapplicationrepository.UpdateApplicationRequest) (*request.Request, *serverlessapplicationrepository.UpdateApplicationOutput)
}

var _ ServerlessApplicationRepositoryAPI = (*serverlessapplicationrepository.ServerlessApplicationRepository)(nil)
