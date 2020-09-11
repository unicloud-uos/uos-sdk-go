package s3

import (
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/request"
)

type DeleteObjectsInput struct {
	Bucket *string
	Delete *Delete
}

// String returns the string representation
func (d DeleteObjectsInput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteObjectsInput"}

	if d.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if d.Delete == nil {
		invalidParams.Add(request.NewErrParamRequired("Delete"))
	}
	if d.Delete != nil {
		if err := d.Delete.Validate(); err != nil {
			invalidParams.AddNested("Delete", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (d DeleteObjectsInput) getBucket() (v string) {
	if d.Bucket == nil {
		return v
	}
	return *d.Bucket
}

func (d DeleteObjectsInput) MarshalForPut(e *request.EncoderForPut) error {
	if d.Bucket != nil {
		v := *d.Bucket
		e.SetValue(helper.PathTarget, "Bucket", request.StringValue(v))
	}
	if d.Delete != nil {
		v := d.Delete
		e.SetStruct(helper.PayloadTarget, "Delete", v)
	}
	return nil
}

type DeleteObjectsOutput struct {
	Deleted []DeletedObject
	Errors  []Error
}

// String returns the string representation
func (d DeleteObjectsOutput) String() string {
	return helper.Prettify(d)
}

func (d DeleteObjectsOutput) UnmarshalBody(r *request.Request) error {
	defer r.HTTPResponse.Body.Close()
	return nil
}

const opDeleteObjects = "DeleteObjects"

func (c *Client) DeleteObjectsRequest(input *DeleteObjectsInput) (req *request.Request) {
	op := &request.Operation{
		Name:       opDeleteObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}?delete",
	}

	if input == nil {
		input = &DeleteObjectsInput{}
	}

	req = c.newRequest(op, input, &DeleteObjectsOutput{})
	req.Handlers.Set.PushBackHandlerItem(request.ContentMd5Handler)

	return
}

func (c *Client) DeleteObjects(input *DeleteObjectsInput) (*DeleteObjectsOutput, error) {
	req := c.DeleteObjectsRequest(input)

	err := req.Do()
	if err != nil {
		return &DeleteObjectsOutput{}, err
	}
	out := req.Data.(*DeleteObjectsOutput)

	return out, err
}
