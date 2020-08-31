package lib

import (
	"github.com/uos-sdk-go/s3"
)

func (s3client *S3Client) MakeBucket(bucketName string) (err error) {
	params := &s3.CreateBucketInput{
		Bucket: bucketName,
	}

	if _, err = s3client.Client.CreateBucket(params); err != nil {
		return err
	}
	return
}
