// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package securityhubiface provides an interface to enable mocking the AWS SecurityHub service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package securityhubiface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/securityhub"
)

// SecurityHubAPI provides an interface to enable mocking the
// securityhub.SecurityHub service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS SecurityHub.
//    func myFunc(svc securityhubiface.SecurityHubAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := securityhub.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSecurityHubClient struct {
//        securityhubiface.SecurityHubAPI
//    }
//    func (m *mockSecurityHubClient) AcceptInvitation(input *securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSecurityHubClient{}
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
type SecurityHubAPI interface {
	AcceptInvitation(*securityhub.AcceptInvitationInput) (*securityhub.AcceptInvitationOutput, error)
	AcceptInvitationWithContext(aws.Context, *securityhub.AcceptInvitationInput, ...request.Option) (*securityhub.AcceptInvitationOutput, error)
	AcceptInvitationRequest(*securityhub.AcceptInvitationInput) (*request.Request, *securityhub.AcceptInvitationOutput)

	BatchDisableStandards(*securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error)
	BatchDisableStandardsWithContext(aws.Context, *securityhub.BatchDisableStandardsInput, ...request.Option) (*securityhub.BatchDisableStandardsOutput, error)
	BatchDisableStandardsRequest(*securityhub.BatchDisableStandardsInput) (*request.Request, *securityhub.BatchDisableStandardsOutput)

	BatchEnableStandards(*securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error)
	BatchEnableStandardsWithContext(aws.Context, *securityhub.BatchEnableStandardsInput, ...request.Option) (*securityhub.BatchEnableStandardsOutput, error)
	BatchEnableStandardsRequest(*securityhub.BatchEnableStandardsInput) (*request.Request, *securityhub.BatchEnableStandardsOutput)

	BatchImportFindings(*securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error)
	BatchImportFindingsWithContext(aws.Context, *securityhub.BatchImportFindingsInput, ...request.Option) (*securityhub.BatchImportFindingsOutput, error)
	BatchImportFindingsRequest(*securityhub.BatchImportFindingsInput) (*request.Request, *securityhub.BatchImportFindingsOutput)

	CreateInsight(*securityhub.CreateInsightInput) (*securityhub.CreateInsightOutput, error)
	CreateInsightWithContext(aws.Context, *securityhub.CreateInsightInput, ...request.Option) (*securityhub.CreateInsightOutput, error)
	CreateInsightRequest(*securityhub.CreateInsightInput) (*request.Request, *securityhub.CreateInsightOutput)

	CreateMembers(*securityhub.CreateMembersInput) (*securityhub.CreateMembersOutput, error)
	CreateMembersWithContext(aws.Context, *securityhub.CreateMembersInput, ...request.Option) (*securityhub.CreateMembersOutput, error)
	CreateMembersRequest(*securityhub.CreateMembersInput) (*request.Request, *securityhub.CreateMembersOutput)

	DeclineInvitations(*securityhub.DeclineInvitationsInput) (*securityhub.DeclineInvitationsOutput, error)
	DeclineInvitationsWithContext(aws.Context, *securityhub.DeclineInvitationsInput, ...request.Option) (*securityhub.DeclineInvitationsOutput, error)
	DeclineInvitationsRequest(*securityhub.DeclineInvitationsInput) (*request.Request, *securityhub.DeclineInvitationsOutput)

	DeleteInsight(*securityhub.DeleteInsightInput) (*securityhub.DeleteInsightOutput, error)
	DeleteInsightWithContext(aws.Context, *securityhub.DeleteInsightInput, ...request.Option) (*securityhub.DeleteInsightOutput, error)
	DeleteInsightRequest(*securityhub.DeleteInsightInput) (*request.Request, *securityhub.DeleteInsightOutput)

	DeleteInvitations(*securityhub.DeleteInvitationsInput) (*securityhub.DeleteInvitationsOutput, error)
	DeleteInvitationsWithContext(aws.Context, *securityhub.DeleteInvitationsInput, ...request.Option) (*securityhub.DeleteInvitationsOutput, error)
	DeleteInvitationsRequest(*securityhub.DeleteInvitationsInput) (*request.Request, *securityhub.DeleteInvitationsOutput)

	DeleteMembers(*securityhub.DeleteMembersInput) (*securityhub.DeleteMembersOutput, error)
	DeleteMembersWithContext(aws.Context, *securityhub.DeleteMembersInput, ...request.Option) (*securityhub.DeleteMembersOutput, error)
	DeleteMembersRequest(*securityhub.DeleteMembersInput) (*request.Request, *securityhub.DeleteMembersOutput)

	DisableImportFindingsForProduct(*securityhub.DisableImportFindingsForProductInput) (*securityhub.DisableImportFindingsForProductOutput, error)
	DisableImportFindingsForProductWithContext(aws.Context, *securityhub.DisableImportFindingsForProductInput, ...request.Option) (*securityhub.DisableImportFindingsForProductOutput, error)
	DisableImportFindingsForProductRequest(*securityhub.DisableImportFindingsForProductInput) (*request.Request, *securityhub.DisableImportFindingsForProductOutput)

	DisableSecurityHub(*securityhub.DisableSecurityHubInput) (*securityhub.DisableSecurityHubOutput, error)
	DisableSecurityHubWithContext(aws.Context, *securityhub.DisableSecurityHubInput, ...request.Option) (*securityhub.DisableSecurityHubOutput, error)
	DisableSecurityHubRequest(*securityhub.DisableSecurityHubInput) (*request.Request, *securityhub.DisableSecurityHubOutput)

	DisassociateFromMasterAccount(*securityhub.DisassociateFromMasterAccountInput) (*securityhub.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountWithContext(aws.Context, *securityhub.DisassociateFromMasterAccountInput, ...request.Option) (*securityhub.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountRequest(*securityhub.DisassociateFromMasterAccountInput) (*request.Request, *securityhub.DisassociateFromMasterAccountOutput)

	DisassociateMembers(*securityhub.DisassociateMembersInput) (*securityhub.DisassociateMembersOutput, error)
	DisassociateMembersWithContext(aws.Context, *securityhub.DisassociateMembersInput, ...request.Option) (*securityhub.DisassociateMembersOutput, error)
	DisassociateMembersRequest(*securityhub.DisassociateMembersInput) (*request.Request, *securityhub.DisassociateMembersOutput)

	EnableImportFindingsForProduct(*securityhub.EnableImportFindingsForProductInput) (*securityhub.EnableImportFindingsForProductOutput, error)
	EnableImportFindingsForProductWithContext(aws.Context, *securityhub.EnableImportFindingsForProductInput, ...request.Option) (*securityhub.EnableImportFindingsForProductOutput, error)
	EnableImportFindingsForProductRequest(*securityhub.EnableImportFindingsForProductInput) (*request.Request, *securityhub.EnableImportFindingsForProductOutput)

	EnableSecurityHub(*securityhub.EnableSecurityHubInput) (*securityhub.EnableSecurityHubOutput, error)
	EnableSecurityHubWithContext(aws.Context, *securityhub.EnableSecurityHubInput, ...request.Option) (*securityhub.EnableSecurityHubOutput, error)
	EnableSecurityHubRequest(*securityhub.EnableSecurityHubInput) (*request.Request, *securityhub.EnableSecurityHubOutput)

	GetEnabledStandards(*securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error)
	GetEnabledStandardsWithContext(aws.Context, *securityhub.GetEnabledStandardsInput, ...request.Option) (*securityhub.GetEnabledStandardsOutput, error)
	GetEnabledStandardsRequest(*securityhub.GetEnabledStandardsInput) (*request.Request, *securityhub.GetEnabledStandardsOutput)

	GetFindings(*securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error)
	GetFindingsWithContext(aws.Context, *securityhub.GetFindingsInput, ...request.Option) (*securityhub.GetFindingsOutput, error)
	GetFindingsRequest(*securityhub.GetFindingsInput) (*request.Request, *securityhub.GetFindingsOutput)

	GetFindingsPages(*securityhub.GetFindingsInput, func(*securityhub.GetFindingsOutput, bool) bool) error
	GetFindingsPagesWithContext(aws.Context, *securityhub.GetFindingsInput, func(*securityhub.GetFindingsOutput, bool) bool, ...request.Option) error

	GetInsightResults(*securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error)
	GetInsightResultsWithContext(aws.Context, *securityhub.GetInsightResultsInput, ...request.Option) (*securityhub.GetInsightResultsOutput, error)
	GetInsightResultsRequest(*securityhub.GetInsightResultsInput) (*request.Request, *securityhub.GetInsightResultsOutput)

	GetInsights(*securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error)
	GetInsightsWithContext(aws.Context, *securityhub.GetInsightsInput, ...request.Option) (*securityhub.GetInsightsOutput, error)
	GetInsightsRequest(*securityhub.GetInsightsInput) (*request.Request, *securityhub.GetInsightsOutput)

	GetInsightsPages(*securityhub.GetInsightsInput, func(*securityhub.GetInsightsOutput, bool) bool) error
	GetInsightsPagesWithContext(aws.Context, *securityhub.GetInsightsInput, func(*securityhub.GetInsightsOutput, bool) bool, ...request.Option) error

	GetInvitationsCount(*securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error)
	GetInvitationsCountWithContext(aws.Context, *securityhub.GetInvitationsCountInput, ...request.Option) (*securityhub.GetInvitationsCountOutput, error)
	GetInvitationsCountRequest(*securityhub.GetInvitationsCountInput) (*request.Request, *securityhub.GetInvitationsCountOutput)

	GetMasterAccount(*securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error)
	GetMasterAccountWithContext(aws.Context, *securityhub.GetMasterAccountInput, ...request.Option) (*securityhub.GetMasterAccountOutput, error)
	GetMasterAccountRequest(*securityhub.GetMasterAccountInput) (*request.Request, *securityhub.GetMasterAccountOutput)

	GetMembers(*securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error)
	GetMembersWithContext(aws.Context, *securityhub.GetMembersInput, ...request.Option) (*securityhub.GetMembersOutput, error)
	GetMembersRequest(*securityhub.GetMembersInput) (*request.Request, *securityhub.GetMembersOutput)

	InviteMembers(*securityhub.InviteMembersInput) (*securityhub.InviteMembersOutput, error)
	InviteMembersWithContext(aws.Context, *securityhub.InviteMembersInput, ...request.Option) (*securityhub.InviteMembersOutput, error)
	InviteMembersRequest(*securityhub.InviteMembersInput) (*request.Request, *securityhub.InviteMembersOutput)

	ListEnabledProductsForImport(*securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListEnabledProductsForImportWithContext(aws.Context, *securityhub.ListEnabledProductsForImportInput, ...request.Option) (*securityhub.ListEnabledProductsForImportOutput, error)
	ListEnabledProductsForImportRequest(*securityhub.ListEnabledProductsForImportInput) (*request.Request, *securityhub.ListEnabledProductsForImportOutput)

	ListEnabledProductsForImportPages(*securityhub.ListEnabledProductsForImportInput, func(*securityhub.ListEnabledProductsForImportOutput, bool) bool) error
	ListEnabledProductsForImportPagesWithContext(aws.Context, *securityhub.ListEnabledProductsForImportInput, func(*securityhub.ListEnabledProductsForImportOutput, bool) bool, ...request.Option) error

	ListInvitations(*securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error)
	ListInvitationsWithContext(aws.Context, *securityhub.ListInvitationsInput, ...request.Option) (*securityhub.ListInvitationsOutput, error)
	ListInvitationsRequest(*securityhub.ListInvitationsInput) (*request.Request, *securityhub.ListInvitationsOutput)

	ListMembers(*securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error)
	ListMembersWithContext(aws.Context, *securityhub.ListMembersInput, ...request.Option) (*securityhub.ListMembersOutput, error)
	ListMembersRequest(*securityhub.ListMembersInput) (*request.Request, *securityhub.ListMembersOutput)

	UpdateFindings(*securityhub.UpdateFindingsInput) (*securityhub.UpdateFindingsOutput, error)
	UpdateFindingsWithContext(aws.Context, *securityhub.UpdateFindingsInput, ...request.Option) (*securityhub.UpdateFindingsOutput, error)
	UpdateFindingsRequest(*securityhub.UpdateFindingsInput) (*request.Request, *securityhub.UpdateFindingsOutput)

	UpdateInsight(*securityhub.UpdateInsightInput) (*securityhub.UpdateInsightOutput, error)
	UpdateInsightWithContext(aws.Context, *securityhub.UpdateInsightInput, ...request.Option) (*securityhub.UpdateInsightOutput, error)
	UpdateInsightRequest(*securityhub.UpdateInsightInput) (*request.Request, *securityhub.UpdateInsightOutput)
}

var _ SecurityHubAPI = (*securityhub.SecurityHub)(nil)
