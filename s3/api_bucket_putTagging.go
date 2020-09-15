package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutBucketTaggingInput struct {
	Bucket  *string
	Tagging *Tagging
}

// String returns the string representation
func (p PutBucketTaggingInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketTaggingInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.Tagging == nil {
		invalidParams.Add(request.NewErrParamRequired("Tagging"))
	}
	if p.Tagging != nil {
		if err := p.Tagging.Validate(); err != nil {
			invalidParams.AddNested("Tagging", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketTaggingInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.Tagging != nil {
		v := p.Tagging
		e.SetStruct(helper.PayloadTarget, "Tagging", v)
	}

	return nil
}

type PutBucketTaggingOutput struct{}

// String returns the string representation
func (p PutBucketTaggingOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketTaggingOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketTagging = "PutBucketTagging"

func (c *Client) PutBucketTaggingRequest(input *PutBucketTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketTagging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &PutBucketTaggingInput{}
	}

	req = c.newRequest(op, input, &PutBucketTaggingOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketTagging(input *PutBucketTaggingInput) (PutBucketTaggingOutput, error) {
	req := c.PutBucketTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketTaggingOutput{}, err
	}
	out := req.Data.(*PutBucketTaggingOutput)

	return *out, err
}
