package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketOutput struct{}

// String returns the string representation
func (d DeleteBucketOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucket = "DeleteBucket"

func (c *Client) DeleteBucketRequest(input *DeleteBucketInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucket,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &DeleteBucketInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketOutput{})
	return
}

func (c *Client) DeleteBucket(input *DeleteBucketInput) (*DeleteBucketOutput, error) {
	req := c.DeleteBucketRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketOutput{}, err
	}
	out := req.Data.(*DeleteBucketOutput)

	return out, err
}
