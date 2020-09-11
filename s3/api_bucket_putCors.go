package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutBucketCorsInput struct {
	Bucket            *string
	CORSConfiguration *CORSConfiguration
}

// String returns the string representation
func (p PutBucketCorsInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketCorsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketCorsInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if p.CORSConfiguration != nil {
		if err := p.CORSConfiguration.Validate(); err != nil {
			invalidParams.AddNested("CORSConfiguration", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketCorsInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketCorsInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.CORSConfiguration != nil {
		v := p.CORSConfiguration
		e.SetStruct(helper.PayloadTarget, "CORSConfiguration", v)
	}

	return nil
}

type PutBucketCorsOutput struct{}

// String returns the string representation
func (p PutBucketCorsOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketCorsOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketCors = "PutBucketCors"

func (c *Client) PutBucketCorsRequest(input *PutBucketCorsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketCors,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &PutBucketCorsInput{}
	}

	req = c.newRequest(op, input, &PutBucketCorsOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketCors(input *PutBucketCorsInput) (PutBucketCorsOutput, error) {
	req := c.PutBucketCorsRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketCorsOutput{}, err
	}
	out := req.Data.(*PutBucketCorsOutput)

	return *out, err
}
