package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketTaggingInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketTaggingInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketTaggingInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketTaggingInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketTaggingOutput struct {
	TagSet []Tag `xml:"TagSet>Tag"`
}

// String returns the string representation
func (g GetBucketTaggingOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketTaggingOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketTagging = "GetBucketTagging"

func (c *Client) GetBucketTaggingRequest(input *GetBucketTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketTagging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &GetBucketTaggingInput{}
	}

	req = c.newRequest(op, input, &GetBucketTaggingOutput{})

	return
}

func (c *Client) GetBucketTagging(input *GetBucketTaggingInput) (*GetBucketTaggingOutput, error) {
	req := c.GetBucketTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketTaggingOutput{}, err
	}
	out := req.Data.(*GetBucketTaggingOutput)

	return out, err
}
