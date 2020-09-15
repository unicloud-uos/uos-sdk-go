package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketEncryptionInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketEncryptionInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketEncryptionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketEncryptionInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketEncryptionInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketEncryptionInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketEncryptionOutput struct {
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration
}

// String returns the string representation
func (g GetBucketEncryptionOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketEncryptionOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &g.ServerSideEncryptionConfiguration)
	if err != nil {
		return err
	}

	r.Data = &g
	return nil
}

const opGetBucketEncryption = "GetBucketEncryption"

func (c *Client) GetBucketEncryptionRequest(input *GetBucketEncryptionInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketEncryption,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?encryption",
	}

	if input == nil {
		input = &GetBucketEncryptionInput{}
	}

	req = c.newRequest(op, input, &GetBucketEncryptionOutput{})
	return
}

func (c *Client) GetBucketEncryption(input *GetBucketEncryptionInput) (*GetBucketEncryptionOutput, error) {
	req := c.GetBucketEncryptionRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketEncryptionOutput{}, err
	}
	out := req.Data.(*GetBucketEncryptionOutput)

	return out, err
}
