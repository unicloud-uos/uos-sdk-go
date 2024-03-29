package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/aws/request"
	"io"
	"time"
)

const opAppendObject = "AppendObject"

type AppendObjectInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The canned ACL to apply to the object.
	ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"ObjectCannedACL"`

	// Object data.
	Body io.ReadSeeker `type:"blob"`

	// Name of the bucket to which the PUT operation was initiated.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes. This parameter is useful when the size of the
	// body cannot be determined automatically.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The base64-encoded 128-bit MD5 digest of the part data.
	ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	//// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	//GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`
	//
	//// Allows grantee to read the object data and its metadata.
	//GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`
	//
	//// Allows grantee to read the object ACL.
	//GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`
	//
	//// Allows grantee to write the ACL for the applicable object.
	//GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Object key for which the PUT operation was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	//// The Legal Hold status that you want to apply to the specified object.
	//ObjectLockLegalHoldStatus *string `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"ObjectLockLegalHoldStatus"`
	//
	//// The Object Lock mode that you want to apply to this object.
	//ObjectLockMode *string `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"ObjectLockMode"`
	//
	//// The date and time when you want this object's Object Lock to expire.
	//ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	//RequestPayer *string `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"RequestPayer"`
	//
	//// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	//SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	//
	//// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	//// data. This value is used to store the object and then it is discarded; Amazon
	//// does not store the encryption key. The key must be appropriate for use with
	//// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	//// header.
	//SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`
	//
	//// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	//// Amazon S3 uses this header for a message integrity check to ensure the encryption
	//// key was transmitted without error.
	//SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. Documentation on configuring any of the officially supported
	// AWS SDKs and CLI can be found at http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
	//SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`

	//// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	//// (For example, "Key1=Value1")
	//Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`
	//
	//// If the bucket is configured as a website, redirects requests for this object
	//// to another object in the same bucket or to an external URL. Amazon S3 stores
	//// the value of this header in the object metadata.
	//WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`

	Position *int64 `location:"querystring" locationName:"position" type:"int64" required:"true"`
}

type AppendObjectOutput struct {
	_ struct{} `type:"structure"`

	// Entity tag for the uploaded object.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	//// If present, indicates that the requester was successfully charged for the
	//// request.
	//RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`

	// If server-side encryption with a customer-provided encryption key was requested,
	//// the response will include this header confirming the encryption algorithm
	//// used.
	//SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	//
	//// If server-side encryption with a customer-provided encryption key was requested,
	//// the response will include this header to provide round trip message integrity
	//// verification of the customer-provided encryption key.
	//SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
	//
	//// If present, specifies the ID of the AWS Key Management Service (KMS) master
	//// encryption key that was used for the object.
	//SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	NextPosition *int64 `location:"header" locationName:"x-amz-next-append-position" type:"int64"`
}

func (c *S3) AppendObjectRequest(input *AppendObjectInput) (req *request.Request, output *AppendObjectOutput) {
	op := &request.Operation{
		Name:       opAppendObject,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?append",
	}

	if input == nil {
		input = &AppendObjectInput{}
	}

	output = &AppendObjectOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *S3) AppendObject(input *AppendObjectInput) (*AppendObjectOutput, error) {
	req, out := c.AppendObjectRequest(input)
	return out, req.Send()
}
