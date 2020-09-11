package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type ListObjectVersionsInput struct {
	Bucket          *string
	Delimiter       *string
	EncodingType    EncodingType
	KeyMarker       *string
	MaxKeys         *int64
	Prefix          *string
	VersionIdMarker *string
}

// String returns the string representation
func (l ListObjectVersionsInput) String() string {
	return helper.Prettify(l)
}

func (l ListObjectVersionsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListObjectVersionsInput"}

	if l.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (l ListObjectVersionsInput) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListObjectVersionsInput) MarshalForPut(e *request.EncoderForPut) error {
	if l.Bucket != nil {
		v := *l.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if l.Delimiter != nil {
		v := *l.Delimiter
		e.SetValue(helper.QueryTarget, "delimiter", request.StringValue(v))
	}
	if len(l.EncodingType) > 0 {
		v := l.EncodingType
		e.SetValue(helper.QueryTarget, "encoding-type", v)
	}
	if l.KeyMarker != nil {
		v := *l.KeyMarker
		e.SetValue(helper.QueryTarget, "key-marker", request.StringValue(v))
	}
	if l.MaxKeys != nil {
		v := *l.MaxKeys
		e.SetValue(helper.QueryTarget, "max-keys", request.Int64Value(v))
	}
	if l.Prefix != nil {
		v := *l.Prefix
		e.SetValue(helper.QueryTarget, "prefix", request.StringValue(v))
	}
	if l.VersionIdMarker != nil {
		v := *l.VersionIdMarker
		e.SetValue(helper.QueryTarget, "version-id-marker", request.StringValue(v))
	}

	return nil
}

type ListObjectVersionsOutput struct {
	CommonPrefixes      []CommonPrefix      `xml:"CommonPrefixes>CommonPrefix"`
	DeleteMarkers       []DeleteMarkerEntry `xml:"DeleteMarkers>DeleteMarkerEntry"`
	Delimiter           *string
	EncodingType        EncodingType
	IsTruncated         *bool
	KeyMarker           *string
	MaxKeys             *int64
	Name                *string
	NextKeyMarker       *string
	NextVersionIdMarker *string
	Prefix              *string
	VersionIdMarker     *string
	Versions            []ObjectVersion `xml:"Versions>ObjectVersion"`
}

// String returns the string representation
func (l ListObjectVersionsOutput) String() string {
	return helper.Prettify(l)
}

func (l ListObjectVersionsOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &l)
	if err != nil {
		return err
	}

	r.Data = &l
	return nil
}

const opListObjectVersions = "ListObjectVersions"

func (c *Client) ListObjectVersionsRequest(input *ListObjectVersionsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opListObjectVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?versions",
	}

	if input == nil {
		input = &ListObjectVersionsInput{}
	}

	req = c.newRequest(op, input, &ListObjectVersionsOutput{})
	return
}

func (c *Client) ListObjectVersions(input *ListObjectVersionsInput) (*ListObjectVersionsOutput, error) {
	req := c.ListObjectVersionsRequest(input)

	err := req.Do()
	if err != nil {
		return &ListObjectVersionsOutput{}, err
	}
	out := req.Data.(*ListObjectVersionsOutput)

	return out, err
}
