package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type ListPartsInput struct {
	Bucket           *string
	Key              *string
	MaxParts         *int64
	PartNumberMarker *int64
	UploadId         *string
}

// String returns the string representation
func (l ListPartsInput) String() string {
	return helper.Prettify(l)
}

func (l ListPartsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListPartsInput"}

	if l.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if l.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if l.Key != nil && len(*l.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if l.UploadId == nil {
		invalidParams.Add(request.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (l ListPartsInput) getBucket() (v string) {
	if l.Bucket == nil {
		return v
	}
	return *l.Bucket
}

func (l ListPartsInput) MarshalForPut(e *request.EncoderForPut) error {
	if l.Bucket != nil {
		v := *l.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if l.Key != nil {
		v := *l.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if l.MaxParts != nil {
		v := *l.MaxParts
		e.SetValue(helper.QueryTarget, "max-parts", request.Int64Value(v))
	}
	if l.PartNumberMarker != nil {
		v := *l.PartNumberMarker
		e.SetValue(helper.QueryTarget, "part-number-marker", request.Int64Value(v))
	}
	if l.UploadId != nil {
		v := *l.UploadId
		e.SetValue(helper.QueryTarget, "uploadId", request.StringValue(v))
	}

	return nil
}

type ListPartsOutput struct {
	Bucket               *string
	Initiator            *Initiator
	IsTruncated          *bool
	Key                  *string
	MaxParts             *int64
	NextPartNumberMarker *int64
	Owner                *Owner
	PartNumberMarker     *int64
	Parts                []Part `xml:"Parts>Part"`
	StorageClass         StorageClass
	UploadId             *string
}

// String returns the string representation
func (l ListPartsOutput) String() string {
	return helper.Prettify(l)
}

func (l ListPartsOutput) UnmarshalBody(r *request.Request) error {
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

const opListParts = "ListParts"

func (c *Client) ListPartsRequest(input *ListPartsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opListParts,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &ListPartsInput{}
	}

	req = c.newRequest(op, input, &ListPartsOutput{})

	return
}

func (c *Client) ListParts(input *ListPartsInput) (*ListPartsOutput, error) {
	req := c.ListPartsRequest(input)

	err := req.Do()
	if err != nil {
		return &ListPartsOutput{}, err
	}
	out := req.Data.(*ListPartsOutput)

	return out, err
}
