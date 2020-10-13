package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type RestoreObjectInput struct {
	Bucket         *string
	Key            *string
	RestoreRequest *RestoreRequest
	VersionId      *string
}

// String returns the string representation
func (r RestoreObjectInput) String() string {
	return helper.Prettify(r)
}

func (r RestoreObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreObjectInput"}

	if r.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if r.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if r.Key != nil && len(*r.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}
	if r.RestoreRequest != nil {
		if err := r.RestoreRequest.Validate(); err != nil {
			invalidParams.AddNested("RestoreRequest", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (r RestoreObjectInput) getBucket() (v string) {
	if r.Bucket == nil {
		return v
	}
	return *r.Bucket
}

func (r RestoreObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if r.Bucket != nil {
		v := *r.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if r.Key != nil {
		v := *r.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if r.RestoreRequest != nil {
		v := r.RestoreRequest
		e.SetStruct(helper.PayloadTarget, "RestoreRequest", v)
	}
	if r.VersionId != nil {
		v := *r.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type RestoreObjectOutput struct {
	RestoreOutputPath *string `location:"header" locationName:"x-uos-restore-output-path"`
}

// String returns the string representation
func (s RestoreObjectOutput) String() string {
	return helper.Prettify(s)
}

func (s RestoreObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opRestoreObject = "RestoreObject"

func (c *Client) RestoreObjectRequest(input *RestoreObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opRestoreObject,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?restore",
	}

	if input == nil {
		input = &RestoreObjectInput{}
	}

	req = c.newRequest(op, input, &RestoreObjectOutput{})
	return
}

func (c *Client) RestoreObject(input *RestoreObjectInput) (RestoreObjectOutput, error) {
	req := c.RestoreObjectRequest(input)

	err := req.Do()
	if err != nil {
		return RestoreObjectOutput{}, err
	}
	out := req.Data.(*RestoreObjectOutput)

	return *out, err
}
