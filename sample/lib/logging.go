package lib

import "github.com/unicloud-uos/uos-sdk-go/s3"

func (s3client *S3Client) PutBucketLogging(bucketName string, rules *s3.LoggingEnabled) (err error) {
	params := &s3.PutBucketLoggingInput{
		Bucket:              s3.String(bucketName),
		BucketLoggingStatus: &s3.BucketLoggingStatus{LoggingEnabled: rules},
	}
	if _, err = s3client.Client.PutBucketLogging(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) GetBucketLogging(bucketName string) (rules *s3.GetBucketLoggingOutput, err error) {
	params := &s3.GetBucketLoggingInput{
		Bucket: s3.String(bucketName),
	}
	rules, err = s3client.Client.GetBucketLogging(params)
	if err != nil {
		return nil, err
	}
	return rules, nil
}