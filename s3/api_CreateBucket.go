package s3

import (
	"context"

	"github.com/uos-sdk-go/s3/auxiliary"
	. "github.com/uos-sdk-go/s3/request"
)

type CreateBucketInput struct {
	_ struct{} `type:"structure" payload:"CreateBucketConfiguration"`

	// The ACL to apply to the bucket.
	ACL auxiliary.BucketACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// The name of the bucket to create.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

type CreateBucketOutput struct {
	_ struct{} `type:"structure"`

	Location *string `location:"header" locationName:"Location" type:"string"`
}

const CreateBucket = "CreateBucket"

func (c *Client) CreateBucketRequest(input *CreateBucketInput) (req *Request, output *CreateBucketOutput) {
	op := &Operation{
		Name:       CreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	output = &CreateBucketOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *Client) CreateBucket(input *CreateBucketInput, ctx context.Context) (*CreateBucketOutput, error) {
	req, out := c.CreateBucketRequest(input)
	req.SetContext(ctx)
	return out, req.Do()
}
