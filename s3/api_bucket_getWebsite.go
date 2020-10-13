package s3

import (
	"encoding/xml"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
	"io/ioutil"
)

type GetBucketWebsiteInput struct {
	Bucket *string
}

// String returns the string representation
func (g GetBucketWebsiteInput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketWebsiteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBucketWebsiteInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetBucketWebsiteInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetBucketWebsiteInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	return nil
}

type GetBucketWebsiteOutput struct {
	ErrorDocument         *ErrorDocument
	IndexDocument         *IndexDocument
	RedirectAllRequestsTo *RedirectAllRequestsTo
	RoutingRules          []RoutingRule `xml:"RoutingRules>RoutingRule"`
}

// String returns the string representation
func (g GetBucketWebsiteOutput) String() string {
	return helper.Prettify(g)
}

func (g GetBucketWebsiteOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &g)
	if err != nil {
		return err
	}

	r.Data = &g
	return nil
}

const opGetBucketWebsite = "GetBucketWebsite"

func (c *Client) GetBucketWebsiteRequest(input *GetBucketWebsiteInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetBucketWebsite,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?website",
	}

	if input == nil {
		input = &GetBucketWebsiteInput{}
	}

	req = c.newRequest(op, input, &GetBucketWebsiteOutput{})
	return
}

func (c *Client) GetBucketWebsite(input *GetBucketWebsiteInput) (*GetBucketWebsiteOutput, error) {
	req := c.GetBucketWebsiteRequest(input)

	err := req.Do()
	if err != nil {
		return &GetBucketWebsiteOutput{}, err
	}
	out := req.Data.(*GetBucketWebsiteOutput)

	return out, err
}
