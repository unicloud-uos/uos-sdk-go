// +build go1.5

package request_test

import (
	"errors"
	"strings"
	"testing"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/awstesting/mock"
)

func TestRequestCancelRetry(t *testing.T) {
	c := make(chan struct{})

	reqNum := 0
	s := mock.NewMockClient(aws.NewConfig().WithMaxRetries(10))
	s.Handlers.Validate.Clear()
	s.Handlers.Unmarshal.Clear()
	s.Handlers.UnmarshalMeta.Clear()
	s.Handlers.UnmarshalError.Clear()
	s.Handlers.Send.PushFront(func(r *request.Request) {
		reqNum++
		r.Error = errors.New("net/http: request canceled")
	})
	out := &testData{}
	r := s.NewRequest(&request.Operation{Name: "Operation"}, nil, out)
	r.HTTPRequest.Cancel = c
	close(c)

	err := r.Send()
	if !strings.Contains(err.Error(), "canceled") {
		t.Errorf("expect canceled in error, %v", err)
	}
	if e, a := 1, reqNum; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
