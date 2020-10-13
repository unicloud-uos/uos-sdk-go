package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteObjectTaggingInput struct {
	Bucket    *string
	Key       *string
	VersionId *string
}

// String returns the string representation
func (d DeleteObjectTaggingInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectTaggingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteObjectTaggingInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if d.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if d.Key != nil && len(*d.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteObjectTaggingInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteObjectTaggingInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if d.Key != nil {
		v := *d.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}

	if d.VersionId != nil {
		v := *d.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type DeleteObjectTaggingOutput struct {
	VersionId *string `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (d DeleteObjectTaggingOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectTaggingOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteObjectTagging = "DeleteObjectTagging"

func (c *Client) DeleteObjectTaggingRequest(input *DeleteObjectTaggingInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteObjectTagging,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &DeleteObjectTaggingInput{}
	}

	req = c.newRequest(op, input, &DeleteObjectTaggingOutput{})

	return
}

func (c *Client) DeleteObjectTagging(input *DeleteObjectTaggingInput) (*DeleteObjectTaggingOutput, error) {
	req := c.DeleteObjectTaggingRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteObjectTaggingOutput{}, err
	}
	out := req.Data.(*DeleteObjectTaggingOutput)

	return out, err
}
