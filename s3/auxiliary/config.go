package auxiliary

import (
	"github.com/uos-sdk-go/s3/log"
	"time"

	"github.com/uos-sdk-go/s3/client"
	"github.com/uos-sdk-go/s3/credential"
)

// Config provides service configuration for service clients.
type Config struct {
	Endpoint    string                  // UOS endpoint
	Credentials *credential.Credentials // The credentials will use to request.
	UserName    string                  // The user who send request.
	Region      string                  // The region to send requests to.
	DisableSSL  bool                    // Set this to `true` to disable SSL when sending requests. Defaults to `false`.
	HTTPClient  *client.HTTPClient
	Logger      *log.Logger // For write logging messages
}

func NewConfig() *Config {
	return &Config{}
}

func GetDefaultConfig() *Config {
	config := &Config{}

	config.Endpoint = "DefaultEndPoint"
	config.Credentials = &credential.DefaultCredentials
	config.Region = "DefaultRegion"
	config.DisableSSL = true

	config.HTTPClient.ConnectTimeout = 30 * time.Second
	config.HTTPClient.KeepAliveTimeout = 30 * time.Second
	config.HTTPClient.IdleConnTimeout = 90 * time.Second
	config.HTTPClient.TLSHandleshakeTimeout = 10 * time.Second
	config.HTTPClient.ExpectContinueTimeout = time.Second

	config.HTTPClient.MaxIdleConns = 100
	config.HTTPClient.MaxIdleConnsPerHost = 10

	config.Logger = log.NewLogger(log.LogOffLevel)

	return config
}

func MergeInConfig(target *Config, source *Config) {
	if source == nil {
		return
	}

	if source.Endpoint != "" {
		target.Endpoint = source.Endpoint
	}

	if source.Credentials != nil {
		target.Credentials = source.Credentials
	}

	if source.UserName != "" {
		target.UserName = source.UserName
	}

	if source.Region != "" {
		target.Region = source.Region
	}

	target.DisableSSL = source.DisableSSL

	if source.Logger != nil {
		target.Logger = source.Logger
	}
}
