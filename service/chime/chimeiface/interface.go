// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chimeiface provides an interface to enable mocking the Amazon Chime service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chimeiface

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/service/chime"
)

// ChimeAPI provides an interface to enable mocking the
// chime.Chime service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Chime.
//    func myFunc(svc chimeiface.ChimeAPI) bool {
//        // Make svc.BatchSuspendUser request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := chime.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockChimeClient struct {
//        chimeiface.ChimeAPI
//    }
//    func (m *mockChimeClient) BatchSuspendUser(input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockChimeClient{}
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
type ChimeAPI interface {
	BatchSuspendUser(*chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserWithContext(aws.Context, *chime.BatchSuspendUserInput, ...request.Option) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserRequest(*chime.BatchSuspendUserInput) (*request.Request, *chime.BatchSuspendUserOutput)

	BatchUnsuspendUser(*chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserWithContext(aws.Context, *chime.BatchUnsuspendUserInput, ...request.Option) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserRequest(*chime.BatchUnsuspendUserInput) (*request.Request, *chime.BatchUnsuspendUserOutput)

	BatchUpdateUser(*chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserWithContext(aws.Context, *chime.BatchUpdateUserInput, ...request.Option) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserRequest(*chime.BatchUpdateUserInput) (*request.Request, *chime.BatchUpdateUserOutput)

	CreateAccount(*chime.CreateAccountInput) (*chime.CreateAccountOutput, error)
	CreateAccountWithContext(aws.Context, *chime.CreateAccountInput, ...request.Option) (*chime.CreateAccountOutput, error)
	CreateAccountRequest(*chime.CreateAccountInput) (*request.Request, *chime.CreateAccountOutput)

	DeleteAccount(*chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error)
	DeleteAccountWithContext(aws.Context, *chime.DeleteAccountInput, ...request.Option) (*chime.DeleteAccountOutput, error)
	DeleteAccountRequest(*chime.DeleteAccountInput) (*request.Request, *chime.DeleteAccountOutput)

	GetAccount(*chime.GetAccountInput) (*chime.GetAccountOutput, error)
	GetAccountWithContext(aws.Context, *chime.GetAccountInput, ...request.Option) (*chime.GetAccountOutput, error)
	GetAccountRequest(*chime.GetAccountInput) (*request.Request, *chime.GetAccountOutput)

	GetAccountSettings(*chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsWithContext(aws.Context, *chime.GetAccountSettingsInput, ...request.Option) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsRequest(*chime.GetAccountSettingsInput) (*request.Request, *chime.GetAccountSettingsOutput)

	GetUser(*chime.GetUserInput) (*chime.GetUserOutput, error)
	GetUserWithContext(aws.Context, *chime.GetUserInput, ...request.Option) (*chime.GetUserOutput, error)
	GetUserRequest(*chime.GetUserInput) (*request.Request, *chime.GetUserOutput)

	InviteUsers(*chime.InviteUsersInput) (*chime.InviteUsersOutput, error)
	InviteUsersWithContext(aws.Context, *chime.InviteUsersInput, ...request.Option) (*chime.InviteUsersOutput, error)
	InviteUsersRequest(*chime.InviteUsersInput) (*request.Request, *chime.InviteUsersOutput)

	ListAccounts(*chime.ListAccountsInput) (*chime.ListAccountsOutput, error)
	ListAccountsWithContext(aws.Context, *chime.ListAccountsInput, ...request.Option) (*chime.ListAccountsOutput, error)
	ListAccountsRequest(*chime.ListAccountsInput) (*request.Request, *chime.ListAccountsOutput)

	ListAccountsPages(*chime.ListAccountsInput, func(*chime.ListAccountsOutput, bool) bool) error
	ListAccountsPagesWithContext(aws.Context, *chime.ListAccountsInput, func(*chime.ListAccountsOutput, bool) bool, ...request.Option) error

	ListUsers(*chime.ListUsersInput) (*chime.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *chime.ListUsersInput, ...request.Option) (*chime.ListUsersOutput, error)
	ListUsersRequest(*chime.ListUsersInput) (*request.Request, *chime.ListUsersOutput)

	ListUsersPages(*chime.ListUsersInput, func(*chime.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *chime.ListUsersInput, func(*chime.ListUsersOutput, bool) bool, ...request.Option) error

	LogoutUser(*chime.LogoutUserInput) (*chime.LogoutUserOutput, error)
	LogoutUserWithContext(aws.Context, *chime.LogoutUserInput, ...request.Option) (*chime.LogoutUserOutput, error)
	LogoutUserRequest(*chime.LogoutUserInput) (*request.Request, *chime.LogoutUserOutput)

	ResetPersonalPIN(*chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINWithContext(aws.Context, *chime.ResetPersonalPINInput, ...request.Option) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINRequest(*chime.ResetPersonalPINInput) (*request.Request, *chime.ResetPersonalPINOutput)

	UpdateAccount(*chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error)
	UpdateAccountWithContext(aws.Context, *chime.UpdateAccountInput, ...request.Option) (*chime.UpdateAccountOutput, error)
	UpdateAccountRequest(*chime.UpdateAccountInput) (*request.Request, *chime.UpdateAccountOutput)

	UpdateAccountSettings(*chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsWithContext(aws.Context, *chime.UpdateAccountSettingsInput, ...request.Option) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsRequest(*chime.UpdateAccountSettingsInput) (*request.Request, *chime.UpdateAccountSettingsOutput)

	UpdateUser(*chime.UpdateUserInput) (*chime.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *chime.UpdateUserInput, ...request.Option) (*chime.UpdateUserOutput, error)
	UpdateUserRequest(*chime.UpdateUserInput) (*request.Request, *chime.UpdateUserOutput)
}

var _ ChimeAPI = (*chime.Chime)(nil)
