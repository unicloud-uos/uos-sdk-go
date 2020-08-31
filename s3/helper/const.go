package helper

const (
	SDKName    = "uos-sdk-go"
	SDKVersion = "v1.0.1"
)

const (
	ServiceName        = "Unicloud S3" // Service's name
	ServiceID          = "UOS"         // Service's identifier
	APIVersion         = "202020801"   // SDK version
	ServiceNameForSign = "s3"          // Name: Credential=*/s3/
	SDKSigningRegion   = "TJ"          // SDK default signing region
)

type BucketACL string

const (
	BucketACLPrivate           BucketACL = "private"
	BucketACLPublicRead        BucketACL = "public-read"
	BucketACLPublicReadWrite   BucketACL = "public-read-write"
	BucketACLAuthenticatedRead BucketACL = "authenticated-read"
)

// signer
const (
	UOSSigningV4Algorithm  = "UOS4-HMAC-SHA256"     // Value for X-Uos-Algorithm
	UOSAlgorithmKey        = "X-Uos-Algorithm"      // Indicates the signing algorithm
	UOSSecurityTokenKey    = "X-Uos-Security-Token" // Indicates the security token to be used with temporary credentials
	UOSDateKey             = "X-Uos-Date"           // Date is the UTC timestamp, format YYYYMMDD'T'HHMMSS'Z'
	UOSExpiresKey          = "X-Uos-Expires"        // Indicates the interval of time the presigned is valid for
	UOSCredentialKey       = "X-Uos-Credential"     // IS access key ID and credential scope
	UOSSignedHeadersKey    = "X-Uos-SignedHeaders"  // Headers signed for the request
	UOSSignatureKey        = "X-Uos-Signature"      // Query parameter to store the SigV4 signature
	UOSSignedTimeFormat    = "20060102T150405Z"     // Format fot X-Uos-Date
	UOSSignShortTimeFormat = "20060102"             // Shorten time format used in the credential scope
)

// header
const (
	UOSRequestID = "X-Uos-Request-Id"
)
