package s3

import (
	"encoding/xml"
	"io/ioutil"
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type UploadPartCopyInput struct {
	Bucket                         *string
	CopySource                     *string
	CopySourceIfMatch              *string
	CopySourceIfModifiedSince      *time.Time
	CopySourceIfNoneMatch          *string
	CopySourceIfUnmodifiedSince    *time.Time
	CopySourceRange                *string
	CopySourceSSECustomerAlgorithm *string
	CopySourceSSECustomerKey       *string
	CopySourceSSECustomerKeyMD5    *string
	Key                            *string
	PartNumber                     *int64
	SSECustomerAlgorithm           *string
	SSECustomerKey                 *string
	SSECustomerKeyMD5              *string
	UploadId                       *string
}

// String returns the string representation
func (u UploadPartCopyInput) String() string {
	return helper.Prettify(u)
}

func (u UploadPartCopyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UploadPartCopyInput"}

	if u.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if u.CopySource == nil {
		invalidParams.Add(request.NewErrParamRequired("CopySource"))
	}
	if u.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if u.Key != nil && len(*u.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}
	if u.PartNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PartNumber"))
	}
	if u.UploadId == nil {
		invalidParams.Add(request.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (u UploadPartCopyInput) getBucket() (v string) {
	if u.Bucket == nil {
		return v
	}
	return *u.Bucket
}

func (u UploadPartCopyInput) MarshalForPut(e *request.EncoderForPut) error {
	if u.Bucket != nil {
		v := *u.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if u.Key != nil {
		v := *u.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}

	if u.PartNumber != nil {
		v := *u.PartNumber
		e.SetValue(helper.QueryTarget, "partNumber", request.Int64Value(v))
	}
	if u.UploadId != nil {
		v := *u.UploadId
		e.SetValue(helper.QueryTarget, "uploadId", request.StringValue(v))
	}

	if u.CopySource != nil {
		v := *u.CopySource
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source", request.StringValue(v))
	}

	if u.CopySourceIfMatch != nil {
		v := *u.CopySourceIfMatch
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-match", request.StringValue(v))
	}
	if u.CopySourceIfModifiedSince != nil {
		v := *u.CopySourceIfModifiedSince
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-modified-since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if u.CopySourceIfNoneMatch != nil {
		v := *u.CopySourceIfNoneMatch
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-none-match", request.StringValue(v))
	}
	if u.CopySourceIfUnmodifiedSince != nil {
		v := *u.CopySourceIfUnmodifiedSince
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-unmodified-since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if u.CopySourceRange != nil {
		v := *u.CopySourceRange
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-range", request.StringValue(v))
	}
	if u.CopySourceSSECustomerAlgorithm != nil {
		v := *u.CopySourceSSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if u.CopySourceSSECustomerKey != nil {
		v := *u.CopySourceSSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-key", request.StringValue(v))
	}
	if u.CopySourceSSECustomerKeyMD5 != nil {
		v := *u.CopySourceSSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-key-MD5", request.StringValue(v))
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

type UploadPartCopyOutput struct {
	CopyPartResult       *CopyPartResult
	CopySourceVersionId  *string              `location:"header" locationName:"x-uos-copy-source-version-id"`
	SSECustomerAlgorithm *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId          *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
}

// String returns the string representation
func (p UploadPartCopyOutput) String() string {
	return helper.Prettify(p)
}

func (p UploadPartCopyOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &p)
	if err != nil {
		return err
	}

	r.Data = &p
	return nil
}

const opUploadPartCopy = "UploadPartCopy"

func (c *Client) UploadPartCopyRequest(input *UploadPartCopyInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opUploadPartCopy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &UploadPartCopyInput{}
	}

	req = c.newRequest(op, input, &UploadPartCopyOutput{})

	return
}

func (c *Client) UploadPartCopy(input *UploadPartCopyInput) (UploadPartCopyOutput, error) {
	req := c.UploadPartCopyRequest(input)

	err := req.Do()
	if err != nil {
		return UploadPartCopyOutput{}, err
	}
	out := req.Data.(*UploadPartCopyOutput)

	return *out, err
}
