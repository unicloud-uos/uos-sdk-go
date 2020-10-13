package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type HeadBucketInput struct {
	Bucket *string
}

// String returns the string representation
func (g HeadBucketInput) String() string {
	return helper.Prettify(g)
}

func (g HeadBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HeadBucketInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g HeadBucketInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g HeadBucketInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type HeadBucketOutput struct{}

// String returns the string representation
func (d HeadBucketOutput) String() string {
	return helper.Prettify(d)
}

func (d HeadBucketOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opHeadBucket = "HeadBucket"

func (c *Client) HeadBucketRequest(input *HeadBucketInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opHeadBucket,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &HeadBucketInput{}
	}

	req = c.newRequest(op, input, &HeadBucketOutput{})
	return
}

func (c *Client) HeadBucket(input *HeadBucketInput) (*HeadBucketOutput, error) {
	req := c.HeadBucketRequest(input)

	err := req.Do()
	if err != nil {
		return &HeadBucketOutput{}, err
	}
	out := req.Data.(*HeadBucketOutput)

	return out, err
}
