package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketPolicyInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketPolicyInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketPolicyInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketPolicyInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketPolicyInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketPolicyOutput struct{}

// String returns the string representation
func (d DeleteBucketPolicyOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketPolicyOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketPolicy = "DeleteBucketPolicy"

func (c *Client) DeleteBucketPolicyRequest(input *DeleteBucketPolicyInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?policy",
	}

	if input == nil {
		input = &DeleteBucketPolicyInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketPolicyOutput{})
	return
}

func (c *Client) DeleteBucketPolicy(input *DeleteBucketPolicyInput) (*DeleteBucketPolicyOutput, error) {
	req := c.DeleteBucketPolicyRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketPolicyOutput{}, err
	}
	out := req.Data.(*DeleteBucketPolicyOutput)

	return out, err
}
