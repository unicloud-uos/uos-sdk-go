package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type CreateBucketInput struct {
	ACL    *string
	Bucket *string
}

// String returns the string representation
func (c CreateBucketInput) String() string {
	return helper.Prettify(c)
}

func (c CreateBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBucketInput"}

	if c.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (c CreateBucketInput) GetBucketName() (v string) {
	if c.Bucket == nil {
		return v
	}
	return *c.Bucket
}

func (c CreateBucketInput) MarshalForPut(e *request.EncoderForPut) error {
	if c.Bucket != nil {
		v := *c.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if c.ACL != nil {
		v := *c.ACL
		e.SetValue(helper.HeaderTarget, "x-uos-acl", request.StringValue(v))
	}

	return nil
}

type CreateBucketOutput struct {
	Location *string `location:"header" locationName:"Location"`
}

// String returns the string representation
func (c CreateBucketOutput) String() string {
	return helper.Prettify(c)
}

func (c CreateBucketOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const CreateBucket = "CreateBucket"

func (c *Client) CreateBucketRequest(input *CreateBucketInput) (req *request.Request) {
	op := &request.Operation{
		Name:       CreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	req = c.newRequest(op, input, &CreateBucketOutput{})
	return
}

func (c *Client) CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error) {
	req := c.CreateBucketRequest(input)

	err := req.Do()
	if err != nil {
		return &CreateBucketOutput{}, err
	}
	out := req.Data.(*CreateBucketOutput)

	return out, err
}
