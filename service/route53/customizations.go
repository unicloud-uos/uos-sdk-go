package route53

import (
	"net/url"
	"regexp"

	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awserr"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/client"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/private/protocol/restxml"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(sanitizeURL)
	}

	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opChangeResourceRecordSets:
			r.Handlers.UnmarshalError.Remove(restxml.UnmarshalErrorHandler)
			r.Handlers.UnmarshalError.PushBack(unmarshalChangeResourceRecordSetsError)
		}
	}
}

var reSanitizeURL = regexp.MustCompile(`\/%2F\w+%2F`)

func sanitizeURL(r *request.Request) {
	r.HTTPRequest.URL.RawPath =
		reSanitizeURL.ReplaceAllString(r.HTTPRequest.URL.RawPath, "/")

	// Update Path so that it reflects the cleaned RawPath
	updated, err := url.Parse(r.HTTPRequest.URL.RawPath)
	if err != nil {
		r.Error = awserr.New("SerializationError", "failed to clean Route53 URL", err)
		return
	}

	// Take the updated path so the requests's URL Path has parity with RawPath.
	r.HTTPRequest.URL.Path = updated.Path
}