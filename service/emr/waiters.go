// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"time"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

// WaitUntilClusterRunning uses the Amazon EMR API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EMR) WaitUntilClusterRunning(input *DescribeClusterInput) error {
	return c.WaitUntilClusterRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterRunningWithContext is an extended version of WaitUntilClusterRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) WaitUntilClusterRunningWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterRunning",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "RUNNING",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "WAITING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "TERMINATING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "TERMINATED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "TERMINATED_WITH_ERRORS",
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

// WaitUntilClusterTerminated uses the Amazon EMR API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EMR) WaitUntilClusterTerminated(input *DescribeClusterInput) error {
	return c.WaitUntilClusterTerminatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterTerminatedWithContext is an extended version of WaitUntilClusterTerminated.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) WaitUntilClusterTerminatedWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterTerminated",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "TERMINATED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Cluster.Status.State",
				Expected: "TERMINATED_WITH_ERRORS",
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

// WaitUntilStepComplete uses the Amazon EMR API operation
// DescribeStep to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EMR) WaitUntilStepComplete(input *DescribeStepInput) error {
	return c.WaitUntilStepCompleteWithContext(aws.BackgroundContext(), input)
}

// WaitUntilStepCompleteWithContext is an extended version of WaitUntilStepComplete.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EMR) WaitUntilStepCompleteWithContext(ctx aws.Context, input *DescribeStepInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilStepComplete",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Step.Status.State",
				Expected: "COMPLETED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Step.Status.State",
				Expected: "FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Step.Status.State",
				Expected: "CANCELLED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeStepInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeStepRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
