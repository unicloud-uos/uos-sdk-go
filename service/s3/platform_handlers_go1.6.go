// +build go1.6

package s3

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

func platformRequestHandlers(r *request.Request) {
	if r.Operation.HTTPMethod == "PUT" {
		// 100-Continue should only be used on put requests.
		r.Handlers.Sign.PushBack(add100Continue)
	}
}

func add100Continue(r *request.Request) {
	if aws.BoolValue(r.Config.S3Disable100Continue) {
		return
	}
	if r.HTTPRequest.ContentLength < 1024*1024*2 {
		// Ignore requests smaller than 2MB. This helps prevent delaying
		// requests unnecessarily.
		return
	}

	r.HTTPRequest.Header.Set("Expect", "100-Continue")
}