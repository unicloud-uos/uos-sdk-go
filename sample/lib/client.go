package lib

import (
	"github.com/uos-sdk-go/s3"
	"github.com/uos-sdk-go/s3/credential"
	"github.com/uos-sdk-go/s3/helper"
	"github.com/uos-sdk-go/s3/log"
)

type S3Client struct {
	Client *s3.Client
}

func NewClient(endpoint, accessKey, secretKey string) *S3Client {
	config := helper.GetDefaultConfig()
	config.Endpoint = endpoint
	config.Credentials = credential.NewStaticCredentials(accessKey, secretKey, "")
	config.Logger = log.NewLogger("debug")

	client := s3.NewClient(*config)

	return &S3Client{Client: client}
}
