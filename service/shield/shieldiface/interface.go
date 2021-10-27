// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package shieldiface provides an interface to enable mocking the AWS Shield service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package shieldiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/shield"
)

// ShieldAPI provides an interface to enable mocking the
// shield.Shield service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Shield.
//    func myFunc(svc shieldiface.ShieldAPI) bool {
//        // Make svc.AssociateDRTLogBucket request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := shield.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockShieldClient struct {
//        shieldiface.ShieldAPI
//    }
//    func (m *mockShieldClient) AssociateDRTLogBucket(input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockShieldClient{}
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
type ShieldAPI interface {
	AssociateDRTLogBucket(*shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error)
	AssociateDRTLogBucketWithContext(aws.Context, *shield.AssociateDRTLogBucketInput, ...request.Option) (*shield.AssociateDRTLogBucketOutput, error)
	AssociateDRTLogBucketRequest(*shield.AssociateDRTLogBucketInput) (*request.Request, *shield.AssociateDRTLogBucketOutput)

	AssociateDRTRole(*shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error)
	AssociateDRTRoleWithContext(aws.Context, *shield.AssociateDRTRoleInput, ...request.Option) (*shield.AssociateDRTRoleOutput, error)
	AssociateDRTRoleRequest(*shield.AssociateDRTRoleInput) (*request.Request, *shield.AssociateDRTRoleOutput)

	CreateProtection(*shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error)
	CreateProtectionWithContext(aws.Context, *shield.CreateProtectionInput, ...request.Option) (*shield.CreateProtectionOutput, error)
	CreateProtectionRequest(*shield.CreateProtectionInput) (*request.Request, *shield.CreateProtectionOutput)

	CreateSubscription(*shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error)
	CreateSubscriptionWithContext(aws.Context, *shield.CreateSubscriptionInput, ...request.Option) (*shield.CreateSubscriptionOutput, error)
	CreateSubscriptionRequest(*shield.CreateSubscriptionInput) (*request.Request, *shield.CreateSubscriptionOutput)

	DeleteProtection(*shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error)
	DeleteProtectionWithContext(aws.Context, *shield.DeleteProtectionInput, ...request.Option) (*shield.DeleteProtectionOutput, error)
	DeleteProtectionRequest(*shield.DeleteProtectionInput) (*request.Request, *shield.DeleteProtectionOutput)

	DeleteSubscription(*shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error)
	DeleteSubscriptionWithContext(aws.Context, *shield.DeleteSubscriptionInput, ...request.Option) (*shield.DeleteSubscriptionOutput, error)
	DeleteSubscriptionRequest(*shield.DeleteSubscriptionInput) (*request.Request, *shield.DeleteSubscriptionOutput)

	DescribeAttack(*shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error)
	DescribeAttackWithContext(aws.Context, *shield.DescribeAttackInput, ...request.Option) (*shield.DescribeAttackOutput, error)
	DescribeAttackRequest(*shield.DescribeAttackInput) (*request.Request, *shield.DescribeAttackOutput)

	DescribeDRTAccess(*shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error)
	DescribeDRTAccessWithContext(aws.Context, *shield.DescribeDRTAccessInput, ...request.Option) (*shield.DescribeDRTAccessOutput, error)
	DescribeDRTAccessRequest(*shield.DescribeDRTAccessInput) (*request.Request, *shield.DescribeDRTAccessOutput)

	DescribeEmergencyContactSettings(*shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error)
	DescribeEmergencyContactSettingsWithContext(aws.Context, *shield.DescribeEmergencyContactSettingsInput, ...request.Option) (*shield.DescribeEmergencyContactSettingsOutput, error)
	DescribeEmergencyContactSettingsRequest(*shield.DescribeEmergencyContactSettingsInput) (*request.Request, *shield.DescribeEmergencyContactSettingsOutput)

	DescribeProtection(*shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error)
	DescribeProtectionWithContext(aws.Context, *shield.DescribeProtectionInput, ...request.Option) (*shield.DescribeProtectionOutput, error)
	DescribeProtectionRequest(*shield.DescribeProtectionInput) (*request.Request, *shield.DescribeProtectionOutput)

	DescribeSubscription(*shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error)
	DescribeSubscriptionWithContext(aws.Context, *shield.DescribeSubscriptionInput, ...request.Option) (*shield.DescribeSubscriptionOutput, error)
	DescribeSubscriptionRequest(*shield.DescribeSubscriptionInput) (*request.Request, *shield.DescribeSubscriptionOutput)

	DisassociateDRTLogBucket(*shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error)
	DisassociateDRTLogBucketWithContext(aws.Context, *shield.DisassociateDRTLogBucketInput, ...request.Option) (*shield.DisassociateDRTLogBucketOutput, error)
	DisassociateDRTLogBucketRequest(*shield.DisassociateDRTLogBucketInput) (*request.Request, *shield.DisassociateDRTLogBucketOutput)

	DisassociateDRTRole(*shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error)
	DisassociateDRTRoleWithContext(aws.Context, *shield.DisassociateDRTRoleInput, ...request.Option) (*shield.DisassociateDRTRoleOutput, error)
	DisassociateDRTRoleRequest(*shield.DisassociateDRTRoleInput) (*request.Request, *shield.DisassociateDRTRoleOutput)

	GetSubscriptionState(*shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error)
	GetSubscriptionStateWithContext(aws.Context, *shield.GetSubscriptionStateInput, ...request.Option) (*shield.GetSubscriptionStateOutput, error)
	GetSubscriptionStateRequest(*shield.GetSubscriptionStateInput) (*request.Request, *shield.GetSubscriptionStateOutput)

	ListAttacks(*shield.ListAttacksInput) (*shield.ListAttacksOutput, error)
	ListAttacksWithContext(aws.Context, *shield.ListAttacksInput, ...request.Option) (*shield.ListAttacksOutput, error)
	ListAttacksRequest(*shield.ListAttacksInput) (*request.Request, *shield.ListAttacksOutput)

	ListProtections(*shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error)
	ListProtectionsWithContext(aws.Context, *shield.ListProtectionsInput, ...request.Option) (*shield.ListProtectionsOutput, error)
	ListProtectionsRequest(*shield.ListProtectionsInput) (*request.Request, *shield.ListProtectionsOutput)

	UpdateEmergencyContactSettings(*shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error)
	UpdateEmergencyContactSettingsWithContext(aws.Context, *shield.UpdateEmergencyContactSettingsInput, ...request.Option) (*shield.UpdateEmergencyContactSettingsOutput, error)
	UpdateEmergencyContactSettingsRequest(*shield.UpdateEmergencyContactSettingsInput) (*request.Request, *shield.UpdateEmergencyContactSettingsOutput)

	UpdateSubscription(*shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error)
	UpdateSubscriptionWithContext(aws.Context, *shield.UpdateSubscriptionInput, ...request.Option) (*shield.UpdateSubscriptionOutput, error)
	UpdateSubscriptionRequest(*shield.UpdateSubscriptionInput) (*request.Request, *shield.UpdateSubscriptionOutput)
}

var _ ShieldAPI = (*shield.Shield)(nil)
