// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package autoscalingiface provides an interface to enable mocking the Auto Scaling service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package autoscalingiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/autoscaling"
)

// AutoScalingAPI provides an interface to enable mocking the
// autoscaling.AutoScaling service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Auto Scaling.
//    func myFunc(svc autoscalingiface.AutoScalingAPI) bool {
//        // Make svc.AttachInstances request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := autoscaling.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAutoScalingClient struct {
//        autoscalingiface.AutoScalingAPI
//    }
//    func (m *mockAutoScalingClient) AttachInstances(input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAutoScalingClient{}
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
type AutoScalingAPI interface {
	AttachInstances(*autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error)
	AttachInstancesWithContext(aws.Context, *autoscaling.AttachInstancesInput, ...request.Option) (*autoscaling.AttachInstancesOutput, error)
	AttachInstancesRequest(*autoscaling.AttachInstancesInput) (*request.Request, *autoscaling.AttachInstancesOutput)

	AttachLoadBalancerTargetGroups(*autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error)
	AttachLoadBalancerTargetGroupsWithContext(aws.Context, *autoscaling.AttachLoadBalancerTargetGroupsInput, ...request.Option) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error)
	AttachLoadBalancerTargetGroupsRequest(*autoscaling.AttachLoadBalancerTargetGroupsInput) (*request.Request, *autoscaling.AttachLoadBalancerTargetGroupsOutput)

	AttachLoadBalancers(*autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error)
	AttachLoadBalancersWithContext(aws.Context, *autoscaling.AttachLoadBalancersInput, ...request.Option) (*autoscaling.AttachLoadBalancersOutput, error)
	AttachLoadBalancersRequest(*autoscaling.AttachLoadBalancersInput) (*request.Request, *autoscaling.AttachLoadBalancersOutput)

	BatchDeleteScheduledAction(*autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error)
	BatchDeleteScheduledActionWithContext(aws.Context, *autoscaling.BatchDeleteScheduledActionInput, ...request.Option) (*autoscaling.BatchDeleteScheduledActionOutput, error)
	BatchDeleteScheduledActionRequest(*autoscaling.BatchDeleteScheduledActionInput) (*request.Request, *autoscaling.BatchDeleteScheduledActionOutput)

	BatchPutScheduledUpdateGroupAction(*autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error)
	BatchPutScheduledUpdateGroupActionWithContext(aws.Context, *autoscaling.BatchPutScheduledUpdateGroupActionInput, ...request.Option) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error)
	BatchPutScheduledUpdateGroupActionRequest(*autoscaling.BatchPutScheduledUpdateGroupActionInput) (*request.Request, *autoscaling.BatchPutScheduledUpdateGroupActionOutput)

	CompleteLifecycleAction(*autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error)
	CompleteLifecycleActionWithContext(aws.Context, *autoscaling.CompleteLifecycleActionInput, ...request.Option) (*autoscaling.CompleteLifecycleActionOutput, error)
	CompleteLifecycleActionRequest(*autoscaling.CompleteLifecycleActionInput) (*request.Request, *autoscaling.CompleteLifecycleActionOutput)

	CreateAutoScalingGroup(*autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error)
	CreateAutoScalingGroupWithContext(aws.Context, *autoscaling.CreateAutoScalingGroupInput, ...request.Option) (*autoscaling.CreateAutoScalingGroupOutput, error)
	CreateAutoScalingGroupRequest(*autoscaling.CreateAutoScalingGroupInput) (*request.Request, *autoscaling.CreateAutoScalingGroupOutput)

	CreateLaunchConfiguration(*autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error)
	CreateLaunchConfigurationWithContext(aws.Context, *autoscaling.CreateLaunchConfigurationInput, ...request.Option) (*autoscaling.CreateLaunchConfigurationOutput, error)
	CreateLaunchConfigurationRequest(*autoscaling.CreateLaunchConfigurationInput) (*request.Request, *autoscaling.CreateLaunchConfigurationOutput)

	CreateOrUpdateTags(*autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error)
	CreateOrUpdateTagsWithContext(aws.Context, *autoscaling.CreateOrUpdateTagsInput, ...request.Option) (*autoscaling.CreateOrUpdateTagsOutput, error)
	CreateOrUpdateTagsRequest(*autoscaling.CreateOrUpdateTagsInput) (*request.Request, *autoscaling.CreateOrUpdateTagsOutput)

	DeleteAutoScalingGroup(*autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error)
	DeleteAutoScalingGroupWithContext(aws.Context, *autoscaling.DeleteAutoScalingGroupInput, ...request.Option) (*autoscaling.DeleteAutoScalingGroupOutput, error)
	DeleteAutoScalingGroupRequest(*autoscaling.DeleteAutoScalingGroupInput) (*request.Request, *autoscaling.DeleteAutoScalingGroupOutput)

	DeleteLaunchConfiguration(*autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error)
	DeleteLaunchConfigurationWithContext(aws.Context, *autoscaling.DeleteLaunchConfigurationInput, ...request.Option) (*autoscaling.DeleteLaunchConfigurationOutput, error)
	DeleteLaunchConfigurationRequest(*autoscaling.DeleteLaunchConfigurationInput) (*request.Request, *autoscaling.DeleteLaunchConfigurationOutput)

	DeleteLifecycleHook(*autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error)
	DeleteLifecycleHookWithContext(aws.Context, *autoscaling.DeleteLifecycleHookInput, ...request.Option) (*autoscaling.DeleteLifecycleHookOutput, error)
	DeleteLifecycleHookRequest(*autoscaling.DeleteLifecycleHookInput) (*request.Request, *autoscaling.DeleteLifecycleHookOutput)

	DeleteNotificationConfiguration(*autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error)
	DeleteNotificationConfigurationWithContext(aws.Context, *autoscaling.DeleteNotificationConfigurationInput, ...request.Option) (*autoscaling.DeleteNotificationConfigurationOutput, error)
	DeleteNotificationConfigurationRequest(*autoscaling.DeleteNotificationConfigurationInput) (*request.Request, *autoscaling.DeleteNotificationConfigurationOutput)

	DeletePolicy(*autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *autoscaling.DeletePolicyInput, ...request.Option) (*autoscaling.DeletePolicyOutput, error)
	DeletePolicyRequest(*autoscaling.DeletePolicyInput) (*request.Request, *autoscaling.DeletePolicyOutput)

	DeleteScheduledAction(*autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionWithContext(aws.Context, *autoscaling.DeleteScheduledActionInput, ...request.Option) (*autoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionRequest(*autoscaling.DeleteScheduledActionInput) (*request.Request, *autoscaling.DeleteScheduledActionOutput)

	DeleteTags(*autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *autoscaling.DeleteTagsInput, ...request.Option) (*autoscaling.DeleteTagsOutput, error)
	DeleteTagsRequest(*autoscaling.DeleteTagsInput) (*request.Request, *autoscaling.DeleteTagsOutput)

	DescribeAccountLimits(*autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *autoscaling.DescribeAccountLimitsInput, ...request.Option) (*autoscaling.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*autoscaling.DescribeAccountLimitsInput) (*request.Request, *autoscaling.DescribeAccountLimitsOutput)

	DescribeAdjustmentTypes(*autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error)
	DescribeAdjustmentTypesWithContext(aws.Context, *autoscaling.DescribeAdjustmentTypesInput, ...request.Option) (*autoscaling.DescribeAdjustmentTypesOutput, error)
	DescribeAdjustmentTypesRequest(*autoscaling.DescribeAdjustmentTypesInput) (*request.Request, *autoscaling.DescribeAdjustmentTypesOutput)

	DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribeAutoScalingGroupsWithContext(aws.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...request.Option) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribeAutoScalingGroupsRequest(*autoscaling.DescribeAutoScalingGroupsInput) (*request.Request, *autoscaling.DescribeAutoScalingGroupsOutput)

	DescribeAutoScalingGroupsPages(*autoscaling.DescribeAutoScalingGroupsInput, func(*autoscaling.DescribeAutoScalingGroupsOutput, bool) bool) error
	DescribeAutoScalingGroupsPagesWithContext(aws.Context, *autoscaling.DescribeAutoScalingGroupsInput, func(*autoscaling.DescribeAutoScalingGroupsOutput, bool) bool, ...request.Option) error

	DescribeAutoScalingInstances(*autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
	DescribeAutoScalingInstancesWithContext(aws.Context, *autoscaling.DescribeAutoScalingInstancesInput, ...request.Option) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
	DescribeAutoScalingInstancesRequest(*autoscaling.DescribeAutoScalingInstancesInput) (*request.Request, *autoscaling.DescribeAutoScalingInstancesOutput)

	DescribeAutoScalingInstancesPages(*autoscaling.DescribeAutoScalingInstancesInput, func(*autoscaling.DescribeAutoScalingInstancesOutput, bool) bool) error
	DescribeAutoScalingInstancesPagesWithContext(aws.Context, *autoscaling.DescribeAutoScalingInstancesInput, func(*autoscaling.DescribeAutoScalingInstancesOutput, bool) bool, ...request.Option) error

	DescribeAutoScalingNotificationTypes(*autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
	DescribeAutoScalingNotificationTypesWithContext(aws.Context, *autoscaling.DescribeAutoScalingNotificationTypesInput, ...request.Option) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
	DescribeAutoScalingNotificationTypesRequest(*autoscaling.DescribeAutoScalingNotificationTypesInput) (*request.Request, *autoscaling.DescribeAutoScalingNotificationTypesOutput)

	DescribeLaunchConfigurations(*autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeLaunchConfigurationsWithContext(aws.Context, *autoscaling.DescribeLaunchConfigurationsInput, ...request.Option) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeLaunchConfigurationsRequest(*autoscaling.DescribeLaunchConfigurationsInput) (*request.Request, *autoscaling.DescribeLaunchConfigurationsOutput)

	DescribeLaunchConfigurationsPages(*autoscaling.DescribeLaunchConfigurationsInput, func(*autoscaling.DescribeLaunchConfigurationsOutput, bool) bool) error
	DescribeLaunchConfigurationsPagesWithContext(aws.Context, *autoscaling.DescribeLaunchConfigurationsInput, func(*autoscaling.DescribeLaunchConfigurationsOutput, bool) bool, ...request.Option) error

	DescribeLifecycleHookTypes(*autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
	DescribeLifecycleHookTypesWithContext(aws.Context, *autoscaling.DescribeLifecycleHookTypesInput, ...request.Option) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
	DescribeLifecycleHookTypesRequest(*autoscaling.DescribeLifecycleHookTypesInput) (*request.Request, *autoscaling.DescribeLifecycleHookTypesOutput)

	DescribeLifecycleHooks(*autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeLifecycleHooksWithContext(aws.Context, *autoscaling.DescribeLifecycleHooksInput, ...request.Option) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeLifecycleHooksRequest(*autoscaling.DescribeLifecycleHooksInput) (*request.Request, *autoscaling.DescribeLifecycleHooksOutput)

	DescribeLoadBalancerTargetGroups(*autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLoadBalancerTargetGroupsWithContext(aws.Context, *autoscaling.DescribeLoadBalancerTargetGroupsInput, ...request.Option) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLoadBalancerTargetGroupsRequest(*autoscaling.DescribeLoadBalancerTargetGroupsInput) (*request.Request, *autoscaling.DescribeLoadBalancerTargetGroupsOutput)

	DescribeLoadBalancers(*autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersWithContext(aws.Context, *autoscaling.DescribeLoadBalancersInput, ...request.Option) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersRequest(*autoscaling.DescribeLoadBalancersInput) (*request.Request, *autoscaling.DescribeLoadBalancersOutput)

	DescribeMetricCollectionTypes(*autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
	DescribeMetricCollectionTypesWithContext(aws.Context, *autoscaling.DescribeMetricCollectionTypesInput, ...request.Option) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
	DescribeMetricCollectionTypesRequest(*autoscaling.DescribeMetricCollectionTypesInput) (*request.Request, *autoscaling.DescribeMetricCollectionTypesOutput)

	DescribeNotificationConfigurations(*autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribeNotificationConfigurationsWithContext(aws.Context, *autoscaling.DescribeNotificationConfigurationsInput, ...request.Option) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribeNotificationConfigurationsRequest(*autoscaling.DescribeNotificationConfigurationsInput) (*request.Request, *autoscaling.DescribeNotificationConfigurationsOutput)

	DescribeNotificationConfigurationsPages(*autoscaling.DescribeNotificationConfigurationsInput, func(*autoscaling.DescribeNotificationConfigurationsOutput, bool) bool) error
	DescribeNotificationConfigurationsPagesWithContext(aws.Context, *autoscaling.DescribeNotificationConfigurationsInput, func(*autoscaling.DescribeNotificationConfigurationsOutput, bool) bool, ...request.Option) error

	DescribePolicies(*autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error)
	DescribePoliciesWithContext(aws.Context, *autoscaling.DescribePoliciesInput, ...request.Option) (*autoscaling.DescribePoliciesOutput, error)
	DescribePoliciesRequest(*autoscaling.DescribePoliciesInput) (*request.Request, *autoscaling.DescribePoliciesOutput)

	DescribePoliciesPages(*autoscaling.DescribePoliciesInput, func(*autoscaling.DescribePoliciesOutput, bool) bool) error
	DescribePoliciesPagesWithContext(aws.Context, *autoscaling.DescribePoliciesInput, func(*autoscaling.DescribePoliciesOutput, bool) bool, ...request.Option) error

	DescribeScalingActivities(*autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesWithContext(aws.Context, *autoscaling.DescribeScalingActivitiesInput, ...request.Option) (*autoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesRequest(*autoscaling.DescribeScalingActivitiesInput) (*request.Request, *autoscaling.DescribeScalingActivitiesOutput)

	DescribeScalingActivitiesPages(*autoscaling.DescribeScalingActivitiesInput, func(*autoscaling.DescribeScalingActivitiesOutput, bool) bool) error
	DescribeScalingActivitiesPagesWithContext(aws.Context, *autoscaling.DescribeScalingActivitiesInput, func(*autoscaling.DescribeScalingActivitiesOutput, bool) bool, ...request.Option) error

	DescribeScalingProcessTypes(*autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error)
	DescribeScalingProcessTypesWithContext(aws.Context, *autoscaling.DescribeScalingProcessTypesInput, ...request.Option) (*autoscaling.DescribeScalingProcessTypesOutput, error)
	DescribeScalingProcessTypesRequest(*autoscaling.DescribeScalingProcessTypesInput) (*request.Request, *autoscaling.DescribeScalingProcessTypesOutput)

	DescribeScheduledActions(*autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsWithContext(aws.Context, *autoscaling.DescribeScheduledActionsInput, ...request.Option) (*autoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsRequest(*autoscaling.DescribeScheduledActionsInput) (*request.Request, *autoscaling.DescribeScheduledActionsOutput)

	DescribeScheduledActionsPages(*autoscaling.DescribeScheduledActionsInput, func(*autoscaling.DescribeScheduledActionsOutput, bool) bool) error
	DescribeScheduledActionsPagesWithContext(aws.Context, *autoscaling.DescribeScheduledActionsInput, func(*autoscaling.DescribeScheduledActionsOutput, bool) bool, ...request.Option) error

	DescribeTags(*autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *autoscaling.DescribeTagsInput, ...request.Option) (*autoscaling.DescribeTagsOutput, error)
	DescribeTagsRequest(*autoscaling.DescribeTagsInput) (*request.Request, *autoscaling.DescribeTagsOutput)

	DescribeTagsPages(*autoscaling.DescribeTagsInput, func(*autoscaling.DescribeTagsOutput, bool) bool) error
	DescribeTagsPagesWithContext(aws.Context, *autoscaling.DescribeTagsInput, func(*autoscaling.DescribeTagsOutput, bool) bool, ...request.Option) error

	DescribeTerminationPolicyTypes(*autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
	DescribeTerminationPolicyTypesWithContext(aws.Context, *autoscaling.DescribeTerminationPolicyTypesInput, ...request.Option) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
	DescribeTerminationPolicyTypesRequest(*autoscaling.DescribeTerminationPolicyTypesInput) (*request.Request, *autoscaling.DescribeTerminationPolicyTypesOutput)

	DetachInstances(*autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error)
	DetachInstancesWithContext(aws.Context, *autoscaling.DetachInstancesInput, ...request.Option) (*autoscaling.DetachInstancesOutput, error)
	DetachInstancesRequest(*autoscaling.DetachInstancesInput) (*request.Request, *autoscaling.DetachInstancesOutput)

	DetachLoadBalancerTargetGroups(*autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error)
	DetachLoadBalancerTargetGroupsWithContext(aws.Context, *autoscaling.DetachLoadBalancerTargetGroupsInput, ...request.Option) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error)
	DetachLoadBalancerTargetGroupsRequest(*autoscaling.DetachLoadBalancerTargetGroupsInput) (*request.Request, *autoscaling.DetachLoadBalancerTargetGroupsOutput)

	DetachLoadBalancers(*autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error)
	DetachLoadBalancersWithContext(aws.Context, *autoscaling.DetachLoadBalancersInput, ...request.Option) (*autoscaling.DetachLoadBalancersOutput, error)
	DetachLoadBalancersRequest(*autoscaling.DetachLoadBalancersInput) (*request.Request, *autoscaling.DetachLoadBalancersOutput)

	DisableMetricsCollection(*autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error)
	DisableMetricsCollectionWithContext(aws.Context, *autoscaling.DisableMetricsCollectionInput, ...request.Option) (*autoscaling.DisableMetricsCollectionOutput, error)
	DisableMetricsCollectionRequest(*autoscaling.DisableMetricsCollectionInput) (*request.Request, *autoscaling.DisableMetricsCollectionOutput)

	EnableMetricsCollection(*autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error)
	EnableMetricsCollectionWithContext(aws.Context, *autoscaling.EnableMetricsCollectionInput, ...request.Option) (*autoscaling.EnableMetricsCollectionOutput, error)
	EnableMetricsCollectionRequest(*autoscaling.EnableMetricsCollectionInput) (*request.Request, *autoscaling.EnableMetricsCollectionOutput)

	EnterStandby(*autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error)
	EnterStandbyWithContext(aws.Context, *autoscaling.EnterStandbyInput, ...request.Option) (*autoscaling.EnterStandbyOutput, error)
	EnterStandbyRequest(*autoscaling.EnterStandbyInput) (*request.Request, *autoscaling.EnterStandbyOutput)

	ExecutePolicy(*autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error)
	ExecutePolicyWithContext(aws.Context, *autoscaling.ExecutePolicyInput, ...request.Option) (*autoscaling.ExecutePolicyOutput, error)
	ExecutePolicyRequest(*autoscaling.ExecutePolicyInput) (*request.Request, *autoscaling.ExecutePolicyOutput)

	ExitStandby(*autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error)
	ExitStandbyWithContext(aws.Context, *autoscaling.ExitStandbyInput, ...request.Option) (*autoscaling.ExitStandbyOutput, error)
	ExitStandbyRequest(*autoscaling.ExitStandbyInput) (*request.Request, *autoscaling.ExitStandbyOutput)

	PutLifecycleHook(*autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error)
	PutLifecycleHookWithContext(aws.Context, *autoscaling.PutLifecycleHookInput, ...request.Option) (*autoscaling.PutLifecycleHookOutput, error)
	PutLifecycleHookRequest(*autoscaling.PutLifecycleHookInput) (*request.Request, *autoscaling.PutLifecycleHookOutput)

	PutNotificationConfiguration(*autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error)
	PutNotificationConfigurationWithContext(aws.Context, *autoscaling.PutNotificationConfigurationInput, ...request.Option) (*autoscaling.PutNotificationConfigurationOutput, error)
	PutNotificationConfigurationRequest(*autoscaling.PutNotificationConfigurationInput) (*request.Request, *autoscaling.PutNotificationConfigurationOutput)

	PutScalingPolicy(*autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyWithContext(aws.Context, *autoscaling.PutScalingPolicyInput, ...request.Option) (*autoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyRequest(*autoscaling.PutScalingPolicyInput) (*request.Request, *autoscaling.PutScalingPolicyOutput)

	PutScheduledUpdateGroupAction(*autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)
	PutScheduledUpdateGroupActionWithContext(aws.Context, *autoscaling.PutScheduledUpdateGroupActionInput, ...request.Option) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)
	PutScheduledUpdateGroupActionRequest(*autoscaling.PutScheduledUpdateGroupActionInput) (*request.Request, *autoscaling.PutScheduledUpdateGroupActionOutput)

	RecordLifecycleActionHeartbeat(*autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)
	RecordLifecycleActionHeartbeatWithContext(aws.Context, *autoscaling.RecordLifecycleActionHeartbeatInput, ...request.Option) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)
	RecordLifecycleActionHeartbeatRequest(*autoscaling.RecordLifecycleActionHeartbeatInput) (*request.Request, *autoscaling.RecordLifecycleActionHeartbeatOutput)

	ResumeProcesses(*autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error)
	ResumeProcessesWithContext(aws.Context, *autoscaling.ScalingProcessQuery, ...request.Option) (*autoscaling.ResumeProcessesOutput, error)
	ResumeProcessesRequest(*autoscaling.ScalingProcessQuery) (*request.Request, *autoscaling.ResumeProcessesOutput)

	SetDesiredCapacity(*autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error)
	SetDesiredCapacityWithContext(aws.Context, *autoscaling.SetDesiredCapacityInput, ...request.Option) (*autoscaling.SetDesiredCapacityOutput, error)
	SetDesiredCapacityRequest(*autoscaling.SetDesiredCapacityInput) (*request.Request, *autoscaling.SetDesiredCapacityOutput)

	SetInstanceHealth(*autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error)
	SetInstanceHealthWithContext(aws.Context, *autoscaling.SetInstanceHealthInput, ...request.Option) (*autoscaling.SetInstanceHealthOutput, error)
	SetInstanceHealthRequest(*autoscaling.SetInstanceHealthInput) (*request.Request, *autoscaling.SetInstanceHealthOutput)

	SetInstanceProtection(*autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error)
	SetInstanceProtectionWithContext(aws.Context, *autoscaling.SetInstanceProtectionInput, ...request.Option) (*autoscaling.SetInstanceProtectionOutput, error)
	SetInstanceProtectionRequest(*autoscaling.SetInstanceProtectionInput) (*request.Request, *autoscaling.SetInstanceProtectionOutput)

	SuspendProcesses(*autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error)
	SuspendProcessesWithContext(aws.Context, *autoscaling.ScalingProcessQuery, ...request.Option) (*autoscaling.SuspendProcessesOutput, error)
	SuspendProcessesRequest(*autoscaling.ScalingProcessQuery) (*request.Request, *autoscaling.SuspendProcessesOutput)

	TerminateInstanceInAutoScalingGroup(*autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
	TerminateInstanceInAutoScalingGroupWithContext(aws.Context, *autoscaling.TerminateInstanceInAutoScalingGroupInput, ...request.Option) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
	TerminateInstanceInAutoScalingGroupRequest(*autoscaling.TerminateInstanceInAutoScalingGroupInput) (*request.Request, *autoscaling.TerminateInstanceInAutoScalingGroupOutput)

	UpdateAutoScalingGroup(*autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error)
	UpdateAutoScalingGroupWithContext(aws.Context, *autoscaling.UpdateAutoScalingGroupInput, ...request.Option) (*autoscaling.UpdateAutoScalingGroupOutput, error)
	UpdateAutoScalingGroupRequest(*autoscaling.UpdateAutoScalingGroupInput) (*request.Request, *autoscaling.UpdateAutoScalingGroupOutput)

	WaitUntilGroupExists(*autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupExistsWithContext(aws.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...request.WaiterOption) error

	WaitUntilGroupInService(*autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupInServiceWithContext(aws.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...request.WaiterOption) error

	WaitUntilGroupNotExists(*autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupNotExistsWithContext(aws.Context, *autoscaling.DescribeAutoScalingGroupsInput, ...request.WaiterOption) error
}

var _ AutoScalingAPI = (*autoscaling.AutoScaling)(nil)
