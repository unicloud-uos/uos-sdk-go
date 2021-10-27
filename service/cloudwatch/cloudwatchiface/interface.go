// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatchiface provides an interface to enable mocking the Amazon CloudWatch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatchiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/cloudwatch"
)

// CloudWatchAPI provides an interface to enable mocking the
// cloudwatch.CloudWatch service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudWatch.
//    func myFunc(svc cloudwatchiface.CloudWatchAPI) bool {
//        // Make svc.DeleteAlarms request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudwatch.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudWatchClient struct {
//        cloudwatchiface.CloudWatchAPI
//    }
//    func (m *mockCloudWatchClient) DeleteAlarms(input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudWatchClient{}
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
type CloudWatchAPI interface {
	DeleteAlarms(*cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsWithContext(aws.Context, *cloudwatch.DeleteAlarmsInput, ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error)
	DeleteAlarmsRequest(*cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput)

	DeleteDashboards(*cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsWithContext(aws.Context, *cloudwatch.DeleteDashboardsInput, ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error)
	DeleteDashboardsRequest(*cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput)

	DescribeAlarmHistory(*cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryWithContext(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarmHistoryRequest(*cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput)

	DescribeAlarmHistoryPages(*cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error
	DescribeAlarmHistoryPagesWithContext(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, ...request.Option) error

	DescribeAlarms(*cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsRequest(*cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput)

	DescribeAlarmsPages(*cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error
	DescribeAlarmsPagesWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool, ...request.Option) error

	DescribeAlarmsForMetric(*cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricWithContext(aws.Context, *cloudwatch.DescribeAlarmsForMetricInput, ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAlarmsForMetricRequest(*cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput)

	DisableAlarmActions(*cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsWithContext(aws.Context, *cloudwatch.DisableAlarmActionsInput, ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error)
	DisableAlarmActionsRequest(*cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput)

	EnableAlarmActions(*cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsWithContext(aws.Context, *cloudwatch.EnableAlarmActionsInput, ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error)
	EnableAlarmActionsRequest(*cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput)

	GetDashboard(*cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardWithContext(aws.Context, *cloudwatch.GetDashboardInput, ...request.Option) (*cloudwatch.GetDashboardOutput, error)
	GetDashboardRequest(*cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput)

	GetMetricData(*cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataWithContext(aws.Context, *cloudwatch.GetMetricDataInput, ...request.Option) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricDataRequest(*cloudwatch.GetMetricDataInput) (*request.Request, *cloudwatch.GetMetricDataOutput)

	GetMetricStatistics(*cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsWithContext(aws.Context, *cloudwatch.GetMetricStatisticsInput, ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStatisticsRequest(*cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput)

	GetMetricWidgetImage(*cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageWithContext(aws.Context, *cloudwatch.GetMetricWidgetImageInput, ...request.Option) (*cloudwatch.GetMetricWidgetImageOutput, error)
	GetMetricWidgetImageRequest(*cloudwatch.GetMetricWidgetImageInput) (*request.Request, *cloudwatch.GetMetricWidgetImageOutput)

	ListDashboards(*cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsWithContext(aws.Context, *cloudwatch.ListDashboardsInput, ...request.Option) (*cloudwatch.ListDashboardsOutput, error)
	ListDashboardsRequest(*cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput)

	ListMetrics(*cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsWithContext(aws.Context, *cloudwatch.ListMetricsInput, ...request.Option) (*cloudwatch.ListMetricsOutput, error)
	ListMetricsRequest(*cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput)

	ListMetricsPages(*cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool) error
	ListMetricsPagesWithContext(aws.Context, *cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool, ...request.Option) error

	PutDashboard(*cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardWithContext(aws.Context, *cloudwatch.PutDashboardInput, ...request.Option) (*cloudwatch.PutDashboardOutput, error)
	PutDashboardRequest(*cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput)

	PutMetricAlarm(*cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmWithContext(aws.Context, *cloudwatch.PutMetricAlarmInput, ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error)
	PutMetricAlarmRequest(*cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput)

	PutMetricData(*cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataWithContext(aws.Context, *cloudwatch.PutMetricDataInput, ...request.Option) (*cloudwatch.PutMetricDataOutput, error)
	PutMetricDataRequest(*cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput)

	SetAlarmState(*cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateWithContext(aws.Context, *cloudwatch.SetAlarmStateInput, ...request.Option) (*cloudwatch.SetAlarmStateOutput, error)
	SetAlarmStateRequest(*cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput)

	WaitUntilAlarmExists(*cloudwatch.DescribeAlarmsInput) error
	WaitUntilAlarmExistsWithContext(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.WaiterOption) error
}

var _ CloudWatchAPI = (*cloudwatch.CloudWatch)(nil)
