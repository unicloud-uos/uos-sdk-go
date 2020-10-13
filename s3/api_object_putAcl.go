package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutObjectAclInput struct {
	ACL                 *string
	AccessControlPolicy *AccessControlPolicy
	Bucket              *string
	Key                 *string
	VersionId           *string
}

// String returns the string representation
func (p PutObjectAclInput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectAclInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if p.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if p.Key != nil && len(*p.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}
	if p.AccessControlPolicy != nil {
		if err := p.AccessControlPolicy.Validate(); err != nil {
			invalidParams.AddNested("AccessControlPolicy", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (p PutObjectAclInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutObjectAclInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if p.Key != nil {
		v := *p.Key
		e.SetValue(helper.PathTarget, "Key", request.StringValue(v))
	}
	if p.VersionId != nil {
		v := *p.VersionId
		e.SetValue(helper.QueryTarget, "versionId", request.StringValue(v))
	}

	if p.ACL != nil {
		v := *p.ACL
		e.SetValue(helper.HeaderTarget, "x-uos-acl", request.StringValue(v))
	}
	if p.AccessControlPolicy != nil {
		v := p.AccessControlPolicy
		e.SetStruct(helper.PayloadTarget, "AccessControlPolicy", v)
	}

	return nil
}

type PutObjectAclOutput struct{}

// String returns the string representation
func (p PutObjectAclOutput) String() string {
	return helper.Prettify(p)
}

func (p PutObjectAclOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutObjectAcl = "PutObjectAcl"

func (c *Client) PutObjectAclRequest(input *PutObjectAclInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutObjectAcl,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?acl",
	}

	if input == nil {
		input = &PutObjectAclInput{}
	}

	req = c.newRequest(op, input, &PutObjectAclOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutObjectAcl(input *PutObjectAclInput) (PutObjectAclOutput, error) {
	req := c.PutObjectAclRequest(input)

	err := req.Do()
	if err != nil {
		return PutObjectAclOutput{}, err
	}
	out := req.Data.(*PutObjectAclOutput)

	return *out, err
}
