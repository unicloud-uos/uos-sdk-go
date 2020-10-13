package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutBucketWebsiteInput struct {
	Bucket               *string
	WebsiteConfiguration *WebsiteConfiguration
}

// String returns the string representation
func (p PutBucketWebsiteInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketWebsiteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketWebsiteInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.WebsiteConfiguration == nil {
		invalidParams.Add(request.NewErrParamRequired("WebsiteConfiguration"))
	}
	if p.WebsiteConfiguration != nil {
		if err := p.WebsiteConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WebsiteConfiguration", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketWebsiteInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketWebsiteInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.WebsiteConfiguration != nil {
		v := p.WebsiteConfiguration
		e.SetStruct(helper.PayloadTarget, "WebsiteConfiguration", v)
	}

	return nil
}

type PutBucketWebsiteOutput struct{}

// String returns the string representation
func (p PutBucketWebsiteOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketWebsiteOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketWebsite = "PutBucketWebsite"

func (c *Client) PutBucketWebsiteRequest(input *PutBucketWebsiteInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketWebsite,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?website",
	}

	if input == nil {
		input = &PutBucketWebsiteInput{}
	}

	req = c.newRequest(op, input, &PutBucketWebsiteOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketWebsite(input *PutBucketWebsiteInput) (PutBucketWebsiteOutput, error) {
	req := c.PutBucketWebsiteRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketWebsiteOutput{}, err
	}
	out := req.Data.(*PutBucketWebsiteOutput)

	return *out, err
}
