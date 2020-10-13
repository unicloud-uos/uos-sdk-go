package s3

import (
	"encoding/xml"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetObjectTaggingInput struct {
	Bucket    *string
	Key       *string
	VersionId *string
}

// String returns the string representation
func (g GetObjectTaggingInput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetObjectTaggingInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if g.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if g.Key != nil && len(*g.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetObjectTaggingInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetObjectTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if g.Key != nil {
		v := *g.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}

	if g.VersionId != nil {
		v := *g.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type GetObjectTaggingOutput struct {
	TagSet    []Tag   `xml:"TagSet>Tag"`
	VersionId *string `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (g GetObjectTaggingOutput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectTaggingOutput) UnmarshalBody(r *request.Request) error {
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

const opGetObjectTagging = "GetObjectTagging"

func (c *Client) GetObjectTaggingRequest(input *GetObjectTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetObjectTagging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &GetObjectTaggingInput{}
	}

	req = c.newRequest(op, input, &GetObjectTaggingOutput{})

	return
}

func (c *Client) GetObjectTagging(input *GetObjectTaggingInput) (*GetObjectTaggingOutput, error) {
	req := c.GetObjectTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return &GetObjectTaggingOutput{}, err
	}
	out := req.Data.(*GetObjectTaggingOutput)

	return out, err
}
