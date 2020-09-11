package s3

import (
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type HeadObjectInput struct {
	Bucket               *string
	IfMatch              *string    // Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed).
	IfModifiedSince      *time.Time // Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified).
	IfNoneMatch          *string    // Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified).
	IfUnmodifiedSince    *time.Time // Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified).
	Key                  *string    // Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed).
	PartNumber           *int64     // Part number of the object being read. This is a positive integer between 1 and 10,000.
	Range                *string    // Downloads the specified range bytes of an object. For more information about the HTTP Range header
	SSECustomerAlgorithm *string
	SSECustomerKey       *string
	SSECustomerKeyMD5    *string
	VersionId            *string
}

// String returns the string representation
func (h HeadObjectInput) String() string {
	return helper.Prettify(h)
}

func (h HeadObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HeadObjectInput"}

	if h.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if h.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if h.Key != nil && len(*h.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (h HeadObjectInput) getBucket() (v string) {
	if h.Bucket == nil {
		return v
	}
	return *h.Bucket
}

func (h HeadObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if h.IfMatch != nil {
		v := *h.IfMatch
		e.SetValue(helper.HeaderTarget, "If-Match", request.StringValue(v))
	}
	if h.IfModifiedSince != nil {
		v := *h.IfModifiedSince
		e.SetValue(helper.HeaderTarget, "If-Modified-Since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if h.IfNoneMatch != nil {
		v := *h.IfNoneMatch
		e.SetValue(helper.HeaderTarget, "If-None-Match", request.StringValue(v))
	}
	if h.IfUnmodifiedSince != nil {
		v := *h.IfUnmodifiedSince
		e.SetValue(helper.HeaderTarget, "If-Unmodified-Since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if h.Range != nil {
		v := *h.Range
		e.SetValue(helper.HeaderTarget, "Range", request.StringValue(v))
	}
	if h.SSECustomerAlgorithm != nil {
		v := *h.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if h.SSECustomerKey != nil {
		v := *h.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if h.SSECustomerKeyMD5 != nil {
		v := *h.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if h.Bucket != nil {
		v := *h.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if h.Key != nil {
		v := *h.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if h.PartNumber != nil {
		v := *h.PartNumber
		e.SetValue(helper.QueryTarget, "partNumber", request.Int64Value(v))
	}

	if h.VersionId != nil {
		v := *h.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type HeadObjectOutput struct {
	AcceptRanges            *string              `location:"header" locationName:"accept-ranges"`
	CacheControl            *string              `location:"header" locationName:"Cache-Control"`
	ContentDisposition      *string              `location:"header" locationName:"Content-Disposition"`
	ContentEncoding         *string              `location:"header" locationName:"Content-Encoding"`
	ContentLanguage         *string              `location:"header" locationName:"Content-Language"`
	ContentLength           *int64               `location:"header" locationName:"Content-Length"`
	ContentRange            *string              `location:"header" locationName:"Content-Range"`
	ContentType             *string              `location:"header" locationName:"Content-Type"`
	DeleteMarker            *bool                `location:"header" locationName:"x-uos-delete-marker"`
	ETag                    *string              `location:"header" locationName:"ETag"`
	Expiration              *string              `location:"header" locationName:"x-uos-expiration"`
	Expires                 *string              `location:"header" locationName:"Expires"`
	LastModified            *time.Time           `location:"header" locationName:"Last-Modified"`
	Metadata                map[string]string    `location:"header" locationName:"x-uos-meta-"`
	PartsCount              *int64               `location:"header" locationName:"x-uos-mp-parts-count"`
	Restore                 *string              `location:"header" locationName:"x-uos-restore"`
	SSECustomerAlgorithm    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId             *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption    ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
	StorageClass            StorageClass         `location:"header" locationName:"x-uos-storage-class"`
	VersionId               *string              `location:"header" locationName:"x-uos-version-id"`
	WebsiteRedirectLocation *string              `location:"header" locationName:"x-uos-website-redirect-location"`
}

// String returns the string representation
func (h HeadObjectOutput) String() string {
	return helper.Prettify(h)
}

func (h HeadObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opHeadObject = "HeadObject"

func (c *Client) HeadObjectRequest(input *HeadObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opHeadObject,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &HeadObjectInput{}
	}

	req = c.newRequest(op, input, &HeadObjectOutput{})
	return
}

func (c *Client) HeadObject(input *HeadObjectInput) (*HeadObjectOutput, error) {
	req := c.HeadObjectRequest(input)

	err := req.Do()
	if err != nil {
		return &HeadObjectOutput{}, err
	}
	out := req.Data.(*HeadObjectOutput)

	return out, err
}
