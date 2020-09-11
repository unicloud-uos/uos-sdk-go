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

// signer
const (
	UOSSigningV4Algorithm  = "UOS4-HMAC-SHA256"                                                 // Value for X-Uos-Algorithm
	EmptyStringSHA256      = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` // The hex encoded sha256 value of an empty string
	UnsignedPayload        = "UNSIGNED-PAYLOAD"                                                 // Indicates that the request payload body is unsigned
	UOSAlgorithmKey        = "X-Uos-Algorithm"                                                  // Indicates the signing algorithm
	UOSSecurityTokenKey    = "X-Uos-Security-Token"                                             // Indicates the security token to be used with temporary credentials
	UOSContentSha256       = "X-Uos-Content-Sha256"                                             // Indicates the Content type
	UOSDateKey             = "X-Uos-Date"                                                       // Date is the UTC timestamp, format YYYYMMDD'T'HHMMSS'Z'
	UOSExpiresKey          = "X-Uos-Expires"                                                    // Indicates the interval of time the presigned is valid for
	UOSCredentialKey       = "X-Uos-Credential"                                                 // IS access key ID and credential scope
	UOSSignedHeadersKey    = "X-Uos-SignedHeaders"                                              // Headers signed for the request
	UOSSignatureKey        = "X-Uos-Signature"                                                  // Query parameter to store the SigV4 signature
	UOSSignedTimeFormat    = "20060102T150405Z"                                                 // Format fot X-Uos-Date
	UOSSignShortTimeFormat = "20060102"                                                         // Shorten time format used in the credential scope
)

// header
const (
	ContentMD5Header = "Content-Md5"
	UOSRequestID     = "X-Uos-Request-Id"
)

// marshal target
const (
	PathTarget       = "Path"
	QueryTarget      = "Query"
	HeaderTarget     = "Header"
	StatusCodeTarget = "StatusCode"
	BodyTarget       = "Body"
	PayloadTarget    = "Payload"
)

// Names of time formats supported by the SDK
const (
	RFC822TimeFormatName  = "rfc822"
	ISO8601TimeFormatName = "iso8601"
	UnixTimeFormatName    = "unixTimestamp"
)

// Time formats supported by the SDK
const (
	RFC822TimeFormat        = "Mon, 2 Jan 2006 15:04:05 GMT"   // RFC 7231#section-7.1.1.1 timetamp format. e.g Tue, 29 Apr 2014 18:30:38 GMT
	ISO8601TimeFormat       = "2006-01-02T15:04:05.999999999Z" // RFC3339 a subset of the ISO8601 timestamp format. e.g 2014-04-29T18:30:38.999999999Z
	RFC822OutputTimeFormat  = "Mon, 02 Jan 2006 15:04:05 GMT"  // RFC Output TimeStamp format is used for output time without seconds precision
	ISO8601OutputTimeFormat = "2006-01-02T15:04:05Z"           // ISO output TimeStamp format is used for output time without seconds precision

)
