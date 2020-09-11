package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type GetBucketAclInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketAclInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketAclInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketAclInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketAclInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketAclOutput struct {
	AccessControlList []Grant `xml:"AccessControlList>Grant"`
	Owner             *Owner  // Container for the bucket owner's display name and ID.
}

// String returns the string representation
func (g GetBucketAclOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketAclOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketAcl = "GetBucketAcl"

func (c *Client) GetBucketAclRequest(input *GetBucketAclInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?acl",
	}

	if input == nil {
		input = &GetBucketAclInput{}
	}

	req = c.newRequest(op, input, &GetBucketAclOutput{})
	return
}

func (c *Client) GetBucketAcl(input *GetBucketAclInput) (*GetBucketAclOutput, error) {
	req := c.GetBucketAclRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketAclOutput{}, err
	}
	out := req.Data.(*GetBucketAclOutput)

	return out, err
}
