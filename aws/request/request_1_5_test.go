// +build !go1.6

package request_test

import (
	"errors"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))
