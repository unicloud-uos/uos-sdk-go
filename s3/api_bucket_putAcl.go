package s3

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type PutBucketAclInput struct {
	ACL                 *string
	AccessControlPolicy *AccessControlPolicy
	Bucket              *string
}

// String returns the string representation
func (p PutBucketAclInput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectInput"}

	if p.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
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

func (p PutBucketAclInput) GetBucketName() (v string) {
	if p.Bucket == nil {
		return v
	}
	return *p.Bucket
}

func (p PutBucketAclInput) MarshalForPut(e *request.EncoderForPut) error {
	if p.Bucket != nil {
		v := *p.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
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

type PutBucketAclOutput struct{}

// String returns the string representation
func (p PutBucketAclOutput) String() string {
	return helper.Prettify(p)
}

func (p PutBucketAclOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opPutBucketAcl = "PutBucketAcl"

func (c *Client) PutBucketAclRequest(input *PutBucketAclInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opPutBucketAcl,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?acl",
	}

	if input == nil {
		input = &PutBucketAclInput{}
	}

	req = c.newRequest(op, input, &PutBucketAclOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) PutBucketAcl(input *PutBucketAclInput) (PutBucketAclOutput, error) {
	req := c.PutBucketAclRequest(input)

	err := req.Do()
	if err != nil {
		return PutBucketAclOutput{}, err
	}
	out := req.Data.(*PutBucketAclOutput)

	return *out, err
}
