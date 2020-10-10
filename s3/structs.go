package s3

import (
	"fmt"
	"github.com/uos-sdk-go/s3/helper"
	"time"

	"github.com/uos-sdk-go/s3/request"
)

type AbortIncompleteMultipartUpload struct {
	DaysAfterInitiation *int64
}

// String returns the string representation
func (a AbortIncompleteMultipartUpload) String() string {
	return helper.Prettify(a)
}

// Contains the elements that set the ACL permissions for an object per grantee.
type AccessControlPolicy struct {
	AccessControlList []Grant // A list of grants.
	Owner             *Owner  // Container for the bucket owner's display name and ID.
}

// String returns the string representation
func (a AccessControlPolicy) String() string {
	return helper.Prettify(a)
}

func (a AccessControlPolicy) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AccessControlPolicy"}
	if a.AccessControlList != nil {
		for i, v := range a.AccessControlList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccessControlList", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// This is used in a Lifecycle Rule Filter to apply a logical AND to two or
// more predicates. The Lifecycle Rule will apply to any object matching all
// of the predicates configured inside the And operator.
type And struct {
	Prefix *string
	Tags   []Tag `xml:"Tags>Tag"`
}

// String returns the string representation
func (a And) String() string {
	return helper.Prettify(a)
}

func (a And) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "And"}
	if a.Tags != nil {
		for i, v := range a.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// In terms of implementation, a Bucket is a resource.
type Bucket struct {
	CreationDate *time.Time
	Name         *string
}

// String returns the string representation
func (b Bucket) String() string {
	return helper.Prettify(b)
}

// Container for logging status information.
type BucketLoggingStatus struct {
	LoggingEnabled *LoggingEnabled
}

// String returns the string representation
func (b BucketLoggingStatus) String() string {
	return helper.Prettify(b)
}

func (b BucketLoggingStatus) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "BucketLoggingStatus"}
	if b.LoggingEnabled != nil {
		if err := b.LoggingEnabled.Validate(); err != nil {
			invalidParams.AddNested("LoggingEnabled", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for all response elements.
type CopyPartResult struct {
	ETag         *string
	LastModified *time.Time
}

// String returns the string representation
func (c CopyPartResult) String() string {
	return helper.Prettify(c)
}

// Container for all response elements.
type CopyObjectResult struct {
	ETag         *string
	LastModified *time.Time
}

// String returns the string representation
func (s CopyObjectResult) String() string {
	return helper.Prettify(s)
}

// Container for all (if there are any) keys between Prefix and the next occurrence
// of the string specified by a delimiter. CommonPrefixes lists keys that act
// like subdirectories in the directory specified by Prefix. For example, if
// the prefix is notes/ and the delimiter is a slash (/) as in notes/summer/july,
// the common prefix is notes/summer/.
type CommonPrefix struct {
	Prefix *string
}

// String returns the string representation
func (c CommonPrefix) String() string {
	return helper.Prettify(c)
}

// A container for describing a condition that must be met for the specified
// redirect to apply. For example, 1. If request is for pages in the /docs folder,
// redirect to the /documents folder. 2. If request results in HTTP error 4xx,
// redirect request to another host where you might process the error.
type Condition struct {
	HttpErrorCodeReturnedEquals *string
	KeyPrefixEquals             *string
}

// String returns the string representation
func (c Condition) String() string {
	return helper.Prettify(c)
}

// The container for the completed multipart upload details.
type CompletedMultipartUpload struct {
	Parts []CompletedPart `xml:"Part"`
}

// String returns the string representation
func (c CompletedMultipartUpload) String() string {
	return helper.Prettify(c)
}

// Details of the parts that were uploaded.
type CompletedPart struct {
	ETag       *string
	PartNumber *int64
}

// String returns the string representation
func (c CompletedPart) String() string {
	return helper.Prettify(c)
}

type CORSConfiguration struct {
	CORSRules []CORSRule `xml:"CORSRule"`
}

// String returns the string representation
func (c CORSConfiguration) String() string {
	return helper.Prettify(c)
}

func (c CORSConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CORSConfiguration"}

	if c.CORSRules == nil {
		invalidParams.Add(request.NewErrParamRequired("CORSRules"))
	}
	if c.CORSRules != nil {
		for i, v := range c.CORSRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CORSRules", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies a cross-origin access rule.
type CORSRule struct {
	AllowedHeaders []string `xml:"AllowedHeader"`
	AllowedMethods []string `xml:"AllowedMethod"`
	AllowedOrigins []string `xml:"AllowedOrigin"`
	ExposeHeaders  []string `xml:"ExposeHeader"`
	MaxAgeSeconds  *int64
}

// String returns the string representation
func (c CORSRule) String() string {
	return helper.Prettify(c)
}

func (c CORSRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CORSRule"}

	if c.AllowedMethods == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowedMethods"))
	}

	if c.AllowedOrigins == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowedOrigins"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes how a CSV-formatted input object is formatted.
type CSVInput struct {
	AllowQuotedRecordDelimiter *bool
	Comments                   *string
	FieldDelimiter             *string
	FileHeaderInfo             *string
	QuoteCharacter             *string
	QuoteEscapeCharacter       *string
	RecordDelimiter            *string
}

// String returns the string representation
func (c CSVInput) String() string {
	return helper.Prettify(c)
}

// Describes how CSV-formatted results are formatted.
type CSVOutput struct {
	FieldDelimiter       *string
	QuoteCharacter       *string
	QuoteEscapeCharacter *string
	QuoteFields          *string
	RecordDelimiter      *string
}

// String returns the string representation
func (s CSVOutput) String() string {
	return helper.Prettify(s)
}

// Container for the objects to delete.
type Delete struct {
	Objects []ObjectIdentifier `xml:"Object"`
	Quiet   *bool
}

// String returns the string representation
func (d Delete) String() string {
	return helper.Prettify(d)
}

func (d Delete) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Delete"}

	if d.Objects == nil {
		invalidParams.Add(request.NewErrParamRequired("Objects"))
	}
	if d.Objects != nil {
		for i, v := range d.Objects {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Objects", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the delete marker.
type DeleteMarkerEntry struct {
	IsLatest     *bool
	Key          *string
	LastModified *time.Time
	Owner        *Owner
	VersionId    *string
}

// String returns the string representation
func (d DeleteMarkerEntry) String() string {
	return helper.Prettify(d)
}

// Information about the deleted object.
type DeletedObject struct {
	DeleteMarker          *bool
	DeleteMarkerVersionId *string
	Key                   *string
	VersionId             *string
}

// String returns the string representation
func (d DeletedObject) String() string {
	return helper.Prettify(d)
}

// Container for all error elements.
type Error struct {
	Code      *string
	Key       *string
	Message   *string
	VersionId *string
}

// String returns the string representation
func (e Error) String() string {
	return helper.Prettify(e)
}

// The error information.
type ErrorDocument struct {
	Key *string
}

// String returns the string representation
func (s ErrorDocument) String() string {
	return helper.Prettify(s)
}

func (s ErrorDocument) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ErrorDocument"}

	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the type of server-side encryption used.
type Encryption struct {
	EncryptionType ServerSideEncryption
	KMSContext     *string
	KMSKeyId       *string
}

// String returns the string representation
func (e Encryption) String() string {
	return helper.Prettify(e)
}

func (e Encryption) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Encryption"}
	if len(e.EncryptionType) == 0 {
		invalidParams.Add(request.NewErrParamRequired("EncryptionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the expiration for the lifecycle of the object.
type Expiration struct {
	Date                      *time.Time
	Days                      *int64
	ExpiredObjectDeleteMarker *bool
}

// String returns the string representation
func (e Expiration) String() string {
	return helper.Prettify(e)
}

// The Filter is used to identify objects that a Lifecycle Rule applies to.
// A Filter must have exactly one of Prefix, Tag, or And specified.
type Filter struct {
	And    *And
	Prefix *string
	Tag    *Tag
}

// String returns the string representation
func (f Filter) String() string {
	return helper.Prettify(f)
}

func (f Filter) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Filter"}
	if f.And != nil {
		if err := f.And.Validate(); err != nil {
			invalidParams.AddNested("And", err.(request.ErrInvalidParams))
		}
	}
	if f.Tag != nil {
		if err := f.Tag.Validate(); err != nil {
			invalidParams.AddNested("Tag", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the Suffix element.
type IndexDocument struct {
	Suffix *string
}

// String returns the string representation
func (i IndexDocument) String() string {
	return helper.Prettify(i)
}

func (i IndexDocument) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "IndexDocument"}

	if i.Suffix == nil {
		invalidParams.Add(request.NewErrParamRequired("Suffix"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container element that identifies who initiated the multipart upload.
type Initiator struct {
	DisplayName *string
	ID          *string
}

// String returns the string representation
func (i Initiator) String() string {
	return helper.Prettify(i)
}

// Describes the serialization format of the object.
type InputSerialization struct {
	CSV             *CSVInput
	CompressionType *string
	JSON            *JSONInput
	Parquet         *ParquetInput
}

// String returns the string representation
func (i InputSerialization) String() string {
	return helper.Prettify(i)
}

type JSONInput struct {
	Type *string
}

// String returns the string representation
func (j JSONInput) String() string {
	return helper.Prettify(j)
}

type JSONOutput struct {
	RecordDelimiter *string
}

// String returns the string representation
func (j JSONOutput) String() string {
	return helper.Prettify(j)
}

// A container of a key value name pair.
type Tag struct {
	Key   *string
	Value *string
}

// String returns the string representation
func (t Tag) String() string {
	return helper.Prettify(t)
}

func (t Tag) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Tag"}

	if t.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if t.Key != nil && len(*t.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if t.Value == nil {
		invalidParams.Add(request.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for S3 Glacier job parameters.
type GlacierJobParameters struct {
	Tier *string
}

// String returns the string representation
func (g GlacierJobParameters) String() string {
	return helper.Prettify(g)
}

func (g GlacierJobParameters) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GlacierJobParameters"}
	if g.Tier == nil {
		invalidParams.Add(request.NewErrParamRequired("Tier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for grant information.
type Grant struct {
	Grantee    *Grantee   // The person being granted permissions.
	Permission Permission // Specifies the permission given to the grantee.
}

func (g Grant) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Grant"}
	if g.Grantee != nil {
		if err := g.Grantee.Validate(); err != nil {
			invalidParams.AddNested("Grantee", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the person being granted permissions.
type Grantee struct {
	DisplayName  *string // Screen name of the grantee.
	EmailAddress *string // Email address of the grantee.
	ID           *string // The canonical user ID of the grantee.
	Type         Type    // Type of grantee
	URI          *string // URI of the grantee group.
}

func (g Grantee) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Grantee"}
	if len(g.Type) == 0 {
		invalidParams.Add(request.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for lifecycle rules. You can add as many as 1000 rules.
type LifecycleConfiguration struct {
	Rules []LifecycleRule `xml:"Rule"`
}

// String returns the string representation
func (l LifecycleConfiguration) String() string {
	return helper.Prettify(l)
}

func (l LifecycleConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LifecycleConfiguration"}

	if l.Rules == nil {
		invalidParams.Add(request.NewErrParamRequired("Rules"))
	}
	if l.Rules != nil {
		for i, v := range l.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type LifecycleRule struct {
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload
	Expiration                     *Expiration
	Filter                         *Filter
	ID                             *string
	NoncurrentVersionExpiration    *NoncurrentVersionExpiration
	NoncurrentVersionTransitions   []NoncurrentVersionTransition `xml:"NoncurrentVersionTransition"`
	Prefix                         *string
	Status                         *string      `xml:"Status"`
	Transitions                    []Transition `xml:"Transition"`
}

// String returns the string representation
func (l LifecycleRule) String() string {
	return helper.Prettify(l)
}

func (l LifecycleRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LifecycleRule"}
	if l.Status == nil {
		invalidParams.Add(request.NewErrParamRequired("Status"))
	}
	if l.Filter != nil {
		if err := l.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type LoggingEnabled struct {
	TargetBucket *string
	TargetPrefix *string
	Grantee      *Grantee
	Permission   PermissionForLogs
}

// String returns the string representation
func (l LoggingEnabled) String() string {
	return helper.Prettify(l)
}

func (l LoggingEnabled) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LoggingEnabled"}

	if l.TargetBucket == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetBucket"))
	}

	if l.TargetPrefix == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetPrefix"))
	}
	if l.Grantee != nil {
		if err := l.Grantee.Validate(); err != nil {
			invalidParams.AddNested("Grantee", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MetaConfiguration struct {
	Headers   []MetadataEntry
	VersionID *string
}

// String returns the string representation
func (m MetaConfiguration) String() string {
	return helper.Prettify(m)
}

func (m MetaConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MetaConfiguration"}

	if m.Headers == nil {
		invalidParams.Add(request.NewErrParamRequired("Headers"))
	}
	if m.Headers != nil {
		for i, v := range m.Headers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Headers", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A metadata key-value pair to store with an object.
type MetadataEntry struct {
	Name  *string
	Value *string
}

// String returns the string representation
func (m MetadataEntry) String() string {
	return helper.Prettify(m)
}

func (m MetadataEntry) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MetadataEntry"}

	if m.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if m.Name != nil && len(*m.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}

	if m.Value == nil {
		invalidParams.Add(request.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MultipartUpload struct {
	Initiated    *time.Time
	Initiator    *Initiator
	Key          *string
	Owner        *Owner
	StorageClass StorageClass
	UploadId     *string
}

// String returns the string representation
func (m MultipartUpload) String() string {
	return helper.Prettify(m)
}

type NoncurrentVersionExpiration struct {
	NoncurrentDays *int64
}

// String returns the string representation
func (n NoncurrentVersionExpiration) String() string {
	return helper.Prettify(n)
}

type NoncurrentVersionTransition struct {
	NoncurrentDays *int64
	StorageClass   *string
}

// String returns the string representation
func (n NoncurrentVersionTransition) String() string {
	return helper.Prettify(n)
}

// An object consists of data and its descriptive metadata.
type Object struct {
	ETag         *string
	Key          *string
	LastModified *time.Time
	Owner        *Owner
	Size         *int64
	StorageClass StorageClass
}

// String returns the string representation
func (o Object) String() string {
	return helper.Prettify(o)
}

// Object Identifier is unique value to identify objects.
type ObjectIdentifier struct {
	Key       *string
	VersionId *string
}

// String returns the string representation
func (o ObjectIdentifier) String() string {
	return helper.Prettify(o)
}

func (o ObjectIdentifier) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ObjectIdentifier"}

	if o.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if o.Key != nil && len(*o.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The version of an object.
type ObjectVersion struct {
	ETag         *string
	IsLatest     *bool
	Key          *string
	LastModified *time.Time
	Owner        *Owner
	Size         *int64
	StorageClass *string
	VersionId    *string
}

// String returns the string representation
func (o ObjectVersion) String() string {
	return helper.Prettify(o)
}

// Describes the location where the restore job's output is stored.
type OutputLocation struct {
	S3 *S3Location
}

// String returns the string representation
func (o OutputLocation) String() string {
	return helper.Prettify(o)
}

func (o OutputLocation) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "OutputLocation"}
	if o.S3 != nil {
		if err := o.S3.Validate(); err != nil {
			invalidParams.AddNested("S3", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the owner's display name and ID.
type Owner struct {
	DisplayName *string // Container for the display name of the owner.
	ID          *string // Container for the ID of the owner.
}

// String returns the string representation
func (o Owner) String() string {
	return helper.Prettify(o)
}

// Describes how results of the Select job are serialized.
type OutputSerialization struct {
	CSV  *CSVOutput
	JSON *JSONOutput
}

// String returns the string representation
func (o OutputSerialization) String() string {
	return helper.Prettify(o)
}

type ParquetInput struct{}

// String returns the string representation
func (p ParquetInput) String() string {
	return helper.Prettify(p)
}

// Container for elements related to a part.
type Part struct {
	ETag         *string
	LastModified *time.Time
	PartNumber   *int64
	Size         *int64
}

// String returns the string representation
func (p Part) String() string {
	return helper.Prettify(p)
}

// Specifies how requests are redirected. In the event of an error, you can
// specify a different error code to return.
type Redirect struct {
	HostName             *string
	HttpRedirectCode     *string
	Protocol             Protocol
	ReplaceKeyPrefixWith *string
	ReplaceKeyWith       *string
}

// String returns the string representation
func (s Redirect) String() string {
	return helper.Prettify(s)
}

// Specifies the redirect behavior of all requests to a website endpoint
type RedirectAllRequestsTo struct {
	HostName *string
	Protocol Protocol
}

// String returns the string representation
func (r RedirectAllRequestsTo) String() string {
	return helper.Prettify(r)
}

func (r RedirectAllRequestsTo) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RedirectAllRequestsTo"}

	if r.HostName == nil {
		invalidParams.Add(request.NewErrParamRequired("HostName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RenameObjectResult struct {
	LastModified *time.Time
}

// String returns the string representation
func (r RenameObjectResult) String() string {
	return helper.Prettify(r)
}

// Container for restore job parameters.
type RestoreRequest struct {
	Days                 *int64
	Description          *string
	GlacierJobParameters *GlacierJobParameters
	OutputLocation       *OutputLocation
	SelectParameters     *SelectParameters
	Tier                 *string
	Type                 *string
}

func (r RestoreRequest) String() string {
	return helper.Prettify(r)
}

func (r RestoreRequest) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreRequest"}
	if r.GlacierJobParameters != nil {
		if err := r.GlacierJobParameters.Validate(); err != nil {
			invalidParams.AddNested("GlacierJobParameters", err.(request.ErrInvalidParams))
		}
	}
	if r.OutputLocation != nil {
		if err := r.OutputLocation.Validate(); err != nil {
			invalidParams.AddNested("OutputLocation", err.(request.ErrInvalidParams))
		}
	}
	if r.SelectParameters != nil {
		if err := r.SelectParameters.Validate(); err != nil {
			invalidParams.AddNested("SelectParameters", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the redirect behavior and when a redirect is applied.
type RoutingRule struct {
	Condition *Condition
	Redirect  *Redirect
}

// String returns the string representation
func (s RoutingRule) String() string {
	return helper.Prettify(s)
}

func (s RoutingRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RoutingRule"}

	if s.Redirect == nil {
		invalidParams.Add(request.NewErrParamRequired("Redirect"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the parameters for Select job types.
type SelectParameters struct {
	Expression          *string
	ExpressionType      *string
	InputSerialization  *InputSerialization
	OutputSerialization *OutputSerialization
}

// String returns the string representation
func (s SelectParameters) String() string {
	return helper.Prettify(s)
}

func (s SelectParameters) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SelectParameters"}

	if s.Expression == nil {
		invalidParams.Add(request.NewErrParamRequired("Expression"))
	}
	if s.ExpressionType == nil {
		invalidParams.Add(request.NewErrParamRequired("ExpressionType"))
	}

	if s.InputSerialization == nil {
		invalidParams.Add(request.NewErrParamRequired("InputSerialization"))
	}

	if s.OutputSerialization == nil {
		invalidParams.Add(request.NewErrParamRequired("OutputSerialization"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the default server-side-encryption configuration.
type ServerSideEncryptionConfiguration struct {
	Rules []ServerSideEncryptionRule `xml:"Rule"`
}

// String returns the string representation
func (s ServerSideEncryptionConfiguration) String() string {
	return helper.Prettify(s)
}

func (s ServerSideEncryptionConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ServerSideEncryptionConfiguration"}

	if s.Rules == nil {
		invalidParams.Add(request.NewErrParamRequired("Rules"))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the default server-side encryption configuration.
type ServerSideEncryptionRule struct {
	ApplyServerSideEncryptionByDefault *ServerSideEncryptionByDefault
}

// String returns the string representation
func (s ServerSideEncryptionRule) String() string {
	return helper.Prettify(s)
}

func (s ServerSideEncryptionRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ServerSideEncryptionRule"}
	if s.ApplyServerSideEncryptionByDefault != nil {
		if err := s.ApplyServerSideEncryptionByDefault.Validate(); err != nil {
			invalidParams.AddNested("ApplyServerSideEncryptionByDefault", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ServerSideEncryptionByDefault struct {
	KMSMasterKeyID *string
	SSEAlgorithm   ServerSideEncryption
}

// String returns the string representation
func (s ServerSideEncryptionByDefault) String() string {
	return helper.Prettify(s)
}

func (s ServerSideEncryptionByDefault) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ServerSideEncryptionByDefault"}
	if len(s.SSEAlgorithm) == 0 {
		invalidParams.Add(request.NewErrParamRequired("SSEAlgorithm"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type S3Location struct {
	AccessControlList []Grant
	BucketName        *string
	CannedACL         *string
	Encryption        *Encryption
	Prefix            *string
	StorageClass      StorageClass
	Tagging           *Tagging
	UserMetadata      []MetadataEntry
}

// String returns the string representation
func (s S3Location) String() string {
	return helper.Prettify(s)
}

func (s S3Location) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "S3Location"}

	if s.BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("BucketName"))
	}

	if s.Prefix == nil {
		invalidParams.Add(request.NewErrParamRequired("Prefix"))
	}
	if s.AccessControlList != nil {
		for i, v := range s.AccessControlList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccessControlList", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(request.ErrInvalidParams))
		}
	}
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			invalidParams.AddNested("Tagging", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for TagSet elements.
type Tagging struct {
	TagSet []Tag `xml:"TagSet>Tag"`
}

// String returns the string representation
func (t Tagging) String() string {
	return helper.Prettify(t)
}

func (t Tagging) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Tagging"}

	if t.TagSet == nil {
		invalidParams.Add(request.NewErrParamRequired("TagSet"))
	}
	if t.TagSet != nil {
		for i, v := range t.TagSet {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagSet", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type Transition struct {
	Date         *time.Time
	Days         *int64
	StorageClass *string
}

// String returns the string representation
func (t Transition) String() string {
	return helper.Prettify(t)
}

type VersioningConfiguration struct {
	Status *string
}

// String returns the string representation
func (v VersioningConfiguration) String() string {
	return helper.Prettify(v)
}

type WebsiteConfiguration struct {
	ErrorDocument         *ErrorDocument
	IndexDocument         *IndexDocument
	RedirectAllRequestsTo *RedirectAllRequestsTo
	RoutingRules          []RoutingRule
}

// String returns the string representation
func (w WebsiteConfiguration) String() string {
	return helper.Prettify(w)
}

func (w WebsiteConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "WebsiteConfiguration"}
	if w.ErrorDocument != nil {
		if err := w.ErrorDocument.Validate(); err != nil {
			invalidParams.AddNested("ErrorDocument", err.(request.ErrInvalidParams))
		}
	}
	if w.IndexDocument != nil {
		if err := w.IndexDocument.Validate(); err != nil {
			invalidParams.AddNested("IndexDocument", err.(request.ErrInvalidParams))
		}
	}
	if w.RedirectAllRequestsTo != nil {
		if err := w.RedirectAllRequestsTo.Validate(); err != nil {
			invalidParams.AddNested("RedirectAllRequestsTo", err.(request.ErrInvalidParams))
		}
	}
	if w.RoutingRules != nil {
		for i, v := range w.RoutingRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RoutingRules", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
