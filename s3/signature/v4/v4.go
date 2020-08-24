package v4

import (
	"encoding/hex"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	. "github.com/uos-sdk-go/s3/auxiliary"
	"github.com/uos-sdk-go/s3/credential"
	"github.com/uos-sdk-go/s3/log"
	"github.com/uos-sdk-go/s3/request"
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

func SignForRequest(req *request.Request, opts ...func(*Signer)) {
	// If the request does not need to be signed ignore the signing of the
	// request if the AnonymousCredentials object is used.
	if req.Config.Credentials == &credential.DefaultCredentials {
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

	isPreSign := req.ExpireTime > 0
	v4 := &Signer{
		Request:     req.HTTPRequest,
		Credential:  req.Config.Credentials,
		ServiceName: name,
		Region:      region,
		Time:        req.Time,
		ExpireTime:  req.ExpireTime,
		IsPreSign:   isPreSign,
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
	v4.Logger.Debug("signedHeaders: %s, signedHeadersStr: %s, canonicalHeaderStr: %s",
		signedHeaders, signedHeadersStr, canonicalHeaderStr)
	if v4.IsPreSign {
		query.Set(UOSSignedHeadersKey, signedHeadersStr)

	}
	v4.Logger.Debug("http header query:", query)

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
	v4.Logger.Debug("Signature: ", signingSignature)

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

func (v4 *Signer) buildCanonicalHeaders(host string, rule Rule, header http.Header) (signed http.Header, signedHeaders, canonicalHeaders string) {
	signed = make(http.Header)

	var headers []string
	headers = append(headers, "host")
	v4.Logger.Debug("header: ", header)
	for k, v := range header {
		canonicalKey := http.CanonicalHeaderKey(k)
		if !rule.IsValid(canonicalKey) {
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
	v4.Logger.Debug("headers: ", headers)
	sort.Strings(headers)

	signedHeaders = strings.Join(headers, ":")
	headerValues := make([]string, len(headers))
	for k, v := range headers {
		if v == "host" {
			headerValues[k] = "host:" + host
		} else {
			headerValues[k] = v + ":" + strings.Join(signed[v], ",")
		}
	}
	v4.Logger.Debug("headerValues: ", headerValues)
	StripExcessSpaces(headerValues)
	v4.Logger.Debug("headerValues: ", headerValues)
	canonicalHeaders = strings.Join(headerValues, "\n")
	v4.Logger.Debug("canonicalHeaders: ", canonicalHeaders)

	return signed, signedHeaders, canonicalHeaders
}

// buildCanonicalString generate a canonical request of style
func buildCanonicalString(method, uri, query, canonicalHeaders, signedHeaders, payloadHash string) string {
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
