package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type CompleteMultipartUploadInput struct {
	Bucket          *string
	Key             *string
	MultipartUpload *CompletedMultipartUpload
	UploadId        *string
}

// String returns the string representation
func (a CompleteMultipartUploadInput) String() string {
	return helper.Prettify(a)
}

func (a CompleteMultipartUploadInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CompleteMultipartUploadInput"}

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

func (a CompleteMultipartUploadInput) getBucket() (v string) {
	if a.Bucket == nil {
		return v
	}
	return *a.Bucket
}

func (a CompleteMultipartUploadInput) MarshalForPut(e *request.EncoderForPut) error {
	if a.Bucket != nil {
		v := *a.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if a.Key != nil {
		v := *a.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if a.MultipartUpload != nil {
		v := a.MultipartUpload
		e.SetStruct(helper.PayloadTarget, "CompleteMultipartUpload", v)
	}
	if a.UploadId != nil {
		v := *a.UploadId
		e.SetValue(helper.QueryTarget, "uploadId", request.StringValue(v))
	}

	return nil
}

type CompleteMultipartUploadOutput struct {
	Bucket               *string
	ETag                 *string
	Expiration           *string `location:"header" locationName:"x-uos-expiration"`
	Key                  *string
	Location             *string
	SSEKMSKeyId          *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
	VersionId            *string              `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (c CompleteMultipartUploadOutput) String() string {
	return helper.Prettify(c)
}

func (c CompleteMultipartUploadOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	r.Data = &c
	return nil
}

const opCompleteMultipartUpload = "CompleteMultipartUpload"

func (c *Client) CompleteMultipartUploadRequest(input *CompleteMultipartUploadInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opCompleteMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &CompleteMultipartUploadInput{}
	}

	req = c.newRequest(op, input, &CompleteMultipartUploadOutput{})

	return
}

func (c *Client) CompleteMultipartUpload(input *CompleteMultipartUploadInput) (CompleteMultipartUploadOutput, error) {
	req := c.CompleteMultipartUploadRequest(input)

	err := req.Do()
	if err != nil {
		return CompleteMultipartUploadOutput{}, err
	}
	out := req.Data.(*CompleteMultipartUploadOutput)

	return *out, err
}
