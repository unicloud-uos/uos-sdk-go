package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type ListBucketsInput struct{}

// String returns the string representation
func (l ListBucketsInput) String() string {
	return helper.Prettify(l)
}

func (l ListBucketsInput) Validate() error {

	return nil
}

func (l ListBucketsInput) MarshalForPut(e *request.EncoderForPut) error {

	return nil
}

type ListBucketsOutput struct {
	Buckets []Bucket `xml:"Buckets>Bucket"`
	Owner   *Owner
}

// String returns the string representation
func (l ListBucketsOutput) String() string {
	return helper.Prettify(l)
}

func (l ListBucketsOutput) UnmarshalBody(r *request.Request) error {
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

const opListBuckets = "ListBuckets"

func (c *Client) ListBucketsRequest(input *ListBucketsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opListBuckets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBucketsInput{}
	}

	req = c.newRequest(op, input, &ListBucketsOutput{})
	return
}

func (c *Client) ListBuckets(input *ListBucketsInput) (*ListBucketsOutput, error) {
	req := c.ListBucketsRequest(input)

	err := req.Do()
	if err != nil {
		return &ListBucketsOutput{}, err
	}
	out := req.Data.(*ListBucketsOutput)

	return out, err
}
