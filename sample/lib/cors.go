package lib

import "github.com/unicloud-uos/uos-sdk-go/s3"

func (s3client *S3Client) SetBucketCORS(bucketName string, rules []s3.CORSRule) error {
	input := &s3.PutBucketCorsInput{
		Bucket: s3.String(bucketName),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: rules,
		},
	}
	_, err := s3client.Client.PutBucketCors(input)
	if err != nil {
		return err
	}
	return nil
}

func (s3client *S3Client) GetBucketCORS(bucketName string) (string, error) {
	input := &s3.GetBucketCorsInput{
		Bucket: s3.String(bucketName),
	}

	out, err := s3client.Client.GetBucketCors(input)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (s3client *S3Client) DeleteBucketCORS(bucketName string) error {
	input := &s3.DeleteBucketCorsInput{
		Bucket: s3.String(bucketName),
	}
	_, err := s3client.Client.DeleteBucketCors(input)
	if err != nil {
		return err
	}
	return nil
}
