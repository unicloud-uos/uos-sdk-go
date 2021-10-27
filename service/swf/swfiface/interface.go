// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package swfiface provides an interface to enable mocking the Amazon Simple Workflow Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package swfiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/swf"
)

// SWFAPI provides an interface to enable mocking the
// swf.SWF service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Workflow Service.
//    func myFunc(svc swfiface.SWFAPI) bool {
//        // Make svc.CountClosedWorkflowExecutions request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := swf.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSWFClient struct {
//        swfiface.SWFAPI
//    }
//    func (m *mockSWFClient) CountClosedWorkflowExecutions(input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSWFClient{}
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
type SWFAPI interface {
	CountClosedWorkflowExecutions(*swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
	CountClosedWorkflowExecutionsWithContext(aws.Context, *swf.CountClosedWorkflowExecutionsInput, ...request.Option) (*swf.WorkflowExecutionCount, error)
	CountClosedWorkflowExecutionsRequest(*swf.CountClosedWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionCount)

	CountOpenWorkflowExecutions(*swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
	CountOpenWorkflowExecutionsWithContext(aws.Context, *swf.CountOpenWorkflowExecutionsInput, ...request.Option) (*swf.WorkflowExecutionCount, error)
	CountOpenWorkflowExecutionsRequest(*swf.CountOpenWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionCount)

	CountPendingActivityTasks(*swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error)
	CountPendingActivityTasksWithContext(aws.Context, *swf.CountPendingActivityTasksInput, ...request.Option) (*swf.PendingTaskCount, error)
	CountPendingActivityTasksRequest(*swf.CountPendingActivityTasksInput) (*request.Request, *swf.PendingTaskCount)

	CountPendingDecisionTasks(*swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error)
	CountPendingDecisionTasksWithContext(aws.Context, *swf.CountPendingDecisionTasksInput, ...request.Option) (*swf.PendingTaskCount, error)
	CountPendingDecisionTasksRequest(*swf.CountPendingDecisionTasksInput) (*request.Request, *swf.PendingTaskCount)

	DeprecateActivityType(*swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error)
	DeprecateActivityTypeWithContext(aws.Context, *swf.DeprecateActivityTypeInput, ...request.Option) (*swf.DeprecateActivityTypeOutput, error)
	DeprecateActivityTypeRequest(*swf.DeprecateActivityTypeInput) (*request.Request, *swf.DeprecateActivityTypeOutput)

	DeprecateDomain(*swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error)
	DeprecateDomainWithContext(aws.Context, *swf.DeprecateDomainInput, ...request.Option) (*swf.DeprecateDomainOutput, error)
	DeprecateDomainRequest(*swf.DeprecateDomainInput) (*request.Request, *swf.DeprecateDomainOutput)

	DeprecateWorkflowType(*swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error)
	DeprecateWorkflowTypeWithContext(aws.Context, *swf.DeprecateWorkflowTypeInput, ...request.Option) (*swf.DeprecateWorkflowTypeOutput, error)
	DeprecateWorkflowTypeRequest(*swf.DeprecateWorkflowTypeInput) (*request.Request, *swf.DeprecateWorkflowTypeOutput)

	DescribeActivityType(*swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error)
	DescribeActivityTypeWithContext(aws.Context, *swf.DescribeActivityTypeInput, ...request.Option) (*swf.DescribeActivityTypeOutput, error)
	DescribeActivityTypeRequest(*swf.DescribeActivityTypeInput) (*request.Request, *swf.DescribeActivityTypeOutput)

	DescribeDomain(*swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error)
	DescribeDomainWithContext(aws.Context, *swf.DescribeDomainInput, ...request.Option) (*swf.DescribeDomainOutput, error)
	DescribeDomainRequest(*swf.DescribeDomainInput) (*request.Request, *swf.DescribeDomainOutput)

	DescribeWorkflowExecution(*swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error)
	DescribeWorkflowExecutionWithContext(aws.Context, *swf.DescribeWorkflowExecutionInput, ...request.Option) (*swf.DescribeWorkflowExecutionOutput, error)
	DescribeWorkflowExecutionRequest(*swf.DescribeWorkflowExecutionInput) (*request.Request, *swf.DescribeWorkflowExecutionOutput)

	DescribeWorkflowType(*swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error)
	DescribeWorkflowTypeWithContext(aws.Context, *swf.DescribeWorkflowTypeInput, ...request.Option) (*swf.DescribeWorkflowTypeOutput, error)
	DescribeWorkflowTypeRequest(*swf.DescribeWorkflowTypeInput) (*request.Request, *swf.DescribeWorkflowTypeOutput)

	GetWorkflowExecutionHistory(*swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error)
	GetWorkflowExecutionHistoryWithContext(aws.Context, *swf.GetWorkflowExecutionHistoryInput, ...request.Option) (*swf.GetWorkflowExecutionHistoryOutput, error)
	GetWorkflowExecutionHistoryRequest(*swf.GetWorkflowExecutionHistoryInput) (*request.Request, *swf.GetWorkflowExecutionHistoryOutput)

	GetWorkflowExecutionHistoryPages(*swf.GetWorkflowExecutionHistoryInput, func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool) error
	GetWorkflowExecutionHistoryPagesWithContext(aws.Context, *swf.GetWorkflowExecutionHistoryInput, func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool, ...request.Option) error

	ListActivityTypes(*swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error)
	ListActivityTypesWithContext(aws.Context, *swf.ListActivityTypesInput, ...request.Option) (*swf.ListActivityTypesOutput, error)
	ListActivityTypesRequest(*swf.ListActivityTypesInput) (*request.Request, *swf.ListActivityTypesOutput)

	ListActivityTypesPages(*swf.ListActivityTypesInput, func(*swf.ListActivityTypesOutput, bool) bool) error
	ListActivityTypesPagesWithContext(aws.Context, *swf.ListActivityTypesInput, func(*swf.ListActivityTypesOutput, bool) bool, ...request.Option) error

	ListClosedWorkflowExecutions(*swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
	ListClosedWorkflowExecutionsWithContext(aws.Context, *swf.ListClosedWorkflowExecutionsInput, ...request.Option) (*swf.WorkflowExecutionInfos, error)
	ListClosedWorkflowExecutionsRequest(*swf.ListClosedWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionInfos)

	ListClosedWorkflowExecutionsPages(*swf.ListClosedWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool) error
	ListClosedWorkflowExecutionsPagesWithContext(aws.Context, *swf.ListClosedWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool, ...request.Option) error

	ListDomains(*swf.ListDomainsInput) (*swf.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *swf.ListDomainsInput, ...request.Option) (*swf.ListDomainsOutput, error)
	ListDomainsRequest(*swf.ListDomainsInput) (*request.Request, *swf.ListDomainsOutput)

	ListDomainsPages(*swf.ListDomainsInput, func(*swf.ListDomainsOutput, bool) bool) error
	ListDomainsPagesWithContext(aws.Context, *swf.ListDomainsInput, func(*swf.ListDomainsOutput, bool) bool, ...request.Option) error

	ListOpenWorkflowExecutions(*swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
	ListOpenWorkflowExecutionsWithContext(aws.Context, *swf.ListOpenWorkflowExecutionsInput, ...request.Option) (*swf.WorkflowExecutionInfos, error)
	ListOpenWorkflowExecutionsRequest(*swf.ListOpenWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionInfos)

	ListOpenWorkflowExecutionsPages(*swf.ListOpenWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool) error
	ListOpenWorkflowExecutionsPagesWithContext(aws.Context, *swf.ListOpenWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool, ...request.Option) error

	ListWorkflowTypes(*swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error)
	ListWorkflowTypesWithContext(aws.Context, *swf.ListWorkflowTypesInput, ...request.Option) (*swf.ListWorkflowTypesOutput, error)
	ListWorkflowTypesRequest(*swf.ListWorkflowTypesInput) (*request.Request, *swf.ListWorkflowTypesOutput)

	ListWorkflowTypesPages(*swf.ListWorkflowTypesInput, func(*swf.ListWorkflowTypesOutput, bool) bool) error
	ListWorkflowTypesPagesWithContext(aws.Context, *swf.ListWorkflowTypesInput, func(*swf.ListWorkflowTypesOutput, bool) bool, ...request.Option) error

	PollForActivityTask(*swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error)
	PollForActivityTaskWithContext(aws.Context, *swf.PollForActivityTaskInput, ...request.Option) (*swf.PollForActivityTaskOutput, error)
	PollForActivityTaskRequest(*swf.PollForActivityTaskInput) (*request.Request, *swf.PollForActivityTaskOutput)

	PollForDecisionTask(*swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error)
	PollForDecisionTaskWithContext(aws.Context, *swf.PollForDecisionTaskInput, ...request.Option) (*swf.PollForDecisionTaskOutput, error)
	PollForDecisionTaskRequest(*swf.PollForDecisionTaskInput) (*request.Request, *swf.PollForDecisionTaskOutput)

	PollForDecisionTaskPages(*swf.PollForDecisionTaskInput, func(*swf.PollForDecisionTaskOutput, bool) bool) error
	PollForDecisionTaskPagesWithContext(aws.Context, *swf.PollForDecisionTaskInput, func(*swf.PollForDecisionTaskOutput, bool) bool, ...request.Option) error

	RecordActivityTaskHeartbeat(*swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error)
	RecordActivityTaskHeartbeatWithContext(aws.Context, *swf.RecordActivityTaskHeartbeatInput, ...request.Option) (*swf.RecordActivityTaskHeartbeatOutput, error)
	RecordActivityTaskHeartbeatRequest(*swf.RecordActivityTaskHeartbeatInput) (*request.Request, *swf.RecordActivityTaskHeartbeatOutput)

	RegisterActivityType(*swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error)
	RegisterActivityTypeWithContext(aws.Context, *swf.RegisterActivityTypeInput, ...request.Option) (*swf.RegisterActivityTypeOutput, error)
	RegisterActivityTypeRequest(*swf.RegisterActivityTypeInput) (*request.Request, *swf.RegisterActivityTypeOutput)

	RegisterDomain(*swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error)
	RegisterDomainWithContext(aws.Context, *swf.RegisterDomainInput, ...request.Option) (*swf.RegisterDomainOutput, error)
	RegisterDomainRequest(*swf.RegisterDomainInput) (*request.Request, *swf.RegisterDomainOutput)

	RegisterWorkflowType(*swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error)
	RegisterWorkflowTypeWithContext(aws.Context, *swf.RegisterWorkflowTypeInput, ...request.Option) (*swf.RegisterWorkflowTypeOutput, error)
	RegisterWorkflowTypeRequest(*swf.RegisterWorkflowTypeInput) (*request.Request, *swf.RegisterWorkflowTypeOutput)

	RequestCancelWorkflowExecution(*swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error)
	RequestCancelWorkflowExecutionWithContext(aws.Context, *swf.RequestCancelWorkflowExecutionInput, ...request.Option) (*swf.RequestCancelWorkflowExecutionOutput, error)
	RequestCancelWorkflowExecutionRequest(*swf.RequestCancelWorkflowExecutionInput) (*request.Request, *swf.RequestCancelWorkflowExecutionOutput)

	RespondActivityTaskCanceled(*swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error)
	RespondActivityTaskCanceledWithContext(aws.Context, *swf.RespondActivityTaskCanceledInput, ...request.Option) (*swf.RespondActivityTaskCanceledOutput, error)
	RespondActivityTaskCanceledRequest(*swf.RespondActivityTaskCanceledInput) (*request.Request, *swf.RespondActivityTaskCanceledOutput)

	RespondActivityTaskCompleted(*swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error)
	RespondActivityTaskCompletedWithContext(aws.Context, *swf.RespondActivityTaskCompletedInput, ...request.Option) (*swf.RespondActivityTaskCompletedOutput, error)
	RespondActivityTaskCompletedRequest(*swf.RespondActivityTaskCompletedInput) (*request.Request, *swf.RespondActivityTaskCompletedOutput)

	RespondActivityTaskFailed(*swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error)
	RespondActivityTaskFailedWithContext(aws.Context, *swf.RespondActivityTaskFailedInput, ...request.Option) (*swf.RespondActivityTaskFailedOutput, error)
	RespondActivityTaskFailedRequest(*swf.RespondActivityTaskFailedInput) (*request.Request, *swf.RespondActivityTaskFailedOutput)

	RespondDecisionTaskCompleted(*swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error)
	RespondDecisionTaskCompletedWithContext(aws.Context, *swf.RespondDecisionTaskCompletedInput, ...request.Option) (*swf.RespondDecisionTaskCompletedOutput, error)
	RespondDecisionTaskCompletedRequest(*swf.RespondDecisionTaskCompletedInput) (*request.Request, *swf.RespondDecisionTaskCompletedOutput)

	SignalWorkflowExecution(*swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error)
	SignalWorkflowExecutionWithContext(aws.Context, *swf.SignalWorkflowExecutionInput, ...request.Option) (*swf.SignalWorkflowExecutionOutput, error)
	SignalWorkflowExecutionRequest(*swf.SignalWorkflowExecutionInput) (*request.Request, *swf.SignalWorkflowExecutionOutput)

	StartWorkflowExecution(*swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error)
	StartWorkflowExecutionWithContext(aws.Context, *swf.StartWorkflowExecutionInput, ...request.Option) (*swf.StartWorkflowExecutionOutput, error)
	StartWorkflowExecutionRequest(*swf.StartWorkflowExecutionInput) (*request.Request, *swf.StartWorkflowExecutionOutput)

	TerminateWorkflowExecution(*swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error)
	TerminateWorkflowExecutionWithContext(aws.Context, *swf.TerminateWorkflowExecutionInput, ...request.Option) (*swf.TerminateWorkflowExecutionOutput, error)
	TerminateWorkflowExecutionRequest(*swf.TerminateWorkflowExecutionInput) (*request.Request, *swf.TerminateWorkflowExecutionOutput)
}

var _ SWFAPI = (*swf.SWF)(nil)
