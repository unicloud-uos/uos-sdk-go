package helper

import (
	"github.com/unicloud-uos/uos-sdk-go/s3/credential"
	"github.com/unicloud-uos/uos-sdk-go/s3/log"
)

// Config provides service configuration for service clients.
type Config struct {
	Endpoint        string                  // UOS endpoint
	Credentials     *credential.Credentials // The credentials will use to request.
	UserName        string                  // The user who send request.
	Region          string                  // The region to send requests to.
	DisableSSL      bool                    // Set this to `true` to disable SSL when sending requests. Defaults to `false`.
	RedirectEnabled bool                    // Set this to `true` to allow redirect
	Logger          *log.Logger             // For write logging messages
}

func NewConfig() *Config {
	return &Config{}
}

func GetDefaultConfig() *Config {
	config := &Config{}

	config.Endpoint = "DefaultEndPoint"
	config.Credentials = credential.DefaultCredentials
	config.Region = "DefaultRegion"
	config.DisableSSL = true
	config.RedirectEnabled = false

	config.Logger = log.NewLogger("off")

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
	target.RedirectEnabled = source.RedirectEnabled

	if source.Logger != nil {
		target.Logger = source.Logger
	}
}
