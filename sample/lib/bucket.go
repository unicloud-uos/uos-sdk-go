package lib

import (
	"github.com/unicloud-uos/uos-sdk-go/s3"
)

func (s3client *S3Client) MakeBucket(bucketName string) (err error) {
	params := &s3.CreateBucketInput{
		Bucket: s3.String(bucketName),
	}

	if _, err = s3client.Client.CreateBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) MakeBucketWithAcl(bucketName string, acl string) (err error) {
	params := &s3.CreateBucketInput{
		Bucket: s3.String(bucketName),
		ACL:    s3.String(acl),
	}
	if _, err = s3client.Client.CreateBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) DeleteBucket(bucketName string) (err error) {
	params := &s3.DeleteBucketInput{
		Bucket: s3.String(bucketName),
	}
	if _, err = s3client.Client.DeleteBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) HeadBucket(bucketName string) (err error) {
	params := &s3.HeadBucketInput{
		Bucket: s3.String(bucketName),
	}
	if _, err = s3client.Client.HeadBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) ListBuckets() (buckets []s3.Bucket, err error) {
	params := &s3.ListBucketsInput{}
	out, err := s3client.Client.ListBuckets(params)
	if err != nil {
		return nil, err
	}
	return out.Buckets, err
}
