package lib

import (
	"github.com/uos-sdk-go/s3"
)

func (s3client *S3Client) PutBucketVersion(bucketName string, status string) error {
	params := &s3.PutBucketVersioningInput{
		Bucket: s3.String(bucketName),
		VersioningConfiguration: &s3.VersioningConfiguration{
			Status: s3.String(status),
		},
	}
	_, err := s3client.Client.PutBucketVersioning(params)
	if err != nil {
		return err
	}
	return nil
}

func (s3client *S3Client) GetBucketVersion(bucketName string) (string, error) {
	params := &s3.GetBucketVersioningInput{
		Bucket: s3.String(bucketName),
	}
	out, err := s3client.Client.GetBucketVersioning(params)
	if err != nil {
		return "", err
	}
	if out.Status == nil {
		return "", nil
	}
	return *out.Status, nil
}

func (s3client *S3Client) GetObjectVersionOutPut(bucketName, key, versionId string) (out *s3.GetObjectOutput, err error) {
	params := &s3.GetObjectInput{
		Bucket:    s3.String(bucketName),
		Key:       s3.String(key),
		VersionId: s3.String(versionId),
	}
	return s3client.Client.GetObject(params)
}

func (s3client *S3Client) ListObjectVersions(bucketName string, prefix string, delimiter, VersionIdMarker string, maxKey int64) (out *s3.ListObjectVersionsOutput, err error) {
	params := &s3.ListObjectVersionsInput{
		Bucket:          s3.String(bucketName),
		MaxKeys:         s3.Int64(maxKey),
		Delimiter:       s3.String(delimiter),
		Prefix:          s3.String(prefix),
		VersionIdMarker: s3.String(VersionIdMarker),
	}
	return s3client.Client.ListObjectVersions(params)
}

func (s3client *S3Client) DeleteObjectVersion(bucketName, objectName, version string) (*s3.DeleteObjectOutput, error) {
	params := &s3.DeleteObjectInput{
		Bucket:    s3.String(bucketName),
		Key:       s3.String(objectName),
		VersionId: s3.String(version),
	}
	return s3client.Client.DeleteObject(params)
}
