package s3

import (
	"encoding/xml"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
	"io/ioutil"
	"time"
)

type CreateMultipartUploadInput struct {
	ACL                     *string
	Bucket                  *string
	ContentDisposition      *string
	ContentEncoding         *string
	ContentLanguage         *string
	ContentType             *string
	Expires                 *time.Time
	ForbidOverwrite         *bool
	Key                     *string
	Metadata                map[string]string
	SSECustomerAlgorithm    *string
	SSECustomerKey          *string
	SSECustomerKeyMD5       *string
	SSEKMSKeyId             *string
	ServerSideEncryption    ServerSideEncryption
	StorageClass            StorageClass
	Tagging                 *string
	WebsiteRedirectLocation *string
}

// String returns the string representation
func (c CreateMultipartUploadInput) String() string {
	return helper.Prettify(c)
}

func (c CreateMultipartUploadInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBucketInput"}

	if c.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
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

func (c CreateMultipartUploadInput) GetBucketName() (v string) {
	if c.Bucket == nil {
		return v
	}
	return *c.Bucket
}

func (c CreateMultipartUploadInput) MarshalForPut(e *request.EncoderForPut) error {
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
	if c.Expires != nil {
		v := *c.Expires
		e.SetValue(helper.HeaderTarget, "Expires", request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if c.ForbidOverwrite != nil {
		v := *c.ForbidOverwrite
		e.SetValue(helper.HeaderTarget, "x-uos-forbid-overwrite", request.BoolValue(v))
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

type CreateMultipartUploadOutput struct {
	Bucket                  *string
	Key                     *string
	SSECustomerAlgorithm    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSEncryptionContext *string              `location:"header" locationName:"x-uos-server-side-encryption-context"`
	SSEKMSKeyId             *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id" `
	ServerSideEncryption    ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption" `
	UploadId                *string
}

// String returns the string representation
func (c CreateMultipartUploadOutput) String() string {
	return helper.Prettify(c)
}

func (c CreateMultipartUploadOutput) UnmarshalBody(r *request.Request) error {
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

const opCreateMultipartUpload = "CreateMultipartUpload"

func (c *Client) CreateMultipartUploadRequest(input *CreateMultipartUploadInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opCreateMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?uploads",
	}

	if input == nil {
		input = &CreateMultipartUploadInput{}
	}

	req = c.newRequest(op, input, &CreateMultipartUploadOutput{})
	return
}

func (c *Client) CreateMultipartUpload(input *CreateMultipartUploadInput) (CreateMultipartUploadOutput, error) {
	req := c.CreateMultipartUploadRequest(input)

	err := req.Do()
	if err != nil {
		return CreateMultipartUploadOutput{}, err
	}
	out := req.Data.(*CreateMultipartUploadOutput)

	return *out, err
}
