package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type GetBucketLoggingInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketLoggingInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLoggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketLoggingInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketLoggingInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketLoggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketLoggingOutput struct {
	LoggingEnabled *LoggingEnabled
}

// String returns the string representation
func (g GetBucketLoggingOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketLoggingOutput) UnmarshalBody(r *request.Request) error {
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

const opGetBucketLogging = "GetBucketLogging"

func (c *Client) GetBucketLoggingRequest(input *GetBucketLoggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketLogging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?logging",
	}

	if input == nil {
		input = &GetBucketLoggingInput{}
	}

	req = c.newRequest(op, input, &GetBucketLoggingOutput{})
	return
}

func (c *Client) GetBucketLogging(input *GetBucketLoggingInput) (*GetBucketLoggingOutput, error) {
	req := c.GetBucketLoggingRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketLoggingOutput{}, err
	}
	out := req.Data.(*GetBucketLoggingOutput)

	return out, err
}
