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

func (a *And) Validate() error {
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

// Validate inspects the fields of the type to determine if they are valid.
func (b *BucketLoggingStatus) Validate() error {
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

type CORSConfiguration struct {
	CORSRules []CORSRule
}

// String returns the string representation
func (c CORSConfiguration) String() string {
	return helper.Prettify(c)
}

func (c *CORSConfiguration) Validate() error {
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
	AllowedHeaders []string
	AllowedMethods []string
	AllowedOrigins []string
	ExposeHeaders  []string
	MaxAgeSeconds  *int64
}

// String returns the string representation
func (c CORSRule) String() string {
	return helper.Prettify(c)
}

func (c *CORSRule) Validate() error {
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

// Container for the objects to delete.
type Delete struct {
	Objects []ObjectIdentifier // The objects to delete.
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

func (s *ErrorDocument) Validate() error {
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

// Validate inspects the fields of the type to determine if they are valid.
func (f *Filter) Validate() error {
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

func (i *IndexDocument) Validate() error {
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

// A container of a key value name pair.
type Tag struct {
	Key   *string
	Value *string
}

// String returns the string representation
func (t Tag) String() string {
	return helper.Prettify(t)
}

// Validate inspects the fields of the type to determine if they are valid.
func (t *Tag) Validate() error {
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
	Rules []LifecycleRule
}

// String returns the string representation
func (l LifecycleConfiguration) String() string {
	return helper.Prettify(l)
}

// Validate inspects the fields of the type to determine if they are valid.
func (l *LifecycleConfiguration) Validate() error {
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

// A lifecycle rule for individual objects in an Amazon S3 bucket.
type LifecycleRule struct {
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload
	Expiration                     *Expiration
	Filter                         *Filter
	ID                             *string
	NoncurrentVersionExpiration    *NoncurrentVersionExpiration
	NoncurrentVersionTransitions   []NoncurrentVersionTransition `xml:"NoncurrentVersionTransitions>NoncurrentVersionTransition"`
	Prefix                         *string
	Status                         *string
	Transitions                    []Transition `xml:"Transitions>Transition"`
}

// String returns the string representation
func (l LifecycleRule) String() string {
	return helper.Prettify(l)
}

func (l *LifecycleRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LifecycleRule"}
	if l.Status != nil {
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

func (l *LoggingEnabled) Validate() error {
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

// Container for the MultipartUpload for the Amazon S3 object.
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

// Container for the owner's display name and ID.
type Owner struct {
	DisplayName *string // Container for the display name of the owner.
	ID          *string // Container for the ID of the owner.
}

// String returns the string representation
func (o Owner) String() string {
	return helper.Prettify(o)
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

func (r *RedirectAllRequestsTo) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RedirectAllRequestsTo"}

	if r.HostName == nil {
		invalidParams.Add(request.NewErrParamRequired("HostName"))
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

func (s *RoutingRule) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RoutingRule"}

	if s.Redirect == nil {
		invalidParams.Add(request.NewErrParamRequired("Redirect"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Specifies the default server-side-encryption configuration.
type ServerSideEncryptionConfiguration struct {
	Rules []ServerSideEncryptionRule `xml:"Rules>ServerSideEncryptionRule"`
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
	ApplyServerSideEncryptionByDefault *ServerSideEncryptionByDefault `type:"structure"`
}

// String returns the string representation
func (s ServerSideEncryptionRule) String() string {
	return helper.Prettify(s)
}

func (s *ServerSideEncryptionRule) Validate() error {
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

type Transition struct {
	Date         *time.Time
	Days         *int64
	StorageClass *string
}

// String returns the string representation
func (t Transition) String() string {
	return helper.Prettify(t)
}
