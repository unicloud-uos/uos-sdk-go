package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutBucketLifecycleInput struct {
	Bucket                 *string
	LifecycleConfiguration *LifecycleConfiguration
}

// String returns the string representation
func (p PutBucketLifecycleInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketLifecycleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketLifecycleInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.LifecycleConfiguration == nil {
		invalidParams.Add(request.NewErrParamRequired("LifecycleConfiguration"))
	}
	if p.LifecycleConfiguration != nil {
		if err := p.LifecycleConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LifecycleConfiguration", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutBucketLifecycleInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketLifecycleInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}

	if p.LifecycleConfiguration != nil {
		v := p.LifecycleConfiguration
		e.SetStruct(helper.PayloadTarget, "LifecycleConfiguration", v)
	}

	return nil
}

type PutBucketLifecycleOutput struct {}

// String returns the string representation
func (p PutBucketLifecycleOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketLifecycleOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketLifecycle = "PutBucketLifecycle"

func (c *Client) PutBucketLifecycleRequest(input *PutBucketLifecycleInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketLifecycle,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &PutBucketLifecycleInput{}
	}

	req = c.newRequest(op, input, &PutBucketLifecycleOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketLifecycle(input *PutBucketLifecycleInput) (PutBucketCorsOutput, error) {
	req := c.PutBucketLifecycleRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketCorsOutput{}, err
	}
	out := req.Data.(*PutBucketCorsOutput)

	return *out, err
}