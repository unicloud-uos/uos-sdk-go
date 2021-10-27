package sqs

import "gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
