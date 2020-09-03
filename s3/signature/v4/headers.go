package v4

// IgnoredHeaders is a list of headers that are ignored during signing
var IgnoredHeaders = Rules{
	MapRule{
		"Authorization":   struct{}{},
		"User-Agent":      struct{}{},
	},
}

// RequiredSignedHeaders is a whitelist for Build canonical headers.
var RequiredSignedHeaders = Rules{
	MapRule{
		"Cache-Control":                         struct{}{},
		"Content-Disposition":                   struct{}{},
		"Content-Encoding":                      struct{}{},
		"Content-Language":                      struct{}{},
		"Content-Md5":                           struct{}{},
		"Content-Type":                          struct{}{},
		"Expires":                               struct{}{},
		"If-Match":                              struct{}{},
		"If-Modified-Since":                     struct{}{},
		"If-None-Match":                         struct{}{},
		"If-Unmodified-Since":                   struct{}{},
		"Range":                                 struct{}{},
		"X-Uos-Acl":                             struct{}{},
		"X-Uos-Copy-Source":                     struct{}{},
		"X-Uos-Copy-Source-If-Match":            struct{}{},
		"X-Uos-Copy-Source-If-Modified-Since":   struct{}{},
		"X-Uos-Copy-Source-If-None-Match":       struct{}{},
		"X-Uos-Copy-Source-If-Unmodified-Since": struct{}{},
		"X-Uos-Copy-Source-Range":               struct{}{},
		"X-Uos-Copy-Source-Server-Side-Encryption-Customer-Algorithm": struct{}{},
		"X-Uos-Copy-Source-Server-Side-Encryption-Customer-Key":       struct{}{},
		"X-Uos-Copy-Source-Server-Side-Encryption-Customer-Key-Md5":   struct{}{},
		"X-Uos-Grant-Full-control":                                    struct{}{},
		"X-Uos-Grant-Read":                                            struct{}{},
		"X-Uos-Grant-Read-Acp":                                        struct{}{},
		"X-Uos-Grant-Write":                                           struct{}{},
		"X-Uos-Grant-Write-Acp":                                       struct{}{},
		"X-Uos-Metadata-Directive":                                    struct{}{},
		"X-Uos-Mfa":                                                   struct{}{},
		"X-Uos-Request-Payer":                                         struct{}{},
		"X-Uos-Server-Side-Encryption":                                struct{}{},
		"X-Uos-Server-Side-Encryption-Uos-Kms-Key-Id":                 struct{}{},
		"X-Uos-Server-Side-Encryption-Customer-Algorithm":             struct{}{},
		"X-Uos-Server-Side-Encryption-Customer-Key":                   struct{}{},
		"X-Uos-Server-Side-Encryption-Customer-Key-Md5":               struct{}{},
		"X-Uos-Storage-Class":                                         struct{}{},
		"X-Uos-Website-Redirect-Location":                             struct{}{},
		"X-Uos-Content-Sha256":                                        struct{}{},
	},
	Prefix{"X-Uos-Meta-"},
}

// AllowedQueryHoisting is a whitelist for Build query headers. The boolean value
// represents whether or not it is a pattern.
var AllowedQueryHoisting = Rules{
	List{RequiredSignedHeaders},
	Prefix{"X-Uos-"},
}
