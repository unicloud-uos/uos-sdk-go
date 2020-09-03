package request

import (
	"net/url"
	"regexp"
	"strings"
)

var SetBucketEndpointHandler = HandlerItem{
	Name: "sdk.request.set.endpoint",
	Fn: func(request *Request) {
		bucket, ok := bucketNameFromReqParams(request.Params)
		if !ok {
			// Ignore operation requests if the bucketname was not provided
			// if this is an input validation error the validation handler
			// will report it.
			return
		}

		if !hostCompatibleBucketName(request.HTTPRequest.URL, bucket) {
			// bucket name must be valid to put into the host
			return
		}

		moveBucketToHost(request.HTTPRequest.URL, bucket)
	},
}

func bucketNameFromReqParams(params interface{}) (string, bool) {
	if iface, ok := params.(bucketGetter); ok {
		b := iface.GetBucketName()
		return b, len(b) > 0
	}

	return "", false
}

// hostCompatibleBucketName returns true if the request should
// put the bucket in the host. This is false if S3ForcePathStyle is
// explicitly set or if the bucket is not DNS compatible.
func hostCompatibleBucketName(u *url.URL, bucket string) bool {
	// Bucket might be DNS compatible but dots in the hostname will fail
	// certificate validation, so do not use host-style.
	if u.Scheme == "https" && strings.Contains(bucket, ".") {
		return false
	}

	// if the bucket is DNS compatible
	return dnsCompatibleBucketName(bucket)
}

var reDomain = regexp.MustCompile(`^[a-z0-9][a-z0-9\.\-]{1,61}[a-z0-9]$`)
var reIPAddress = regexp.MustCompile(`^(\d+\.){3}\d+$`)

// dnsCompatibleBucketName returns true if the bucket name is DNS compatible.
// Buckets created outside of the classic region MUST be DNS compatible.
func dnsCompatibleBucketName(bucket string) bool {
	return reDomain.MatchString(bucket) &&
		!reIPAddress.MatchString(bucket) &&
		!strings.Contains(bucket, "..")
}

// moveBucketToHost moves the bucket name from the URI path to URL host.
func moveBucketToHost(u *url.URL, bucket string) {
	u.Host = bucket + "." + u.Host
	u.Path = strings.Replace(u.Path, "/{Bucket}", "", -1)
	if u.Path == "" {
		u.Path = "/"
	}
}

type bucketGetter interface {
	GetBucketName() string
}