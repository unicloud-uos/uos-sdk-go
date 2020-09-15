package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketLifecycleInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketLifecycleInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLifecycleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketLifecycleInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketLifecycleInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketLifecycleInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketLifecycleOutput struct {
	Rules []LifecycleRule `xml:"Rule"`
}

// String returns the string representation
func (g GetBucketLifecycleOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLifecycleOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &g)
	if err != nil {
		return err
	}

	r.Data = &g
	return nil
}

const opGetBucketLifecycle = "GetBucketLifecycle"

func (c *Client) GetBucketLifecycleRequest(input *GetBucketLifecycleInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketLifecycle,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &GetBucketLifecycleInput{}
	}

	req = c.newRequest(op, input, &GetBucketLifecycleOutput{})
	return
}

func (c *Client) GetBucketLifecycle(input *GetBucketLifecycleInput) (*GetBucketLifecycleOutput, error) {
	req := c.GetBucketLifecycleRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketLifecycleOutput{}, err
	}
	out := req.Data.(*GetBucketLifecycleOutput)

	return out, err
}
