package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type AbortMultipartUploadInput struct {
	Bucket   *string
	Key      *string
	UploadId *string
}

// String returns the string representation
func (a AbortMultipartUploadInput) String() string {
	return helper.Prettify(a)
}

func (a AbortMultipartUploadInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AbortMultipartUploadInput"}

	if a.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if a.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if a.Key != nil && len(*a.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if a.UploadId == nil {
		invalidParams.Add(request.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (a AbortMultipartUploadInput) getBucket() (v string) {
	if a.Bucket == nil {
		return v
	}
	return *a.Bucket
}

func (a AbortMultipartUploadInput) MarshalForPut(e *request.EncoderForPut) error {
	if a.Bucket != nil {
		v := *a.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if a.Key != nil {
		v := *a.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if a.UploadId != nil {
		v := *a.UploadId
		e.SetValue(helper.QueryTarget, "uploadId", request.StringValue(v))
	}

	return nil
}

type AbortMultipartUploadOutput struct{}

// String returns the string representation
func (a AbortMultipartUploadOutput) String() string {
	return helper.Prettify(a)
}

func (a AbortMultipartUploadOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opAbortMultipartUpload = "AbortMultipartUpload"

func (c *Client) AbortMultipartUploadRequest(input *AbortMultipartUploadInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opAbortMultipartUpload,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &AbortMultipartUploadInput{}
	}

	req = c.newRequest(op, input, &AbortMultipartUploadOutput{})

	return
}

func (c *Client) AbortMultipartUpload(input *AbortMultipartUploadInput) (AbortMultipartUploadOutput, error) {
	req := c.AbortMultipartUploadRequest(input)

	err := req.Do()
	if err != nil {
		return AbortMultipartUploadOutput{}, err
	}
	out := req.Data.(*AbortMultipartUploadOutput)

	return *out, err
}
