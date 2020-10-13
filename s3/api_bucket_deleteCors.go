package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketCorsInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketCorsInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketCorsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketCorsInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketCorsInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketCorsInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketCorsOutput struct{}

// String returns the string representation
func (d DeleteBucketCorsOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketCorsOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketCors = "DeleteBucketCors"

func (c *Client) DeleteBucketCorsRequest(input *DeleteBucketCorsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketCors,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &DeleteBucketCorsInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketCorsOutput{})
	return
}

func (c *Client) DeleteBucketCors(input *DeleteBucketCorsInput) (*DeleteBucketCorsOutput, error) {
	req := c.DeleteBucketCorsRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketCorsOutput{}, err
	}
	out := req.Data.(*DeleteBucketCorsOutput)

	return out, err
}
