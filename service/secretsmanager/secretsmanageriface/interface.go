// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package secretsmanageriface provides an interface to enable mocking the AWS Secrets Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package secretsmanageriface

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"github.com/unicloud-uos/uos-sdk-go/service/secretsmanager"
)

// SecretsManagerAPI provides an interface to enable mocking the
// secretsmanager.SecretsManager service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Secrets Manager.
//    func myFunc(svc secretsmanageriface.SecretsManagerAPI) bool {
//        // Make svc.CancelRotateSecret request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := secretsmanager.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSecretsManagerClient struct {
//        secretsmanageriface.SecretsManagerAPI
//    }
//    func (m *mockSecretsManagerClient) CancelRotateSecret(input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSecretsManagerClient{}
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
type SecretsManagerAPI interface {
	CancelRotateSecret(*secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error)
	CancelRotateSecretWithContext(aws.Context, *secretsmanager.CancelRotateSecretInput, ...request.Option) (*secretsmanager.CancelRotateSecretOutput, error)
	CancelRotateSecretRequest(*secretsmanager.CancelRotateSecretInput) (*request.Request, *secretsmanager.CancelRotateSecretOutput)

	CreateSecret(*secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error)
	CreateSecretWithContext(aws.Context, *secretsmanager.CreateSecretInput, ...request.Option) (*secretsmanager.CreateSecretOutput, error)
	CreateSecretRequest(*secretsmanager.CreateSecretInput) (*request.Request, *secretsmanager.CreateSecretOutput)

	DeleteResourcePolicy(*secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *secretsmanager.DeleteResourcePolicyInput, ...request.Option) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*secretsmanager.DeleteResourcePolicyInput) (*request.Request, *secretsmanager.DeleteResourcePolicyOutput)

	DeleteSecret(*secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error)
	DeleteSecretWithContext(aws.Context, *secretsmanager.DeleteSecretInput, ...request.Option) (*secretsmanager.DeleteSecretOutput, error)
	DeleteSecretRequest(*secretsmanager.DeleteSecretInput) (*request.Request, *secretsmanager.DeleteSecretOutput)

	DescribeSecret(*secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error)
	DescribeSecretWithContext(aws.Context, *secretsmanager.DescribeSecretInput, ...request.Option) (*secretsmanager.DescribeSecretOutput, error)
	DescribeSecretRequest(*secretsmanager.DescribeSecretInput) (*request.Request, *secretsmanager.DescribeSecretOutput)

	GetRandomPassword(*secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error)
	GetRandomPasswordWithContext(aws.Context, *secretsmanager.GetRandomPasswordInput, ...request.Option) (*secretsmanager.GetRandomPasswordOutput, error)
	GetRandomPasswordRequest(*secretsmanager.GetRandomPasswordInput) (*request.Request, *secretsmanager.GetRandomPasswordOutput)

	GetResourcePolicy(*secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error)
	GetResourcePolicyWithContext(aws.Context, *secretsmanager.GetResourcePolicyInput, ...request.Option) (*secretsmanager.GetResourcePolicyOutput, error)
	GetResourcePolicyRequest(*secretsmanager.GetResourcePolicyInput) (*request.Request, *secretsmanager.GetResourcePolicyOutput)

	GetSecretValue(*secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
	GetSecretValueWithContext(aws.Context, *secretsmanager.GetSecretValueInput, ...request.Option) (*secretsmanager.GetSecretValueOutput, error)
	GetSecretValueRequest(*secretsmanager.GetSecretValueInput) (*request.Request, *secretsmanager.GetSecretValueOutput)

	ListSecretVersionIds(*secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsWithContext(aws.Context, *secretsmanager.ListSecretVersionIdsInput, ...request.Option) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsRequest(*secretsmanager.ListSecretVersionIdsInput) (*request.Request, *secretsmanager.ListSecretVersionIdsOutput)

	ListSecretVersionIdsPages(*secretsmanager.ListSecretVersionIdsInput, func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool) error
	ListSecretVersionIdsPagesWithContext(aws.Context, *secretsmanager.ListSecretVersionIdsInput, func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool, ...request.Option) error

	ListSecrets(*secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsWithContext(aws.Context, *secretsmanager.ListSecretsInput, ...request.Option) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsRequest(*secretsmanager.ListSecretsInput) (*request.Request, *secretsmanager.ListSecretsOutput)

	ListSecretsPages(*secretsmanager.ListSecretsInput, func(*secretsmanager.ListSecretsOutput, bool) bool) error
	ListSecretsPagesWithContext(aws.Context, *secretsmanager.ListSecretsInput, func(*secretsmanager.ListSecretsOutput, bool) bool, ...request.Option) error

	PutResourcePolicy(*secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *secretsmanager.PutResourcePolicyInput, ...request.Option) (*secretsmanager.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*secretsmanager.PutResourcePolicyInput) (*request.Request, *secretsmanager.PutResourcePolicyOutput)

	PutSecretValue(*secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error)
	PutSecretValueWithContext(aws.Context, *secretsmanager.PutSecretValueInput, ...request.Option) (*secretsmanager.PutSecretValueOutput, error)
	PutSecretValueRequest(*secretsmanager.PutSecretValueInput) (*request.Request, *secretsmanager.PutSecretValueOutput)

	RestoreSecret(*secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error)
	RestoreSecretWithContext(aws.Context, *secretsmanager.RestoreSecretInput, ...request.Option) (*secretsmanager.RestoreSecretOutput, error)
	RestoreSecretRequest(*secretsmanager.RestoreSecretInput) (*request.Request, *secretsmanager.RestoreSecretOutput)

	RotateSecret(*secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error)
	RotateSecretWithContext(aws.Context, *secretsmanager.RotateSecretInput, ...request.Option) (*secretsmanager.RotateSecretOutput, error)
	RotateSecretRequest(*secretsmanager.RotateSecretInput) (*request.Request, *secretsmanager.RotateSecretOutput)

	TagResource(*secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *secretsmanager.TagResourceInput, ...request.Option) (*secretsmanager.TagResourceOutput, error)
	TagResourceRequest(*secretsmanager.TagResourceInput) (*request.Request, *secretsmanager.TagResourceOutput)

	UntagResource(*secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *secretsmanager.UntagResourceInput, ...request.Option) (*secretsmanager.UntagResourceOutput, error)
	UntagResourceRequest(*secretsmanager.UntagResourceInput) (*request.Request, *secretsmanager.UntagResourceOutput)

	UpdateSecret(*secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretWithContext(aws.Context, *secretsmanager.UpdateSecretInput, ...request.Option) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretRequest(*secretsmanager.UpdateSecretInput) (*request.Request, *secretsmanager.UpdateSecretOutput)

	UpdateSecretVersionStage(*secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	UpdateSecretVersionStageWithContext(aws.Context, *secretsmanager.UpdateSecretVersionStageInput, ...request.Option) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	UpdateSecretVersionStageRequest(*secretsmanager.UpdateSecretVersionStageInput) (*request.Request, *secretsmanager.UpdateSecretVersionStageOutput)
}

var _ SecretsManagerAPI = (*secretsmanager.SecretsManager)(nil)
