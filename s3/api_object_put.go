package s3

import (
	"io"
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type PutObjectInput struct {
	ACL                     *string
	Body                    io.ReadSeeker
	Bucket                  *string
	ContentLength           *int64
	ContentMD5              *string
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
func (p PutObjectInput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if p.Key != nil && len(*p.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutObjectInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if p.Key != nil {
		v := *p.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if p.Body != nil {
		v := p.Body
		e.SetStream(helper.PayloadTarget, "Body", request.ReadSeekerStream{V: v})
	}
	if p.Metadata != nil {
		v := p.Metadata
		for key, value := range v {
			// TODO: maybe need to judge key
			e.SetValue(helper.HeaderTarget, "x-uos-meta-"+key, request.StringValue(value))
		}
	}

	if p.ACL != nil {
		v := *p.ACL
		e.SetValue(helper.HeaderTarget, "x-uos-acl", request.StringValue(v))
	}
	if p.ContentLength != nil {
		v := *p.ContentLength
		e.SetValue(helper.HeaderTarget, "Content-Length", request.Int64Value(v))
	}
	if p.ContentMD5 != nil {
		v := *p.ContentMD5
		e.SetValue(helper.HeaderTarget, "Content-MD5", request.StringValue(v))
	}
	if p.ContentType != nil {
		v := *p.ContentType
		e.SetValue(helper.HeaderTarget, "Content-Type", request.StringValue(v))
	}
	if p.Expires != nil {
		v := *p.Expires
		e.SetValue(helper.HeaderTarget, "Expires", request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if p.ForbidOverwrite != nil {
		v := *p.ForbidOverwrite
		e.SetValue(helper.HeaderTarget, "x-uos-forbid-overwrite", request.BoolValue(v))
	}
	if p.SSECustomerAlgorithm != nil {
		v := *p.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if p.SSECustomerKey != nil {
		v := *p.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if p.SSECustomerKeyMD5 != nil {
		v := *p.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if p.SSEKMSKeyId != nil {
		v := *p.SSEKMSKeyId
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-uos-kms-key-id", request.StringValue(v))
	}
	if len(p.ServerSideEncryption) > 0 {
		v := p.ServerSideEncryption
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption", v)
	}
	if len(p.StorageClass) > 0 {
		v := p.StorageClass
		e.SetValue(helper.HeaderTarget, "x-uos-storage-class", v)
	}
	if p.Tagging != nil {
		v := *p.Tagging
		e.SetValue(helper.HeaderTarget, "x-uos-tagging", request.StringValue(v))
	}
	if p.WebsiteRedirectLocation != nil {
		v := *p.WebsiteRedirectLocation
		e.SetValue(helper.HeaderTarget, "x-uos-website-redirect-location", request.StringValue(v))
	}

	return nil
}

type PutObjectOutput struct {
	ETag                 *string              `location:"header" locationName:"ETag"`
	Expiration           *string              `location:"header" locationName:"x-uos-expiration"`
	SSECustomerAlgorithm *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId          *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id" type:"string"`
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
	VersionId            *string              `location:"header" locationName:"x-uos-version-id"`
}

// String returns the string representation
func (p PutObjectOutput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutObject = "PutObject"

func (c *Client) PutObjectRequest(input *PutObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &PutObjectInput{}
	}

	req = c.newRequest(op, input, &PutObjectOutput{})
	return
}

func (c *Client) PutObject(input *PutObjectInput) (PutObjectOutput, error) {
	req := c.PutObjectRequest(input)

	err := req.Do()
	if err != nil {
		return PutObjectOutput{}, err
	}
	out := req.Data.(*PutObjectOutput)
	c.Logger.Debug("Create bucket output: ", out)

	return *out, err
}
