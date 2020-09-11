package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type GetBucketCorsInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketCorsInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketCorsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketCorsInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketCorsInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketCorsInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketCorsOutput struct {
	CORSRules []CORSRule `xml:"CORSRules>CORSRule"`
}

// String returns the string representation
func (g GetBucketCorsOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketCorsOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketCors = "GetBucketCors"

func (c *Client) GetBucketCorsRequest(input *GetBucketCorsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketCors,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &GetBucketCorsInput{}
	}

	req = c.newRequest(op, input, &GetBucketCorsOutput{})

	return
}

func (c *Client) GetBucketCors(input *GetBucketCorsInput) (*GetBucketCorsOutput, error) {
	req := c.GetBucketCorsRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketCorsOutput{}, err
	}
	out := req.Data.(*GetBucketCorsOutput)

	return out, err
}
