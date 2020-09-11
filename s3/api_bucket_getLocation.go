package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketLocationInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketLocationInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLocationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketLocationInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketLocationInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketLocationInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketLocationOutput struct {
	LocationConstraint *string
}

// String returns the string representation
func (g GetBucketLocationOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLocationOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketLocation = "GetBucketLocation"

func (c *Client) GetBucketLocationRequest(input *GetBucketLocationInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketLocation,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?location",
	}

	if input == nil {
		input = &GetBucketLocationInput{}
	}

	req = c.newRequest(op, input, &GetBucketLocationOutput{})
	return
}

func (c *Client) GetBucketLocation(input *GetBucketLocationInput) (*GetBucketLocationOutput, error) {
	req := c.GetBucketLocationRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketLocationOutput{}, err
	}
	out := req.Data.(*GetBucketLocationOutput)

	return out, err
}
