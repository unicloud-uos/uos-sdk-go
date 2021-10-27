package apigateway

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
