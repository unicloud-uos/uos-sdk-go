package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type DeleteObjectInput struct {
	Bucket    *string
	Key       *string
	VersionId *string
}

// String returns the string representation
func (d DeleteObjectInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteObjectInput"}

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

func (d DeleteObjectInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteObjectInput) MarshalForPut(e *request.EncoderForPut) error {
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

type DeleteObjectOutput struct {
	DeleteMarker *bool   `location:"header" locationName:"x-uos-delete-marker"`
	VersionId    *string `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (d DeleteObjectOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteObject = "DeleteObject"

func (c *Client) DeleteObjectRequest(input *DeleteObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteObject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &DeleteObjectInput{}
	}

	req = c.newRequest(op, input, &DeleteObjectOutput{})
	return
}

func (c *Client) DeleteObject(input *DeleteObjectInput) (*DeleteObjectOutput, error) {
	req := c.DeleteObjectRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteObjectOutput{}, err
	}
	out := req.Data.(*DeleteObjectOutput)

	return out, err
}
