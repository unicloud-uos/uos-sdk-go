package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type ListObjectsInput struct {
	Bucket       *string
	Delimiter    *string
	EncodingType EncodingType
	Marker       *string
	MaxKeys      *int64
	Prefix       *string
}

// String returns the string representation
func (l ListObjectsInput) String() string {
	return helper.Prettify(l)
}

func (l ListObjectsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListObjectsInput"}

	if l.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (l ListObjectsInput) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListObjectsInput) MarshalForPut(e *request.EncoderForPut) error {
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
	if l.Marker != nil {
		v := *l.Marker
		e.SetValue(helper.QueryTarget, "marker", request.StringValue(v))
	}
	if l.MaxKeys != nil {
		v := *l.MaxKeys
		e.SetValue(helper.QueryTarget, "max-keys", request.Int64Value(v))
	}
	if l.Prefix != nil {
		v := *l.Prefix
		e.SetValue(helper.QueryTarget, "prefix", request.StringValue(v))
	}

	return nil
}

type ListObjectsOutput struct {
	CommonPrefixes []CommonPrefix
	Contents       []Object
	Delimiter      *string
	EncodingType   EncodingType
	IsTruncated    *bool
	Marker         *string
	MaxKeys        *int64
	Bucket         *string
	NextMarker     *string
	Prefix         *string
}

// String returns the string representation
func (l ListObjectsOutput) String() string {
	return helper.Prettify(l)
}

func (l ListObjectsOutput) UnmarshalBody(r *request.Request) error {
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

const opListObjects = "ListObjects"

func (c *Client) ListObjectsRequest(input *ListObjectsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opListObjects,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &ListObjectsInput{}
	}

	req = c.newRequest(op, input, &ListObjectsOutput{})

	return
}

func (c *Client) ListObjects(input *ListObjectsInput) (*ListObjectsOutput, error) {
	req := c.ListObjectsRequest(input)

	err := req.Do()
	if err != nil {
		return &ListObjectsOutput{}, err
	}
	out := req.Data.(*ListObjectsOutput)

	return out, err
}
