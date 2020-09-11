package s3

type EncodingType string

// Enum values for EncodingType
const (
	EncodingTypeUrl EncodingType = "url"
)

func (et EncodingType) MarshalValue() (string, error) {
	return string(et), nil
}

type ServerSideEncryption string

// Enum values for ServerSideEncryption
const (
	ServerSideEncryptionAes256 ServerSideEncryption = "AES256"
	ServerSideEncryptionUosKms ServerSideEncryption = "uos:kms"
)

func (se ServerSideEncryption) MarshalValue() (string, error) {
	return string(se), nil
}

type StorageClass string

// Enum values for StorageClass
const (
	StorageClassStandard           StorageClass = "STANDARD"
	StorageClassReducedRedundancy  StorageClass = "REDUCED_REDUNDANCY"
	StorageClassStandardIa         StorageClass = "STANDARD_IA"
	StorageClassOnezoneIa          StorageClass = "ONEZONE_IA"
	StorageClassIntelligentTiering StorageClass = "INTELLIGENT_TIERING"
	StorageClassGlacier            StorageClass = "GLACIER"
	StorageClassDeepArchive        StorageClass = "DEEP_ARCHIVE"
)

func (sc StorageClass) MarshalValue() (string, error) {
	return string(sc), nil
}

type Permission string

// Enum values for Permission
const (
	PermissionFullControl Permission = "FULL_CONTROL"
	PermissionWrite       Permission = "WRITE"
	PermissionWriteAcp    Permission = "WRITE_ACP"
	PermissionRead        Permission = "READ"
	PermissionReadAcp     Permission = "READ_ACP"
)

func (p Permission) MarshalValue() (string, error) {
	return string(p), nil
}

type PermissionForLogs string

// Enum values for BucketLogsPermission
const (
	BucketLogsPermissionFullControl PermissionForLogs = "FULL_CONTROL"
	BucketLogsPermissionRead        PermissionForLogs = "READ"
	BucketLogsPermissionWrite       PermissionForLogs = "WRITE"
)

func (p PermissionForLogs) MarshalValue() (string, error) {
	return string(p), nil
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolHttp  Protocol = "http"
	ProtocolHttps Protocol = "https"
)

func (p Protocol) MarshalValue() (string, error) {
	return string(p), nil
}

type Type string

// Enum values for Type
const (
	TypeCanonicalUser         Type = "CanonicalUser"
	TypeAmazonCustomerByEmail Type = "AmazonCustomerByEmail"
	TypeGroup                 Type = "Group"
)
