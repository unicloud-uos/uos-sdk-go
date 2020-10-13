package request

import (
	"crypto/md5"
	"encoding/base64"
	"io"

	. "github.com/unicloud-uos/uos-sdk-go/s3/error"
	"github.com/unicloud-uos/uos-sdk-go/s3/helper"
)

// AddBodyContentMD5Handler computes and sets the HTTP Content-MD5 header for requests that
// require it.
func AddBodyContentMD5Handler(r *Request) {
	// if Content-MD5 header is already present, return
	if v := r.HTTPRequest.Header.Get(helper.ContentMD5Header); len(v) != 0 {
		return
	}

	h := md5.New()

	// hash the body.  seek back to the first position after reading to reset
	// the body for transmission.  copy errors may be assumed to be from the
	// body.
	_, err := io.Copy(h, r.Body)
	if err != nil {
		r.Error = NewBaseError("ContentMD5", "failed to read body", err)
		return
	}
	_, err = r.Body.Seek(0, io.SeekStart)
	if err != nil {
		r.Error = NewBaseError("ContentMD5", "failed to seek body", err)
		return
	}

	// encode the md5 checksum in base64 and set the request header.
	sum := h.Sum(nil)
	sum64 := make([]byte, base64.StdEncoding.EncodedLen(len(sum)))
	base64.StdEncoding.Encode(sum64, sum)
	r.HTTPRequest.Header.Set(helper.ContentMD5Header, string(sum64))
}
