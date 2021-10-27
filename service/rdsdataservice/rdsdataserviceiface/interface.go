// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsdataserviceiface provides an interface to enable mocking the AWS RDS DataService service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsdataserviceiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/rdsdataservice"
)

// RDSDataServiceAPI provides an interface to enable mocking the
// rdsdataservice.RDSDataService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS RDS DataService.
//    func myFunc(svc rdsdataserviceiface.RDSDataServiceAPI) bool {
//        // Make svc.ExecuteSql request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rdsdataservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRDSDataServiceClient struct {
//        rdsdataserviceiface.RDSDataServiceAPI
//    }
//    func (m *mockRDSDataServiceClient) ExecuteSql(input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRDSDataServiceClient{}
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
type RDSDataServiceAPI interface {
	ExecuteSql(*rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error)
	ExecuteSqlWithContext(aws.Context, *rdsdataservice.ExecuteSqlInput, ...request.Option) (*rdsdataservice.ExecuteSqlOutput, error)
	ExecuteSqlRequest(*rdsdataservice.ExecuteSqlInput) (*request.Request, *rdsdataservice.ExecuteSqlOutput)
}

var _ RDSDataServiceAPI = (*rdsdataservice.RDSDataService)(nil)
