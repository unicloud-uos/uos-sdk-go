package lib

import "github.com/uos-sdk-go/s3"

func (s3client *S3Client) PutBucketWebsite(bucketName string) error {
	params := &s3.PutBucketWebsiteInput{
		Bucket: s3.String(bucketName),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: s3.String("index.html"),
			},
			ErrorDocument: &s3.ErrorDocument{
				Key: s3.String("error.html"),
			},
		},
	}
	_, err := s3client.Client.PutBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}

func (s3client *S3Client) GetBucketWebsite(bucketName string) error {
	params := &s3.GetBucketWebsiteInput{
		Bucket: s3.String(bucketName),
	}
	_, err := s3client.Client.GetBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}
func (s3client *S3Client) DeleteBucketWebsite(bucketName string) error {
	params := &s3.DeleteBucketWebsiteInput{
		Bucket: s3.String(bucketName),
	}
	_, err := s3client.Client.DeleteBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}
