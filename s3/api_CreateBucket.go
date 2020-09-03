package s3

import (
	"encoding/xml"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type CreateBucketInput struct {
	_ struct{} `type:"structure" payload:"CreateBucketConfiguration"`

	// The ACL to apply to the bucket.
	ACL helper.BucketACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// The name of the bucket to create.
	//
	// Bucket is a required field
	Bucket string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

func (c CreateBucketInput) MarshalForPut(e *request.EncoderForPut) error {

	if c.Bucket != "" {
		v := c.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	return nil
}

func (c CreateBucketInput) GetBucketName() (v string) {
	if c.Bucket == "" {
		return v
	}
	return c.Bucket
}

func (c *CreateBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBucketInput"}

	if c.Bucket == "" {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateBucketOutput struct {
	_ struct{} `type:"structure"`

	Location *string `location:"header" locationName:"Location" type:"string"`
}

func (c CreateBucketOutput) UnmarshalHeader(r *request.Request) error {
	*c.Location = r.HTTPResponse.Header.Get("Location")
	r.Data = c
	return nil
}

func (c CreateBucketOutput) UnmarshalBody(coder *xml.Decoder) error {
	return nil
}

const CreateBucket = "CreateBucket"

func (c *Client) CreateBucketRequest(input *CreateBucketInput) (req *request.Request, output *CreateBucketOutput) {
	op := &request.Operation{
		Name:       CreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	output = &CreateBucketOutput{}
	req = c.newRequest(op, input, output)

	c.Logger.Error("#@@@@@@req: ", req)

	return
}

func (c *Client) CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error) {
	c.Logger.Debug("Create bucket input: ", input)
	req, out := c.CreateBucketRequest(input)
	c.Logger.Debug("Create bucket request: ", req)

	return out, req.Do()
}
