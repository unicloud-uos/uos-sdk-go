// +build !go1.7

package request

import "gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"

// setContext updates the Request to use the passed in context for cancellation.
// Context will also be used for request retry delay.
//
// Creates shallow copy of the http.Request with the WithContext method.
func setRequestContext(r *Request, ctx aws.Context) {
	r.context = ctx
	r.HTTPRequest.Cancel = ctx.Done()
}
