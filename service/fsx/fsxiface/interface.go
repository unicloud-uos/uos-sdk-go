// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fsxiface provides an interface to enable mocking the Amazon FSx service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fsxiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/fsx"
)

// FSxAPI provides an interface to enable mocking the
// fsx.FSx service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon FSx.
//    func myFunc(svc fsxiface.FSxAPI) bool {
//        // Make svc.CreateBackup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := fsx.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFSxClient struct {
//        fsxiface.FSxAPI
//    }
//    func (m *mockFSxClient) CreateBackup(input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFSxClient{}
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
type FSxAPI interface {
	CreateBackup(*fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error)
	CreateBackupWithContext(aws.Context, *fsx.CreateBackupInput, ...request.Option) (*fsx.CreateBackupOutput, error)
	CreateBackupRequest(*fsx.CreateBackupInput) (*request.Request, *fsx.CreateBackupOutput)

	CreateFileSystem(*fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error)
	CreateFileSystemWithContext(aws.Context, *fsx.CreateFileSystemInput, ...request.Option) (*fsx.CreateFileSystemOutput, error)
	CreateFileSystemRequest(*fsx.CreateFileSystemInput) (*request.Request, *fsx.CreateFileSystemOutput)

	CreateFileSystemFromBackup(*fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error)
	CreateFileSystemFromBackupWithContext(aws.Context, *fsx.CreateFileSystemFromBackupInput, ...request.Option) (*fsx.CreateFileSystemFromBackupOutput, error)
	CreateFileSystemFromBackupRequest(*fsx.CreateFileSystemFromBackupInput) (*request.Request, *fsx.CreateFileSystemFromBackupOutput)

	DeleteBackup(*fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error)
	DeleteBackupWithContext(aws.Context, *fsx.DeleteBackupInput, ...request.Option) (*fsx.DeleteBackupOutput, error)
	DeleteBackupRequest(*fsx.DeleteBackupInput) (*request.Request, *fsx.DeleteBackupOutput)

	DeleteFileSystem(*fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error)
	DeleteFileSystemWithContext(aws.Context, *fsx.DeleteFileSystemInput, ...request.Option) (*fsx.DeleteFileSystemOutput, error)
	DeleteFileSystemRequest(*fsx.DeleteFileSystemInput) (*request.Request, *fsx.DeleteFileSystemOutput)

	DescribeBackups(*fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error)
	DescribeBackupsWithContext(aws.Context, *fsx.DescribeBackupsInput, ...request.Option) (*fsx.DescribeBackupsOutput, error)
	DescribeBackupsRequest(*fsx.DescribeBackupsInput) (*request.Request, *fsx.DescribeBackupsOutput)

	DescribeBackupsPages(*fsx.DescribeBackupsInput, func(*fsx.DescribeBackupsOutput, bool) bool) error
	DescribeBackupsPagesWithContext(aws.Context, *fsx.DescribeBackupsInput, func(*fsx.DescribeBackupsOutput, bool) bool, ...request.Option) error

	DescribeFileSystems(*fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error)
	DescribeFileSystemsWithContext(aws.Context, *fsx.DescribeFileSystemsInput, ...request.Option) (*fsx.DescribeFileSystemsOutput, error)
	DescribeFileSystemsRequest(*fsx.DescribeFileSystemsInput) (*request.Request, *fsx.DescribeFileSystemsOutput)

	DescribeFileSystemsPages(*fsx.DescribeFileSystemsInput, func(*fsx.DescribeFileSystemsOutput, bool) bool) error
	DescribeFileSystemsPagesWithContext(aws.Context, *fsx.DescribeFileSystemsInput, func(*fsx.DescribeFileSystemsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *fsx.ListTagsForResourceInput, ...request.Option) (*fsx.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*fsx.ListTagsForResourceInput) (*request.Request, *fsx.ListTagsForResourceOutput)

	TagResource(*fsx.TagResourceInput) (*fsx.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *fsx.TagResourceInput, ...request.Option) (*fsx.TagResourceOutput, error)
	TagResourceRequest(*fsx.TagResourceInput) (*request.Request, *fsx.TagResourceOutput)

	UntagResource(*fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *fsx.UntagResourceInput, ...request.Option) (*fsx.UntagResourceOutput, error)
	UntagResourceRequest(*fsx.UntagResourceInput) (*request.Request, *fsx.UntagResourceOutput)

	UpdateFileSystem(*fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error)
	UpdateFileSystemWithContext(aws.Context, *fsx.UpdateFileSystemInput, ...request.Option) (*fsx.UpdateFileSystemOutput, error)
	UpdateFileSystemRequest(*fsx.UpdateFileSystemInput) (*request.Request, *fsx.UpdateFileSystemOutput)
}

var _ FSxAPI = (*fsx.FSx)(nil)
