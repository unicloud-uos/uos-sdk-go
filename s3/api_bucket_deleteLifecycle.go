package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type DeleteBucketLifecycleInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketLifecycleInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketLifecycleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketLifecycleInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketLifecycleInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketLifecycleInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketLifecycleOutput struct{}

// String returns the string representation
func (d DeleteBucketLifecycleOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketLifecycleOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketLifecycle = "DeleteBucketLifecycle"

func (c *Client) DeleteBucketLifecycleRequest(input *DeleteBucketLifecycleInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketLifecycle,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &DeleteBucketLifecycleInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketLifecycleOutput{})
	return
}

func (c *Client) DeleteBucketLifecycle(input *DeleteBucketLifecycleInput) (*DeleteBucketLifecycleOutput, error) {
	req := c.DeleteBucketLifecycleRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketLifecycleOutput{}, err
	}
	out := req.Data.(*DeleteBucketLifecycleOutput)

	return out, err
}
