package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketEncryptionInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketEncryptionInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketEncryptionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketEncryptionInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketEncryptionInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketEncryptionInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketEncryptionOutput struct{}

// String returns the string representation
func (d DeleteBucketEncryptionOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketEncryptionOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketEncryption = "DeleteBucketEncryption"

func (c *Client) DeleteBucketEncryptionRequest(input *DeleteBucketEncryptionInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketEncryption,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &DeleteBucketEncryptionInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketEncryptionOutput{})
	return
}

func (c *Client) DeleteBucketEncryption(input *DeleteBucketEncryptionInput) (*DeleteBucketEncryptionOutput, error) {
	req := c.DeleteBucketEncryptionRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketEncryptionOutput{}, err
	}
	out := req.Data.(*DeleteBucketEncryptionOutput)

	return out, err
}
