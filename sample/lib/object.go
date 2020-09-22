package lib

import (
	"fmt"
	"io"
	"time"

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

func (s3client *S3Client) PutObjectWithMeta(bucketName, key string, value io.Reader, meta map[string]string) (err error) {
	params := &s3.PutObjectInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		Body:     s3.ReadSeekCloser(value),
		Metadata: meta,
	}
	if _, err = s3client.Client.PutObject(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) PutObjectWithForbidOverwrite(bucketName, key string, value io.Reader, forbidOverwrite bool) (out s3.PutObjectOutput, err error) {
	params := &s3.PutObjectInput{
		Bucket:          s3.String(bucketName),
		Key:             s3.String(key),
		Body:            s3.ReadSeekCloser(value),
		ForbidOverwrite: s3.Bool(forbidOverwrite),
	}
	return s3client.Client.PutObject(params)
}

func (s3client *S3Client) PutObjectPreSignedWithSpecifiedBody(bucketName, key string, value io.Reader, expire time.Duration) (url string, err error) {
	params := &s3.PutObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
		Body:   s3.ReadSeekCloser(value),
	}
	req := s3client.Client.PutObjectRequest(params)
	return req.Presign(expire)
}

func (s3client *S3Client) PutObjectPreSignedWithoutSpecifiedBody(bucketName, key string, expire time.Duration) (url string, err error) {
	params := &s3.PutObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
	}
	req := s3client.Client.PutObjectRequest(params)
	return req.Presign(expire)
}

func (s3client *S3Client) HeadObject(bucketName, key string) (out *s3.HeadObjectOutput, err error) {
	params := &s3.HeadObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
	}
	out, err = s3client.Client.HeadObject(params)
	if err != nil {
		return nil, err
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

func (s3client *S3Client) GetObjectWithRange(bucketName, key, Range string) (out *s3.GetObjectOutput, err error) {
	params := &s3.GetObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
		Range:  s3.String(Range),
	}
	out, err = s3client.Client.GetObject(params)
	if err != nil {
		return
	}
	return out, err
}

func (s3client *S3Client) GetObjectPreSigned(bucketName, key string, expire time.Duration) (url string, err error) {
	params := &s3.GetObjectInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(key),
	}
	req := s3client.Client.GetObjectRequest(params)
	return req.Presign(expire)
}

func (s3client *S3Client) CopyObject(bucketName, sourceObject, newName string) (err error) {
	params := &s3.CopyObjectInput{
		Bucket:     s3.String(bucketName),
		CopySource: s3.String(sourceObject),
		Key:        s3.String(newName),
	}
	if _, err = s3client.Client.CopyObject(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) CopyObjectWithForbidOverwrite(bucketName, sourceObject, newName string, forbidOverwrite bool) (out s3.CopyObjectOutput, err error) {
	params := &s3.CopyObjectInput{
		Bucket:          s3.String(bucketName),
		CopySource:      s3.String(sourceObject),
		Key:             s3.String(newName),
		ForbidOverwrite: s3.Bool(forbidOverwrite),
	}
	return s3client.Client.CopyObject(params)
}

func (s3client *S3Client) AppendObject(bucketName, key string, value io.ReadSeeker, position int64) (nextPos int64, err error) {
	var out s3.AppendObjectOutput
	params := &s3.AppendObjectInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		Body:     value,
		Position: s3.Int64(position),
	}
	if out, err = s3client.Client.AppendObject(params); err != nil {
		return 0, err
	}

	return *out.NextPosition, nil
}

func (s3client *S3Client) AppendObjectWithAclAndMeta(bucketName, key string, value io.ReadSeeker, position int64, acl string, meta map[string]string) (nextPos int64, err error) {
	var out s3.AppendObjectOutput
	params := &s3.AppendObjectInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		Body:     value,
		Position: s3.Int64(position),
		ACL:      s3.String(acl),
		Metadata: meta,
	}
	if out, err = s3client.Client.AppendObject(params); err != nil {
		return 0, err
	}

	return *out.NextPosition, nil
}

func (s3client *S3Client) GetObjectNextAppendPosition(bucketName, key string, position int64) (nextPos int64, err error) {
	var out s3.AppendObjectOutput
	params := &s3.AppendObjectInput{
		Bucket:   s3.String(bucketName),
		Key:      s3.String(key),
		Position: s3.Int64(position),
	}
	if out, err = s3client.Client.AppendObject(params); err != nil {
		return 0, err
	}

	return *out.NextPosition, nil
}

func (s3client *S3Client) ListObjects(bucketName string, prefix string, delimiter string, maxKey int64) (
	keys []string, isTruncated bool, nextMarker string, err error) {
	params := &s3.ListObjectsInput{
		Bucket:    s3.String(bucketName),
		MaxKeys:   s3.Int64(maxKey),
		Delimiter: s3.String(delimiter),
		Prefix:    s3.String(prefix),
	}
	out, err := s3client.Client.ListObjects(params)
	if err != nil {
		return
	}
	isTruncated = *out.IsTruncated
	if out.NextMarker != nil {
		nextMarker = *out.NextMarker
	}
	for _, v := range out.CommonPrefixes {
		keys = append(keys, *v.Prefix)
		fmt.Println("Prefix:", *v.Prefix)
	}
	for _, v := range out.Contents {
		keys = append(keys, *v.Key)
		fmt.Println("Key:", *v.Key)
	}

	return
}
