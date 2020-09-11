package lib

import "github.com/uos-sdk-go/s3"

func (s3client *S3Client) PutBucketLifecycle(bucketName string, config *s3.BucketLifecycle) (err error) {
	params := &s3.PutBucketLifecycleInput{
		Bucket:                 s3.String(bucketName),
		Lifecycle: config,
	}
	_, err = s3client.Client.PutBucketLifecycleConfiguration(params)
	return err
}

func (s3client *S3Client) GetBucketLifecycle(bucketName string) (ret string, err error) {
	params := &s3.GetBucketLifecycleInput{
		Bucket: s3.String(bucketName),
	}
	out, err := s3client.Client.GetBucketLifecycle(params)
	return out.String(), err
}

func (s3client *S3Client) DeleteBucketLifecycle(bucketName string) (ret string, err error) {
	params := &s3.DeleteBucketLifecycleInput{
		Bucket: s3.String(bucketName),
	}
	out, err := s3client.Client.DeleteBucketLifecycle(params)
	return out.String(), err
}
