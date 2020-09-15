package lib

import (
	"bytes"
	"github.com/uos-sdk-go/s3"
)

func (s3client *S3Client) CreateMultiPartUpload(bucketName, key string, storageClass s3.StorageClass, ) (uploadId string, err error) {
	params := &s3.CreateMultipartUploadInput{
		Bucket:       s3.String(bucketName),
		Key:          s3.String(key),
		StorageClass: storageClass,
	}
	out, err := s3client.Client.CreateMultipartUpload(params)
	if err != nil {
		return
	}
	return *out.UploadId, nil
}

func (s3client *S3Client) CreateMultiPartUploadWithForbidOverwrite(bucketName, key string, storageClass s3.StorageClass, forbidOverwrite bool) (uploadId string, err error) {
	params := &s3.CreateMultipartUploadInput{
		Bucket:          s3.String(bucketName),
		Key:             s3.String(key),
		StorageClass:    storageClass,
		ForbidOverwrite: s3.Bool(forbidOverwrite),
	}
	out, err := s3client.Client.CreateMultipartUpload(params)
	if err != nil {
		return
	}
	return *out.UploadId, nil
}

func (s3client *S3Client) UploadPart(bucketName, key string, value []byte, uploadId string, partNumber int64) (etag string, err error) {
	params := &s3.UploadPartInput{
		Bucket:     s3.String(bucketName),
		Key:        s3.String(key),
		Body:       bytes.NewReader(value),
		PartNumber: s3.Int64(partNumber),
		UploadId:   s3.String(uploadId),
	}
	out, err := s3client.Client.UploadPart(params)
	if err != nil {
		return
	}
	return *out.ETag, nil
}

func (s3client *S3Client) ListMultiPartUpload(bucketName, key string, uploadId *string) (result []s3.Part, err error) {
	params := &s3.ListPartsInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		UploadId: uploadId,
	}
	out, err := s3client.Client.ListParts(params)
	if err != nil {
		return nil, err
	}
	return out.Parts, err
}

func (s3client *S3Client) CompleteMultiPartUpload(bucketName, key, uploadId string, completed *s3.CompletedMultipartUpload) (err error) {
	params := &s3.CompleteMultipartUploadInput{
		Bucket:          s3.String(bucketName),
		Key:             s3.String(key),
		MultipartUpload: completed,
		UploadId:        s3.String(uploadId),
	}
	if _, err = s3client.Client.CompleteMultipartUpload(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) AbortMultiPartUpload(bucketName, key, uploadId string) (err error) {
	params := &s3.AbortMultipartUploadInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		UploadId: s3.String(uploadId),
	}
	_, err = s3client.Client.AbortMultipartUpload(params)
	return
}
