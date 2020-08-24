package client

import "time"

type Metadata struct {
	ServiceName string
	ServiceID   string
	APIVersion  string

	SigningName   string
	SigningRegion string
}

type HTTPClient struct {
	ConnectTimeout        time.Duration
	KeepAliveTimeout      time.Duration
	IdleConnTimeout       time.Duration
	TLSHandleshakeTimeout time.Duration
	ExpectContinueTimeout time.Duration
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
}
