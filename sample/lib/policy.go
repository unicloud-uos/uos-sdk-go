package lib

import (
	"github.com/unicloud-uos/uos-sdk-go/s3"
)

func (s3client *S3Client) PutBucketPolicy(bucketName, policy string) (err error) {
	params := &s3.PutBucketPolicyInput{
		Bucket: s3.String(bucketName),
		Policy: s3.String(policy),
	}
	_, err = s3client.Client.PutBucketPolicy(params)
	if err != nil {
		return
	}
	return
}

func (s3client *S3Client) GetBucketPolicy(bucketName string) (policy string, err error) {
	params := &s3.GetBucketPolicyInput{
		Bucket: s3.String(bucketName),
	}
	out, err := s3client.Client.GetBucketPolicy(params)
	if err != nil {
		return "", err
	}
	return *out.Policy, err
}

func (s3client *S3Client) DeleteBucketPolicy(bucketName string) (err error) {
	params := &s3.DeleteBucketPolicyInput{
		Bucket: s3.String(bucketName),
	}
	_, err = s3client.Client.DeleteBucketPolicy(params)
	if err != nil {
		return
	}
	return
}
