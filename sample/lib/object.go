package lib

import (
	"io"

	"github.com/uos-sdk-go/s3"
)

func (s3client *S3Client) DeleteObject(bucketName, key string) (err error) {
	params := &s3.DeleteObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
	}
	_, err = s3client.Client.DeleteObject(params)
	if err != nil {
		return err
	}
	return
}

func (s3client *S3Client) DeleteObjects(bucketName string, key ...string) (deletedKeys []string, err error) {
	var objects []s3.ObjectIdentifier
	for _, k := range key {
		oi := s3.ObjectIdentifier{
			Key: s3.String(k),
		}
		objects = append(objects, oi)
	}

	params := &s3.DeleteObjectsInput{
		Bucket: s3.String(bucketName),
		Delete: &s3.Delete{
			Objects: objects,
		},
	}
	out, err := s3client.Client.DeleteObjects(params)
	if err != nil {
		return nil, err
	}
	for _, dk := range out.Deleted {
		deletedKeys = append(deletedKeys, *dk.Key)
	}
	return
}

func (s3client *S3Client) PutObject(bucketName, key string, value io.Reader) (err error) {
	params := &s3.PutObjectInput{
		Bucket:               s3.String(bucketName),
		Key:                  s3.String(key),
		Body:                 s3.ReadSeekCloser(value),
		ServerSideEncryption: s3.ServerSideEncryptionAes256,
	}
	if _, err = s3client.Client.PutObject(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) GetObject(bucketName, key string) (out *s3.GetObjectOutput, err error) {
	params := &s3.GetObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
	}
	out, err = s3client.Client.GetObject(params)
	if err != nil {
		return nil, err
	}
	return out, err
}