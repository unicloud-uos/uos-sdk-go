package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutObjectMetaInput struct {
	Bucket            *string
	Key               *string
	MetaConfiguration *MetaConfiguration
}

// String returns the string representation
func (p PutObjectMetaInput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectMetaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectMetaInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if p.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if p.Key != nil && len(*p.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if p.MetaConfiguration != nil {
		if err := p.MetaConfiguration.Validate(); err != nil {
			invalidParams.AddNested("MetaConfiguration", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutObjectMetaInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutObjectMetaInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if p.Key != nil {
		v := *p.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}

	if p.MetaConfiguration != nil {
		v := p.MetaConfiguration
		e.SetStruct(helper.PayloadTarget, "MetaConfiguration", v)
	}

	return nil
}

type PutObjectMetaOutput struct{}

// String returns the string representation
func (p PutObjectMetaOutput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectMetaOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutObjectMeta = "PutObjectMeta"

func (c *Client) PutObjectMetaRequest(input *PutObjectMetaInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutObjectMeta,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?meta",
	}

	if input == nil {
		input = &PutObjectMetaInput{}
	}

	req = c.newRequest(op, input, &PutObjectMetaOutput{})
	return
}

func (c *Client) PutObjectMeta(input *PutObjectMetaInput) (PutObjectMetaOutput, error) {
	req := c.PutObjectMetaRequest(input)

	err := req.Do()
	if err != nil {
		return PutObjectMetaOutput{}, err
	}
	out := req.Data.(*PutObjectMetaOutput)

	return *out, err
}
