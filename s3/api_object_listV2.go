package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type ListObjectsV2Input struct {
	Bucket            *string
	ContinuationToken *string
	Delimiter         *string
	EncodingType      EncodingType
	FetchOwner        *bool
	MaxKeys           *int64
	Prefix            *string
	StartAfter        *string
}

// String returns the string representation
func (l ListObjectsV2Input) String() string {
	return helper.Prettify(l)
}

func (l ListObjectsV2Input) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListObjectsV2Input"}

	if l.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (l ListObjectsV2Input) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListObjectsV2Input) MarshalForPut(e *request.EncoderForPut) error {
	if l.Bucket != nil {
		v := *l.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if l.ContinuationToken != nil {
		v := *l.ContinuationToken
		e.SetValue(helper.QueryTarget, "continuation-token", request.StringValue(v))
	}
	if l.Delimiter != nil {
		v := *l.Delimiter
		e.SetValue(helper.QueryTarget, "delimiter", request.StringValue(v))
	}
	if len(l.EncodingType) > 0 {
		v := l.EncodingType
		e.SetValue(helper.QueryTarget, "encoding-type", v)
	}
	if l.FetchOwner != nil {
		v := *l.FetchOwner
		e.SetValue(helper.QueryTarget, "fetch-owner", request.BoolValue(v))
	}
	if l.MaxKeys != nil {
		v := *l.MaxKeys
		e.SetValue(helper.QueryTarget, "max-keys", request.Int64Value(v))
	}
	if l.Prefix != nil {
		v := *l.Prefix
		e.SetValue(helper.QueryTarget, "prefix", request.StringValue(v))
	}
	if l.StartAfter != nil {
		v := *l.StartAfter
		e.SetValue(helper.QueryTarget, "start-after", request.StringValue(v))
	}
	return nil
}

type ListObjectsV2Output struct {
	CommonPrefixes        []CommonPrefix `xml:"CommonPrefixes>CommonPrefix"`
	Contents              []Object       `xml:"Contents>Object"`
	ContinuationToken     *string
	Delimiter             *string
	EncodingType          EncodingType
	IsTruncated           *bool
	KeyCount              *int64
	MaxKeys               *int64
	Name                  *string
	NextContinuationToken *string
	Prefix                *string
	StartAfter            *string
}

// String returns the string representation
func (l ListObjectsV2Output) String() string {
	return helper.Prettify(l)
}

func (l ListObjectsV2Output) UnmarshalBody(r *request.Request) error {
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

const opListObjectsV2 = "ListObjectsV2"

func (c *Client) ListObjectsV2Request(input *ListObjectsV2Input) (req *request.Request) {
	op := &request.Operation{
		Name:       opListObjectsV2,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?list-type=2",
	}

	if input == nil {
		input = &ListObjectsV2Input{}
	}

	req = c.newRequest(op, input, &ListObjectsV2Output{})
	return
}

func (c *Client) ListObjectsV2(input *ListObjectsV2Input) (*ListObjectsV2Output, error) {
	req := c.ListObjectsV2Request(input)

	err := req.Do()
	if err != nil {
		return &ListObjectsV2Output{}, err
	}
	out := req.Data.(*ListObjectsV2Output)

	return out, err
}
