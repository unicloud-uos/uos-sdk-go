package s3

import (
	"io"
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type AppendObjectInput struct {
	ACL                  *string
	Body                 io.ReadSeeker
	Bucket               *string
	ContentDisposition   *string
	ContentEncoding      *string
	ContentLanguage      *string
	ContentLength        *int64
	ContentMD5           *string
	ContentType          *string
	Expires              *time.Time
	Key                  *string
	Metadata             map[string]string
	SSECustomerAlgorithm *string
	SSECustomerKey       *string
	SSECustomerKeyMD5    *string
	SSEKMSKeyId          *string
	ServerSideEncryption ServerSideEncryption
	StorageClass         StorageClass
	Position             *int64
}

// String returns the string representation
func (a AppendObjectInput) String() string {
	return helper.Prettify(a)
}

func (a AppendObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AppendObjectInput"}

	if a.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if a.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if a.Key != nil && len(*a.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (a AppendObjectInput) GetBucketName() (v string) {
	if a.Bucket == nil {
		return v
	}
	return *a.Bucket
}

func (a AppendObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if a.Bucket != nil {
		v := *a.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if a.Key != nil {
		v := *a.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if a.Body != nil {
		v := a.Body
		e.SetStream(helper.PayloadTarget, "Body", request.ReadSeekerStream{V: v})
	}
	if a.Metadata != nil {
		v := a.Metadata
		for key, value := range v {
			// TODO: maybe need to judge key
			e.SetValue(helper.HeaderTarget, "x-uos-meta-"+key, request.StringValue(value))
		}
	}

	if a.ACL != nil {
		v := *a.ACL
		e.SetValue(helper.HeaderTarget, "x-uos-acl", request.StringValue(v))
	}
	if a.ContentDisposition != nil {
		v := *a.ContentDisposition
		e.SetValue(helper.HeaderTarget, "Content-Disposition", request.StringValue(v))
	}
	if a.ContentEncoding != nil {
		v := *a.ContentEncoding
		e.SetValue(helper.HeaderTarget, "Content-Encoding", request.StringValue(v))
	}
	if a.ContentLanguage != nil {
		v := *a.ContentLanguage
		e.SetValue(helper.HeaderTarget, "Content-Language", request.StringValue(v))
	}
	if a.ContentLength != nil {
		v := *a.ContentLength
		e.SetValue(helper.HeaderTarget, "Content-Length", request.Int64Value(v))
	}
	if a.ContentMD5 != nil {
		v := *a.ContentMD5
		e.SetValue(helper.HeaderTarget, "Content-MD5", request.StringValue(v))
	}
	if a.ContentType != nil {
		v := *a.ContentType
		e.SetValue(helper.HeaderTarget, "Content-Type", request.StringValue(v))
	}
	if a.Expires != nil {
		v := *a.Expires
		e.SetValue(helper.HeaderTarget, "Expires", request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if a.SSECustomerAlgorithm != nil {
		v := *a.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if a.SSECustomerKey != nil {
		v := *a.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if a.SSECustomerKeyMD5 != nil {
		v := *a.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if a.SSEKMSKeyId != nil {
		v := *a.SSEKMSKeyId
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-uos-kms-key-id", request.StringValue(v))
	}
	if len(a.ServerSideEncryption) > 0 {
		v := a.ServerSideEncryption
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption", v)
	}
	if len(a.StorageClass) > 0 {
		v := a.StorageClass
		e.SetValue(helper.HeaderTarget, "x-uos-storage-class", v)
	}
	if a.Position != nil {
		v := *a.Position
		e.SetValue(helper.QueryTarget, "position", request.Int64Value(v))
	}

	return nil
}

type AppendObjectOutput struct {
	ETag                 *string `location:"header" locationName:"ETag"`
	Expiration           *string `location:"header" locationName:"x-uos-expiration"`
	SSECustomerAlgorithm *string `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5    *string `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId          *string `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption *string `location:"header" locationName:"x-uos-server-side-encryption"`
	VersionId            *string `location:"header" locationName:"x-uos-version-id"`
	NextPosition         *int64  `location:"header" locationName:"x-uos-next-append-position"`
}

// String returns the string representation
func (p AppendObjectOutput) String() string {
	return helper.Prettify(p)
}

func (p AppendObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opAppendObject = "AppendObject"

func (c *Client) AppendObjectRequest(input *AppendObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opAppendObject,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?append",
	}

	if input == nil {
		input = &AppendObjectInput{}
	}

	req = c.newRequest(op, input, &AppendObjectOutput{})
	return
}

func (c *Client) AppendObject(input *AppendObjectInput) (AppendObjectOutput, error) {
	req := c.AppendObjectRequest(input)

	err := req.Do()
	if err != nil {
		return AppendObjectOutput{}, err
	}
	out := req.Data.(*AppendObjectOutput)

	return *out, err
}
