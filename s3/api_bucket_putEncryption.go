package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutBucketEncryptionInput struct {
	Bucket                            *string
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration
}

// String returns the string representation
func (p PutBucketEncryptionInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketEncryptionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketEncryptionInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.ServerSideEncryptionConfiguration == nil {
		invalidParams.Add(request.NewErrParamRequired("ServerSideEncryptionConfiguration"))
	}
	if p.ServerSideEncryptionConfiguration != nil {
		if err := p.ServerSideEncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ServerSideEncryptionConfiguration", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketEncryptionInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketEncryptionInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.ServerSideEncryptionConfiguration != nil {
		v := p.ServerSideEncryptionConfiguration
		e.SetStruct(helper.PayloadTarget, "ServerSideEncryptionConfiguration", v)
	}

	return nil
}

type PutBucketEncryptionOutput struct{}

// String returns the string representation
func (p PutBucketEncryptionOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketEncryptionOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketEncryption = "PutBucketEncryption"

func (c *Client) PutBucketEncryptionRequest(input *PutBucketEncryptionInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketEncryption,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &PutBucketEncryptionInput{}
	}

	req = c.newRequest(op, input, &PutBucketEncryptionOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketEncryption(input *PutBucketEncryptionInput) (PutBucketEncryptionOutput, error) {
	req := c.PutBucketEncryptionRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketEncryptionOutput{}, err
	}
	out := req.Data.(*PutBucketEncryptionOutput)

	return *out, err
}
