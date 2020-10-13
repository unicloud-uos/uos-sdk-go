package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutBucketVersioningInput struct {
	Bucket                  *string
	VersioningConfiguration *VersioningConfiguration
}

// String returns the string representation
func (p PutBucketVersioningInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketVersioningInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketVersioningInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.VersioningConfiguration == nil {
		invalidParams.Add(request.NewErrParamRequired("VersioningConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketVersioningInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketVersioningInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.VersioningConfiguration != nil {
		v := *p.VersioningConfiguration
		e.SetStruct(helper.PayloadTarget, "VersioningConfiguration", v)
	}

	return nil
}

type PutBucketVersioningOutput struct{}

// String returns the string representation
func (p PutBucketVersioningOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketVersioningOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketVersioning = "PutBucketVersioning"

func (c *Client) PutBucketVersioningRequest(input *PutBucketVersioningInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketVersioning,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?versioning",
	}

	if input == nil {
		input = &PutBucketVersioningInput{}
	}

	req = c.newRequest(op, input, &PutBucketVersioningOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketVersioning(input *PutBucketVersioningInput) (PutBucketVersioningOutput, error) {
	req := c.PutBucketVersioningRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketVersioningOutput{}, err
	}
	out := req.Data.(*PutBucketVersioningOutput)

	return *out, err
}
