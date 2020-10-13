package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutBucketLoggingInput struct {
	Bucket              *string
	BucketLoggingStatus *BucketLoggingStatus
}

// String returns the string representation
func (p PutBucketLoggingInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketLoggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketLoggingInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.BucketLoggingStatus == nil {
		invalidParams.Add(request.NewErrParamRequired("BucketLoggingStatus"))
	}
	if p.BucketLoggingStatus != nil {
		if err := p.BucketLoggingStatus.Validate(); err != nil {
			invalidParams.AddNested("BucketLoggingStatus", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketLoggingInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketLoggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.BucketLoggingStatus != nil {
		v := p.BucketLoggingStatus
		e.SetStruct(helper.PayloadTarget, "BucketLoggingStatus", v)
	}

	return nil
}

type PutBucketLoggingOutput struct{}

// String returns the string representation
func (p PutBucketLoggingOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketLoggingOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketLogging = "PutBucketLogging"

func (c *Client) PutBucketLoggingRequest(input *PutBucketLoggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketLogging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?logging",
	}

	if input == nil {
		input = &PutBucketLoggingInput{}
	}

	req = c.newRequest(op, input, &PutBucketLoggingOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketLogging(input *PutBucketLoggingInput) (PutBucketLoggingOutput, error) {
	req := c.PutBucketLoggingRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketLoggingOutput{}, err
	}
	out := req.Data.(*PutBucketLoggingOutput)

	return *out, err
}
