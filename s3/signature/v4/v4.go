package v4

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/unicloud-uos/uos-sdk-go/s3/credential"
	. "github.com/unicloud-uos/uos-sdk-go/s3/helper"
	"github.com/unicloud-uos/uos-sdk-go/s3/log"
	"github.com/unicloud-uos/uos-sdk-go/s3/request"
)

type Signer struct {
	Request     *http.Request
	Credential  *credential.Credentials
	ServiceName string
	Region      string
	Time        time.Time
	ExpireTime  time.Duration
	IsPreSign   bool
	PayloadHash string
	Logger      *log.Logger
}

var SignV4Handler = request.HandlerItem{
	Name: "v4.sign.Handler",
	Fn: func(r *request.Request) {
		SignForRequest(r)
	},
}

func SignForRequest(req *request.Request, opts ...func(*Signer)) {
	// If the request does not need to be signed ignore the signing of the
	// request if the AnonymousCredentials object is used.
	if req.Config.Credentials == credential.DefaultCredentials {
		return
	}

	region := req.Config.Region
	if region == "" {
		region = req.Metadata.SigningRegion
	}

	name := req.Config.UserName
	if name == "" {
		name = req.Metadata.SigningName
	}
	req.Config.Logger.Debug("Header in signer: ", req.HTTPRequest.Header)
	isPreSign := req.ExpireTime > 0

	bodyDigest, err := buildBodyDigest(req.HTTPRequest, req.GetBody(), name, isPreSign)
	if err != nil {
		req.SignedHeaderValues = http.Header{}
	}

	v4 := &Signer{
		Request:     req.HTTPRequest,
		Credential:  req.Config.Credentials,
		ServiceName: name,
		Region:      region,
		Time:        req.Time.UTC(),
		ExpireTime:  req.ExpireTime,
		IsPreSign:   isPreSign,
		PayloadHash: bodyDigest,
		Logger:      req.Config.Logger,
	}
	r, signedHeaders := v4.sign()

	*req.HTTPRequest = *r
	req.SignedHeaderValues = signedHeaders
}

func (v4 Signer) sign() (*http.Request, http.Header) {
	req := v4.Request.Clone(v4.Request.Context())
	query := req.URL.Query()
	headers := req.Header

	v4.setHeadersForSignV4(headers, query)
	for key := range query {
		sort.Strings(query[key])
	}

	request.ResetHostForHeader(req)

	credentialScope := v4.SetCredentialScope()
	credentialStr := v4.Credential.AccessKeyID + "/" + credentialScope
	if v4.IsPreSign {
		query.Set(UOSCredentialKey, credentialStr)
	}

	unsignedHeaders := headers
	if v4.IsPreSign {
		urlValues := url.Values{}
		urlValues, unsignedHeaders = forQuery(AllowedQueryHoisting, unsignedHeaders)
		for k, v := range urlValues {
			query[k] = v
		}
	}

	host := req.URL.Host
	if len(req.Host) > 0 {
		host = req.Host
	}
	signedHeaders, signedHeadersStr, canonicalHeaderStr := v4.buildCanonicalHeaders(host, IgnoredHeaders, unsignedHeaders)
	if v4.IsPreSign {
		query.Set(UOSSignedHeadersKey, signedHeadersStr)

	}

	rawQuery := strings.Replace(query.Encode(), "+", "%20", -1)

	canonicalURI := GetURIPath(req.URL)
	canonicalString := buildCanonicalString(
		req.Method,
		canonicalURI,
		rawQuery,
		signedHeadersStr,
		canonicalHeaderStr,
		v4.PayloadHash,
	)

	strToSign := buildStringToSign(v4.Time, credentialScope, canonicalString)

	// signingKey = credentials
	signingKey := buildSigningKey(v4.Credential.SecretAccessKey, v4.Time, v4.Region, v4.ServiceName)
	signingSignature := buildSignature(signingKey, strToSign)

	if v4.IsPreSign {
		rawQuery += "&X-Uos-Signature=" + signingSignature
	} else {
		parts := []string{
			"Credential=" + credentialStr,
			"SignedHeaders=" + signedHeadersStr,
			"Signature=" + signingSignature,
		}
		headers.Set("Authorization", UOSSigningV4Algorithm+" "+strings.Join(parts, ", "))
	}
	req.URL.RawQuery = rawQuery

	return req, signedHeaders
}

func (v4 *Signer) setHeadersForSignV4(headers http.Header, query url.Values) {
	date := v4.Time.Format(UOSSignedTimeFormat)

	if v4.IsPreSign {
		query.Set(UOSAlgorithmKey, UOSSigningV4Algorithm)
		if len(v4.Credential.SessionToken) > 0 {
			query.Set(UOSSecurityTokenKey, v4.Credential.SessionToken)
		}

		duration := int64(v4.ExpireTime / time.Second)
		query.Set(UOSDateKey, date)
		query.Set(UOSExpiresKey, strconv.FormatInt(duration, 10))
		return

	}

	headers.Set(UOSDateKey, date)
	if len(v4.Credential.SessionToken) > 0 {
		headers.Set(UOSSecurityTokenKey, v4.Credential.SessionToken)
	}
}

func (v4 *Signer) SetCredentialScope() string {
	return strings.Join([]string{
		v4.Time.Format(UOSSignShortTimeFormat),
		v4.Region,
		v4.ServiceName,
		"uos4_request",
	}, "/")
}

func forQuery(r Rule, header http.Header) (url.Values, http.Header) {
	query := url.Values{}
	unsignedHeaders := http.Header{}
	for k, h := range header {
		if r.IsValid(k) {
			query[k] = h
		} else {
			unsignedHeaders[k] = h
		}
	}

	return query, unsignedHeaders
}

func buildBodyDigest(r *http.Request, body io.ReadSeeker, service string, presigned bool) (string, error) {
	hash := r.Header.Get(UOSContentSha256)
	if hash == "" {
		includeSHA256Header := service == ServiceNameForSign

		s3Presign := presigned && service == ServiceNameForSign

		if s3Presign {
			hash = UnsignedPayload
			includeSHA256Header = !s3Presign
		} else if body == nil {
			hash = EmptyStringSHA256
		} else {
			if !isReaderSeekable(body) {
				return "", fmt.Errorf("cannot use unseekable request body %T, for signed request with body", body)
			}
			hashBytes, err := makeSha256Reader(body)
			if err != nil {
				return "", err
			}
			hash = hex.EncodeToString(hashBytes)
		}

		if includeSHA256Header {
			r.Header.Set(UOSContentSha256, hash)
		}
	}
	return hash, nil
}

func (v4 *Signer) buildCanonicalHeaders(host string, ignoredRule Rule, header http.Header) (signed http.Header, signedHeaders, canonicalHeaders string) {
	signed = make(http.Header)

	var headers []string
	headers = append(headers, "host")
	for k, v := range header {
		canonicalKey := http.CanonicalHeaderKey(k)
		if ignoredRule.IsValid(canonicalKey) {
			// ignored header
			continue
		}

		lowerKey := strings.ToLower(k)
		if _, ok := signed[lowerKey]; ok {
			// include additional values
			signed[lowerKey] = append(signed[lowerKey], v...)
			continue
		}

		headers = append(headers, lowerKey)
		signed[lowerKey] = v
	}
	sort.Strings(headers)

	signedHeaders = strings.Join(headers, ";")
	headerValues := make([]string, len(headers))
	for k, v := range headers {
		if v == "host" {
			headerValues[k] = "host:" + host
		} else {
			headerValues[k] = v + ":" + strings.Join(signed[v], ",")
		}
	}
	StripExcessSpaces(headerValues)
	canonicalHeaders = strings.Join(headerValues, "\n")

	return signed, signedHeaders, canonicalHeaders
}

// buildCanonicalString generate a canonical request of style
func buildCanonicalString(method, uri, query, signedHeaders, canonicalHeaders, payloadHash string) string {
	return strings.Join([]string{
		method,
		uri,
		query,
		canonicalHeaders + "\n",
		signedHeaders,
		payloadHash,
	}, "\n")
}

// buildStringToSign a string based on selected query values.
func buildStringToSign(time time.Time, credentialScope, canonicalRequestString string) string {
	return strings.Join([]string{
		UOSSigningV4Algorithm,
		time.Format(UOSSignedTimeFormat),
		credentialScope,
		hex.EncodeToString(sum256([]byte(canonicalRequestString))),
	}, "\n")
}

// buildSigningKey hmac seed to calculate final signature.
func buildSigningKey(secretKey string, t time.Time, region, serviceName string) []byte {
	date := sumHMAC([]byte("UOS4"+secretKey), []byte(t.Format(UOSSignShortTimeFormat)))
	regionBytes := sumHMAC(date, []byte(region))
	service := sumHMAC(regionBytes, []byte(serviceName))
	signingKey := sumHMAC(service, []byte("uos4_request"))
	return signingKey
}

// buildSignature final signature in hexadecimal form.
func buildSignature(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(sumHMAC(signingKey, []byte(stringToSign)))
}

func isReaderSeekable(r io.Reader) bool {
	switch r.(type) {
	//case ReaderSeekerCloser:
	//	return v.IsSeeker()
	//case *ReaderSeekerCloser:
	//	return v.IsSeeker()
	case io.ReadSeeker:
		return true
	default:
		return false
	}
}

func makeSha256Reader(reader io.ReadSeeker) (hashBytes []byte, err error) {
	hash := sha256.New()
	start, err := reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	defer func() {
		// ensure error is return if unable to seek back to start if payload
		_, err = reader.Seek(start, io.SeekStart)
	}()

	io.Copy(hash, reader)
	return hash.Sum(nil), nil
}
