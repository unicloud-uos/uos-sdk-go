package s3

import (
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type GetBucketVersioningInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketVersioningInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketVersioningInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketVersioningInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketVersioningInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketVersioningInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketVersioningOutput struct {
	Status *string
}

// String returns the string representation
func (g GetBucketVersioningOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketVersioningOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	str := string(buf)
	g.Status = &str

	r.Data = &g
	return nil
}

const opGetBucketVersioning = "GetBucketVersioning"

func (c *Client) GetBucketVersioningRequest(input *GetBucketVersioningInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketVersioning,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?versioning",
	}

	if input == nil {
		input = &GetBucketVersioningInput{}
	}

	req = c.newRequest(op, input, &GetBucketVersioningOutput{})

	return
}

func (c *Client) GetBucketVersioning(input *GetBucketVersioningInput) (*GetBucketVersioningOutput, error) {
	req := c.GetBucketVersioningRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketVersioningOutput{}, err
	}
	out := req.Data.(*GetBucketVersioningOutput)

	return out, err
}
