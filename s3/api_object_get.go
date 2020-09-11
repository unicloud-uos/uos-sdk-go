package s3

import (
	"io"
	"time"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type GetObjectInput struct {
	Bucket                     *string
	IfMatch                    *string    // Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed).
	IfModifiedSince            *time.Time // Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified).
	IfNoneMatch                *string    // Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified).
	IfUnmodifiedSince          *time.Time // Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified).
	Key                        *string    // Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed).
	PartNumber                 *int64     // Part number of the object being read. This is a positive integer between 1 and 10,000.
	Range                      *string    // Downloads the specified range bytes of an object. For more information about the HTTP Range header
	ResponseCacheControl       *string    // Sets the Cache-Control header of the response.
	ResponseContentDisposition *string    // Sets the Content-Disposition header of the response
	ResponseContentEncoding    *string    // Sets the Content-Encoding header of the response.
	ResponseContentLanguage    *string    // Sets the Content-Language header of the response.
	ResponseContentType        *string    // Sets the Content-Type header of the response.
	ResponseExpires            *time.Time // Sets the Expires header of the response.
	SSECustomerAlgorithm       *string
	SSECustomerKey             *string
	SSECustomerKeyMD5          *string
	VersionId                  *string
}

// String returns the string representation
func (g GetObjectInput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetObjectInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if g.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if g.Key != nil && len(*g.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetObjectInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.IfMatch != nil {
		v := *g.IfMatch
		e.SetValue(helper.HeaderTarget, "If-Match", request.StringValue(v))
	}
	if g.IfModifiedSince != nil {
		v := *g.IfModifiedSince
		e.SetValue(helper.HeaderTarget, "If-Modified-Since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if g.IfNoneMatch != nil {
		v := *g.IfNoneMatch
		e.SetValue(helper.HeaderTarget, "If-None-Match", request.StringValue(v))
	}
	if g.IfUnmodifiedSince != nil {
		v := *g.IfUnmodifiedSince
		e.SetValue(helper.HeaderTarget, "If-Unmodified-Since",
			request.TimeValue{T: v, Format: helper.RFC822TimeFormatName})
	}
	if g.Range != nil {
		v := *g.Range
		e.SetValue(helper.HeaderTarget, "Range", request.StringValue(v))
	}
	if g.SSECustomerAlgorithm != nil {
		v := *g.SSECustomerAlgorithm
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-algorithm", request.StringValue(v))
	}
	if g.SSECustomerKey != nil {
		v := *g.SSECustomerKey
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key", request.StringValue(v))
	}
	if g.SSECustomerKeyMD5 != nil {
		v := *g.SSECustomerKeyMD5
		e.SetValue(helper.HeaderTarget, "x-uos-server-side-encryption-customer-key-MD5", request.StringValue(v))
	}
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if g.Key != nil {
		v := *g.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if g.PartNumber != nil {
		v := *g.PartNumber
		e.SetValue(helper.QueryTarget, "partNumber", request.Int64Value(v))
	}
	if g.ResponseCacheControl != nil {
		v := *g.ResponseCacheControl
		e.SetValue(helper.QueryTarget, "response-cache-control", request.StringValue(v))
	}
	if g.ResponseContentDisposition != nil {
		v := *g.ResponseContentDisposition
		e.SetValue(helper.QueryTarget, "response-content-disposition", request.StringValue(v))
	}
	if g.ResponseContentEncoding != nil {
		v := *g.ResponseContentEncoding
		e.SetValue(helper.QueryTarget, "response-content-encoding", request.StringValue(v))
	}
	if g.ResponseContentLanguage != nil {
		v := *g.ResponseContentLanguage
		e.SetValue(helper.QueryTarget, "response-content-language", request.StringValue(v))
	}
	if g.ResponseContentType != nil {
		v := *g.ResponseContentType
		e.SetValue(helper.QueryTarget, "response-content-type", request.StringValue(v))
	}
	if g.ResponseExpires != nil {
		v := *g.ResponseExpires
		e.SetValue(helper.QueryTarget, "response-expires",
			request.TimeValue{T: v, Format: helper.ISO8601TimeFormatName})
	}
	if g.VersionId != nil {
		v := *g.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type GetObjectOutput struct {
	Body                    io.ReadCloser
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
	ObjectType              *string              `location:"header" locationName:"x-uos-object-type"`
	PartsCount              *int64               `location:"header" locationName:"x-uos-mp-parts-count"`
	Restore                 *string              `location:"header" locationName:"x-uos-restore"`
	SSECustomerAlgorithm    *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-algorithm"`
	SSECustomerKeyMD5       *string              `location:"header" locationName:"x-uos-server-side-encryption-customer-key-MD5"`
	SSEKMSKeyId             *string              `location:"header" locationName:"x-uos-server-side-encryption-uos-kms-key-id"`
	ServerSideEncryption    ServerSideEncryption `location:"header" locationName:"x-uos-server-side-encryption"`
	StorageClass            StorageClass         `location:"header" locationName:"x-uos-storage-class"`
	TagCount                *int64               `location:"header" locationName:"x-uos-tagging-count"`
	VersionId               *string              `location:"header" locationName:"x-uos-version-id"`
	WebsiteRedirectLocation *string              `location:"header" locationName:"x-uos-website-redirect-location"`
}

// String returns the string representation
func (g GetObjectOutput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectOutput) UnmarshalBody(r *request.Request) error {
	g.Body = r.HTTPResponse.Body
	r.Data = &g

	return nil
}

const opGetObject = "GetObject"

func (c *Client) GetObjectRequest(input *GetObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &GetObjectInput{}
	}

	req = c.newRequest(op, input, &GetObjectOutput{})
	return
}

func (c *Client) GetObject(input *GetObjectInput) (*GetObjectOutput, error) {
	req := c.GetObjectRequest(input)

	err := req.Do()
	if err != nil {
		return &GetObjectOutput{}, err
	}
	out := req.Data.(*GetObjectOutput)

	return out, err
}
