package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutObjectTaggingInput struct {
	Bucket    *string
	Key       *string
	Tagging   *Tagging
	VersionId *string
}

// String returns the string representation
func (p PutObjectTaggingInput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectTaggingInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if p.Key != nil && len(*p.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
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

func (p PutObjectTaggingInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutObjectTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if p.Key != nil {
		v := *p.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if p.Tagging != nil {
		v := p.Tagging
		e.SetStruct(helper.PayloadTarget, "Tagging", v)
	}
	if p.VersionId != nil {
		v := *p.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}

	return nil
}

type PutObjectTaggingOutput struct {
	VersionId *string `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (p PutObjectTaggingOutput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectTaggingOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutObjectTagging = "PutObjectTagging"

func (c *Client) PutObjectTaggingRequest(input *PutObjectTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutObjectTagging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &PutObjectTaggingInput{}
	}

	req = c.newRequest(op, input, &PutObjectTaggingOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutObjectTagging(input *PutObjectTaggingInput) (PutObjectTaggingOutput, error) {
	req := c.PutObjectTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return PutObjectTaggingOutput{}, err
	}
	out := req.Data.(*PutObjectTaggingOutput)

	return *out, err
}
