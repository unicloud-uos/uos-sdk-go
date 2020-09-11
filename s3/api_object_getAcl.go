package s3

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type GetObjectAclInput struct {
	Bucket    *string
	Key       *string
	VersionId *string
}

// String returns the string representation
func (g GetObjectAclInput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetObjectAclInput"}

	if g.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if g.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if g.Key != nil && len(*g.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (g GetObjectAclInput) getBucket() (v string) {
	if g.Bucket == nil {
		return v
	}
	return *g.Bucket
}

func (g GetObjectAclInput) MarshalForPut(e *request.EncoderForPut) error {
	if g.Bucket != nil {
		v := *g.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if g.Key != nil {
		v := *g.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}

	if g.VersionId != nil {
		v := *g.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}
	return nil
}

type GetObjectAclOutput struct {
	AccessControlList []Grant `xml:"AccessControlList>Grant"`
	Owner             *Owner  // Container for the bucket owner's display name and ID.
}

// String returns the string representation
func (g GetObjectAclOutput) String() string {
	return helper.Prettify(g)
}

func (g GetObjectAclOutput) UnmarshalBody(r *request.Request) error {
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

const opGetObjectAcl = "GetObjectAcl"

func (c *Client) GetObjectAclRequest(input *GetObjectAclInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opGetObjectAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?acl",
	}

	if input == nil {
		input = &GetObjectAclInput{}
	}

	req = c.newRequest(op, input, &GetObjectAclOutput{})
	return
}

func (c *Client) GetObjectAcl(input *GetObjectAclInput) (*GetObjectAclOutput, error) {
	req := c.GetObjectAclRequest(input)

	err := req.Do()
	if err != nil {
		return &GetObjectAclOutput{}, err
	}
	out := req.Data.(*GetObjectAclOutput)

	return out, err
}
