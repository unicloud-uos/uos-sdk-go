// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package comprehendmedicaliface provides an interface to enable mocking the AWS Comprehend Medical service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package comprehendmedicaliface

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/comprehendmedical"
)

// ComprehendMedicalAPI provides an interface to enable mocking the
// comprehendmedical.ComprehendMedical service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Comprehend Medical.
//    func myFunc(svc comprehendmedicaliface.ComprehendMedicalAPI) bool {
//        // Make svc.DetectEntities request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := comprehendmedical.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComprehendMedicalClient struct {
//        comprehendmedicaliface.ComprehendMedicalAPI
//    }
//    func (m *mockComprehendMedicalClient) DetectEntities(input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComprehendMedicalClient{}
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
type ComprehendMedicalAPI interface {
	DetectEntities(*comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesWithContext(aws.Context, *comprehendmedical.DetectEntitiesInput, ...request.Option) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesRequest(*comprehendmedical.DetectEntitiesInput) (*request.Request, *comprehendmedical.DetectEntitiesOutput)

	DetectPHI(*comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIWithContext(aws.Context, *comprehendmedical.DetectPHIInput, ...request.Option) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIRequest(*comprehendmedical.DetectPHIInput) (*request.Request, *comprehendmedical.DetectPHIOutput)
}

var _ ComprehendMedicalAPI = (*comprehendmedical.ComprehendMedical)(nil)
