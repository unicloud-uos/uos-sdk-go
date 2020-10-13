package s3

import (
	"encoding/xml"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
	"io/ioutil"
)

type ListMultipartUploadsInput struct {
	Bucket         *string
	Delimiter      *string
	EncodingType   EncodingType
	KeyMarker      *string
	MaxUploads     *int64
	Prefix         *string
	UploadIdMarker *string
}

// String returns the string representation
func (l ListMultipartUploadsInput) String() string {
	return helper.Prettify(l)
}

func (l ListMultipartUploadsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListMultipartUploadsInput"}

	if l.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (l ListMultipartUploadsInput) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListMultipartUploadsInput) MarshalForPut(e *request.EncoderForPut) error {
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
	if l.MaxUploads != nil {
		v := *l.MaxUploads
		e.SetValue(helper.QueryTarget, "max-uploads", request.Int64Value(v))
	}
	if l.Prefix != nil {
		v := *l.Prefix
		e.SetValue(helper.QueryTarget, "prefix", request.StringValue(v))
	}
	if l.UploadIdMarker != nil {
		v := *l.UploadIdMarker
		e.SetValue(helper.QueryTarget, "upload-id-marker", request.StringValue(v))
	}

	return nil
}

type ListMultipartUploadsOutput struct {
	Bucket             *string
	CommonPrefixes     []CommonPrefix `xml:"CommonPrefixes>CommonPrefix"`
	Delimiter          *string
	EncodingType       EncodingType
	IsTruncated        *bool
	KeyMarker          *string
	MaxUploads         *int64
	NextKeyMarker      *string
	NextUploadIdMarker *string
	Prefix             *string
	UploadIdMarker     *string
	Uploads            []MultipartUpload `xml:"Uploads>MultipartUpload"`
}

// String returns the string representation
func (l ListMultipartUploadsOutput) String() string {
	return helper.Prettify(l)
}

func (l *ListMultipartUploadsOutput) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListMultipartUploadsOutput) UnmarshalBody(r *request.Request) error {
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

const opListMultipartUploads = "ListMultipartUploads"

func (c *Client) ListMultipartUploadsRequest(input *ListMultipartUploadsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opListMultipartUploads,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?uploads",
	}

	if input == nil {
		input = &ListMultipartUploadsInput{}
	}

	req = c.newRequest(op, input, &ListMultipartUploadsOutput{})

	return
}

func (c *Client) ListMultipartUploads(input *ListMultipartUploadsInput) (*ListMultipartUploadsOutput, error) {
	req := c.ListMultipartUploadsRequest(input)

	err := req.Do()
	if err != nil {
		return &ListMultipartUploadsOutput{}, err
	}
	out := req.Data.(*ListMultipartUploadsOutput)

	return out, err
}
