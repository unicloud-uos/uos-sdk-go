package s3

import (
	"encoding/xml"
	"io/ioutil"
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type CopyObjectInput struct {
	ACL                            *string
	Bucket                         *string
	ContentDisposition             *string
	ContentEncoding                *string
	ContentLanguage                *string
	ContentType                    *string
	CopySource                     *string
	CopySourceIfMatch              *string
	CopySourceIfModifiedSince      *time.Time
	CopySourceIfNoneMatch          *string
	CopySourceIfUnmodifiedSince    *time.Time
	CopySourceSSECustomerAlgorithm *string
	CopySourceSSECustomerKey       *string
	CopySourceSSECustomerKeyMD5    *string
	Expires                        *time.Time
	ForbidOverwrite                *bool
	Key                            *string
	Metadata                       map[string]string
	MetadataDirective              MetadataDirective
	SSECustomerAlgorithm           *string
	SSECustomerKey                 *string
	SSECustomerKeyMD5              *string
	SSEKMSEncryptionContext        *string
	SSEKMSKeyId                    *string
	ServerSideEncryption           ServerSideEncryption
	StorageClass                   StorageClass
	Tagging                        *string
	WebsiteRedirectLocation        *string
}

// String returns the string representation
func (c CopyObjectInput) String() string {
	return helper.Prettify(c)
}

func (c CopyObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CopyObjectInput"}

	if c.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if c.CopySource == nil {
		invalidParams.Add(request.NewErrParamRequired("CopySource"))
	}

	if c.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if c.Key != nil && len(*c.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (c CopyObjectInput) GetBucketName() (v string) {
	if c.Bucket == nil {
		return v
	}
	return *c.Bucket
}

func (c CopyObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if c.Bucket != nil {
		v := *c.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if c.Key != nil {
		v := *c.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if c.Metadata != nil {
		v := c.Metadata
		for key, value := range v {
			// TODO: maybe need to judge key
			e.SetValue(helper.HeaderTarget, "x-uos-meta-"+key, request.StringValue(value))
		}
	}

	if c.ACL != nil {
		v := *c.ACL
		e.SetValue(helper.HeaderTarget, "x-uos-acl", request.StringValue(v))
	}
	if c.ContentDisposition != nil {
		v := *c.ContentDisposition
		e.SetValue(helper.HeaderTarget, "Content-Disposition", request.StringValue(v))
	}
	if c.ContentEncoding != nil {
		v := *c.ContentEncoding
		e.SetValue(helper.HeaderTarget, "Content-Encoding", request.StringValue(v))
	}
	if c.ContentLanguage != nil {
		v := *c.ContentLanguage
		e.SetValue(helper.HeaderTarget, "Content-Language", request.StringValue(v))
	}
	if c.ContentType != nil {
		v := *c.ContentType
		e.SetValue(helper.HeaderTarget, "Content-Type", request.StringValue(v))
	}
	if c.CopySource != nil {
		v := *c.CopySource
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source", request.StringValue(v))
	}
	if c.CopySourceIfMatch != nil {
		v := *c.CopySourceIfMatch
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-match", request.StringValue(v))
	}
	if c.CopySourceIfModifiedSince != nil {
		v := *c.CopySourceIfModifiedSince
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-modified-since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if c.CopySourceIfNoneMatch != nil {
		v := *c.CopySourceIfNoneMatch
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-none-match", request.StringValue(v))
	}
	if c.CopySourceIfUnmodifiedSince != nil {
		v := *c.CopySourceIfUnmodifiedSince
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-if-unmodified-since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if c.CopySourceSSECustomerAlgorithm != nil {
		v := *c.CopySourceSSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if c.CopySourceSSECustomerKey != nil {
		v := *c.CopySourceSSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-key", request.StringValue(v))
	}
	if c.CopySourceSSECustomerKeyMD5 != nil {
		v := *c.CopySourceSSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-copy-source-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if c.Expires != nil {
		v := *c.Expires
		e.SetValue(helper.HeaderTarget, "Expires", request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if c.ForbidOverwrite != nil {
		v := *c.ForbidOverwrite
		e.SetValue(helper.HeaderTarget, "x-uos-forbid-overwrite", request.BoolValue(v))
	}
	if len(c.MetadataDirective) > 0 {
		v := c.MetadataDirective
		e.SetValue(helper.HeaderTarget, "x-uos-metadata-directive", v)
	}
	if c.SSECustomerAlgorithm != nil {
		v := *c.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if c.SSECustomerKey != nil {
		v := *c.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if c.SSECustomerKeyMD5 != nil {
		v := *c.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if c.SSEKMSKeyId != nil {
		v := *c.SSEKMSKeyId
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-uos-kms-key-id", request.StringValue(v))
	}
	if len(c.ServerSideEncryption) > 0 {
		v := c.ServerSideEncryption
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption", v)
	}
	if len(c.StorageClass) > 0 {
		v := c.StorageClass
		e.SetValue(helper.HeaderTarget, "x-uos-storage-class", v)
	}
	if c.Tagging != nil {
		v := *c.Tagging
		e.SetValue(helper.HeaderTarget, "x-uos-tagging", request.StringValue(v))
	}
	if c.WebsiteRedirectLocation != nil {
		v := *c.WebsiteRedirectLocation
		e.SetValue(helper.HeaderTarget, "x-uos-website-redirect-location", request.StringValue(v))
	}

	return nil
}

type CopyObjectOutput struct {
	CopyObjectResult        *CopyObjectResult
	CopySourceVersionId     *string              `location:"header" locationName:"x-uos-copy-source-version-id"`
	Expiration              *string              `location:"header" locationName:"x-uos-expiration"`
	SSECustomerAlgorithm    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSEncryptionContext *string              `location:"header" locationName:"x-uos-server-side-encryption-context"`
	SSEKMSKeyId             *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption    ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
	VersionId               *string              `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (c CopyObjectOutput) String() string {
	return helper.Prettify(c)
}

func (c CopyObjectOutput) UnmarshalBody(r *request.Request) error {
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

const opCopyObject = "CopyObject"

func (c *Client) CopyObjectRequest(input *CopyObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opCopyObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &CopyObjectInput{}
	}

	req = c.newRequest(op, input, &CopyObjectOutput{})
	return
}

func (c *Client) CopyObject(input *CopyObjectInput) (CopyObjectOutput, error) {
	req := c.CopyObjectRequest(input)

	err := req.Do()
	if err != nil {
		return CopyObjectOutput{}, err
	}
	out := req.Data.(*CopyObjectOutput)

	return *out, err
}
