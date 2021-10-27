// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"time"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

// WaitUntilClusterActive uses the Amazon EKS API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilClusterActive(input *DescribeClusterInput) error {
	return c.WaitUntilClusterActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterActiveWithContext is an extended version of WaitUntilClusterActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilClusterActiveWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterActive",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "DELETING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "FAILED",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeClusterInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeClusterRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilClusterDeleted uses the Amazon EKS API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilClusterDeleted(input *DescribeClusterInput) error {
	return c.WaitUntilClusterDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterDeletedWithContext is an extended version of WaitUntilClusterDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilClusterDeletedWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterDeleted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "ACTIVE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "CREATING",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeClusterInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeClusterRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
