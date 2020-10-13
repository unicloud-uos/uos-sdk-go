package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketTaggingInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketTaggingInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketTaggingInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketTaggingInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketTaggingOutput struct{}

// String returns the string representation
func (d DeleteBucketTaggingOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketTaggingOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketTagging = "DeleteBucketTagging"

func (c *Client) DeleteBucketTaggingRequest(input *DeleteBucketTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketTagging,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &DeleteBucketTaggingInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketTaggingOutput{})
	return
}

func (c *Client) DeleteBucketTagging(input *DeleteBucketTaggingInput) (*DeleteBucketTaggingOutput, error) {
	req := c.DeleteBucketTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketTaggingOutput{}, err
	}
	out := req.Data.(*DeleteBucketTaggingOutput)

	return out, err
}
