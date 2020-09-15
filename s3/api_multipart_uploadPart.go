package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io"
)

type UploadPartInput struct {
	Body                 io.ReadSeeker
	Bucket               *string
	ContentLength        *int64
	ContentMD5           *string
	Key                  *string
	PartNumber           *int64
	SSECustomerAlgorithm *string
	SSECustomerKey       *string
	SSECustomerKeyMD5    *string
	UploadId             *string
}

// String returns the string representation
func (u UploadPartInput) String() string {
	return helper.Prettify(u)
}

func (u UploadPartInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UploadPartInput"}

	if u.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (u UploadPartInput) getBucket() (v string) {
	if u.Bucket == nil {
		return v
	}
	return *u.Bucket
}

func (u UploadPartInput) MarshalForPut(e *request.EncoderForPut) error {
	if u.Bucket != nil {
		v := *u.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if u.Key != nil {
		v := *u.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if u.Body != nil {
		v := u.Body
		e.SetStream(helper.PayloadTarget, "Body", request.ReadSeekerStream{V: v})
	}
	if u.PartNumber != nil {
		v := *u.PartNumber
		e.SetValue(helper.QueryTarget, "partNumber", request.Int64Value(v))
	}
	if u.UploadId != nil {
		v := *u.UploadId
		e.SetValue(helper.QueryTarget, "uploadId", request.StringValue(v))
	}

	if u.ContentLength != nil {
		v := *u.ContentLength
		e.SetValue(helper.HeaderTarget, "Content-Length", request.Int64Value(v))
	}
	if u.ContentMD5 != nil {
		v := *u.ContentMD5
		e.SetValue(helper.HeaderTarget, "Content-MD5", request.StringValue(v))
	}
	if u.SSECustomerAlgorithm != nil {
		v := *u.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if u.SSECustomerKey != nil {
		v := *u.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if u.SSECustomerKeyMD5 != nil {
		v := *u.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}

	return nil
}

type UploadPartOutput struct {
	ETag                 *string              `location:"header" locationName:"ETag"`
	SSECustomerAlgorithm *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm" `
	SSECustomerKeyMD5    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId          *string              `location:"header" locationName:"x-uos-server-side-uos-kms-key-id"`
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
}

// String returns the string representation
func (p UploadPartOutput) String() string {
	return helper.Prettify(p)
}

func (p UploadPartOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opUploadPart = "UploadPart"

func (c *Client) UploadPartRequest(input *UploadPartInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opUploadPart,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &UploadPartInput{}
	}

	req = c.newRequest(op, input, &UploadPartOutput{})

	return
}

func (c *Client) UploadPart(input *UploadPartInput) (UploadPartOutput, error) {
	req := c.UploadPartRequest(input)

	err := req.Do()
	if err != nil {
		return UploadPartOutput{}, err
	}
	out := req.Data.(*UploadPartOutput)

	return *out, err
}
