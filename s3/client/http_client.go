package client

import (
	"net"
	"net/http"
	"time"
)

// Defaults for the HTTPTransportBuilder.
var (
	DefaultHTTPTransportMaxIdleConns        = 100
	DefaultHTTPTransportMaxIdleConnsPerHost = 10

	DefaultHTTPTransportIdleConnTimeout       = 90 * time.Second
	DefaultHTTPTransportTLSHandleshakeTimeout = 10 * time.Second
	DefaultHTTPTransportExpectContinueTimeout = 1 * time.Second
)

// Timeouts for net.Dialer's network connection.
var (
	DefaultDialConnectTimeout   = 30 * time.Second
	DefaultDialKeepAliveTimeout = 30 * time.Second
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: defaultHTTPTransport(),
	}

}

func defaultHTTPTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   DefaultDialConnectTimeout,
		KeepAlive: DefaultDialKeepAliveTimeout,
		DualStack: true,
	}

	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		TLSHandshakeTimeout:   DefaultHTTPTransportTLSHandleshakeTimeout,
		MaxIdleConns:          DefaultHTTPTransportMaxIdleConns,
		MaxIdleConnsPerHost:   DefaultHTTPTransportMaxIdleConnsPerHost,
		IdleConnTimeout:       DefaultHTTPTransportIdleConnTimeout,
		ExpectContinueTimeout: DefaultHTTPTransportExpectContinueTimeout,
		ForceAttemptHTTP2:     true,
	}

	return tr
}

func DisableHTTPRedirect(client *http.Client) {
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}


