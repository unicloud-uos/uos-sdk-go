package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketPolicyInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketPolicyInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketPolicyInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketPolicyInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketPolicyInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketPolicyOutput struct {
	Policy *string
}

// String returns the string representation
func (g GetBucketPolicyOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketPolicyOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketPolicy = "GetBucketPolicy"

func (c *Client) GetBucketPolicyRequest(input *GetBucketPolicyInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?policy",
	}

	if input == nil {
		input = &GetBucketPolicyInput{}
	}

	req = c.newRequest(op, input, &GetBucketPolicyOutput{})
	return
}

func (c *Client) GetBucketPolicy(input *GetBucketPolicyInput) (*GetBucketPolicyOutput, error) {
	req := c.GetBucketPolicyRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketPolicyOutput{}, err
	}
	out := req.Data.(*GetBucketPolicyOutput)

	return out, err
}
