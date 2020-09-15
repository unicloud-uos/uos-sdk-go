package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutBucketPolicyInput struct {
	Bucket *string
	Policy *string
}

// String returns the string representation
func (p PutBucketPolicyInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketPolicyInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.Policy == nil {
		invalidParams.Add(request.NewErrParamRequired("Policy"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketPolicyInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketPolicyInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.Policy != nil {
		v := *p.Policy
		e.SetStream(helper.PayloadTarget, "Policy", request.StringStream(v))
	}

	return nil
}

type PutBucketPolicyOutput struct{}

// String returns the string representation
func (p PutBucketPolicyOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketPolicyOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketPolicy = "PutBucketPolicy"

func (c *Client) PutBucketPolicyRequest(input *PutBucketPolicyInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?policy",
	}

	if input == nil {
		input = &PutBucketPolicyInput{}
	}

	req = c.newRequest(op, input, &PutBucketPolicyOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketPolicy(input *PutBucketPolicyInput) (PutBucketPolicyOutput, error) {
	req := c.PutBucketPolicyRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketPolicyOutput{}, err
	}
	out := req.Data.(*PutBucketPolicyOutput)

	return *out, err
}
