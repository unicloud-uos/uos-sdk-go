package s3

import (
	"encoding/xml"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
	"io/ioutil"
)

type RenameObjectInput struct {
	Bucket          *string
	Key             *string
	RenameSourceKey *string
}

// String returns the string representation
func (r RenameObjectInput) String() string {
	return helper.Prettify(r)
}

func (r RenameObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RenameObjectInput"}

	if r.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if r.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if r.Key != nil && len(*r.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}
	if r.RenameSourceKey == nil {
		invalidParams.Add(request.NewErrParamRequired("RenameSourceKey"))
	}
	if r.RenameSourceKey != nil && len(*r.RenameSourceKey) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RenameSourceKey", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (r RenameObjectInput) getBucket() (v string) {
	if r.Bucket == nil {
		return v
	}
	return *r.Bucket
}

func (r RenameObjectInput) MarshalForPut(e *request.EncoderForPut) error {
	if r.Bucket != nil {
		v := *r.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if r.Key != nil {
		v := *r.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if r.RenameSourceKey != nil {
		v := *r.Key
		e.SetValue(helper.HeaderTarget, "x-uos-rename-source-key", request.StringValue(v))

	}
	return nil
}

type RenameObjectOutput struct {
	RenameObjectResult *RenameObjectResult
}

// String returns the string representation
func (s RenameObjectOutput) String() string {
	return helper.Prettify(s)
}

func (s RenameObjectOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	buf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &s)
	if err != nil {
		return err
	}

	r.Data = &s
	return nil
}

const opRenameObject = "RenameObject"

func (c *Client) RenameObjectRequest(input *RenameObjectInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opRenameObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &RenameObjectInput{}
	}

	req = c.newRequest(op, input, &RenameObjectOutput{})
	return
}

func (c *Client) RenameObject(input *RenameObjectInput) (RenameObjectOutput, error) {
	req := c.RenameObjectRequest(input)

	err := req.Do()
	if err != nil {
		return RenameObjectOutput{}, err
	}
	out := req.Data.(*RenameObjectOutput)

	return *out, err
}
