// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaiface provides an interface to enable mocking the Managed Streaming for Kafka service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafkaiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/kafka"
)

// KafkaAPI provides an interface to enable mocking the
// kafka.Kafka service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Managed Streaming for Kafka.
//    func myFunc(svc kafkaiface.KafkaAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kafka.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKafkaClient struct {
//        kafkaiface.KafkaAPI
//    }
//    func (m *mockKafkaClient) CreateCluster(input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKafkaClient{}
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
type KafkaAPI interface {
	CreateCluster(*kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *kafka.CreateClusterInput, ...request.Option) (*kafka.CreateClusterOutput, error)
	CreateClusterRequest(*kafka.CreateClusterInput) (*request.Request, *kafka.CreateClusterOutput)

	DeleteCluster(*kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *kafka.DeleteClusterInput, ...request.Option) (*kafka.DeleteClusterOutput, error)
	DeleteClusterRequest(*kafka.DeleteClusterInput) (*request.Request, *kafka.DeleteClusterOutput)

	DescribeCluster(*kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *kafka.DescribeClusterInput, ...request.Option) (*kafka.DescribeClusterOutput, error)
	DescribeClusterRequest(*kafka.DescribeClusterInput) (*request.Request, *kafka.DescribeClusterOutput)

	GetBootstrapBrokers(*kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersWithContext(aws.Context, *kafka.GetBootstrapBrokersInput, ...request.Option) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersRequest(*kafka.GetBootstrapBrokersInput) (*request.Request, *kafka.GetBootstrapBrokersOutput)

	ListClusters(*kafka.ListClustersInput) (*kafka.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *kafka.ListClustersInput, ...request.Option) (*kafka.ListClustersOutput, error)
	ListClustersRequest(*kafka.ListClustersInput) (*request.Request, *kafka.ListClustersOutput)

	ListNodes(*kafka.ListNodesInput) (*kafka.ListNodesOutput, error)
	ListNodesWithContext(aws.Context, *kafka.ListNodesInput, ...request.Option) (*kafka.ListNodesOutput, error)
	ListNodesRequest(*kafka.ListNodesInput) (*request.Request, *kafka.ListNodesOutput)
}

var _ KafkaAPI = (*kafka.Kafka)(nil)
