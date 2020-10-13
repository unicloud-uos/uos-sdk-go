package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteBucketWebsiteInput struct {
	Bucket *string
}

// String returns the string representation
func (d DeleteBucketWebsiteInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketWebsiteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketWebsiteInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteBucketWebsiteInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteBucketWebsiteInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type DeleteBucketWebsiteOutput struct{}

// String returns the string representation
func (d DeleteBucketWebsiteOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteBucketWebsiteOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteBucketWebsite = "DeleteBucketWebsite"

func (c *Client) DeleteBucketWebsiteRequest(input *DeleteBucketWebsiteInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteBucketWebsite,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?website",
	}

	if input == nil {
		input = &DeleteBucketWebsiteInput{}
	}

	req = c.newRequest(op, input, &DeleteBucketWebsiteOutput{})
	return
}

func (c *Client) DeleteBucketWebsite(input *DeleteBucketWebsiteInput) (*DeleteBucketWebsiteOutput, error) {
	req := c.DeleteBucketWebsiteRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteBucketWebsiteOutput{}, err
	}
	out := req.Data.(*DeleteBucketWebsiteOutput)

	return out, err
}
