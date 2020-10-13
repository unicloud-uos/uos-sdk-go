package lib

import "github.com/unicloud-uos/uos-sdk-go/s3"

const (
	// BucketCannedACLPrivate is a BucketCannedACL enum value
	BucketCannedACLPrivate = "private"

	// BucketCannedACLPublicRead is a BucketCannedACL enum value
	BucketCannedACLPublicRead = "public-read"

	// BucketCannedACLPublicReadWrite is a BucketCannedACL enum value
	BucketCannedACLPublicReadWrite = "public-read-write"

	// BucketCannedACLAuthenticatedRead is a BucketCannedACL enum value
	BucketCannedACLAuthenticatedRead = "authenticated-read"

	// ObjectCannedACLPrivate is a ObjectCannedACL enum value
	ObjectCannedACLPrivate = "private"

	// ObjectCannedACLPublicRead is a ObjectCannedACL enum value
	ObjectCannedACLPublicRead = "public-read"

	// ObjectCannedACLPublicReadWrite is a ObjectCannedACL enum value
	ObjectCannedACLPublicReadWrite = "public-read-write"

	// ObjectCannedACLAuthenticatedRead is a ObjectCannedACL enum value
	ObjectCannedACLAuthenticatedRead = "authenticated-read"

	// ObjectCannedACLBucketOwnerRead is a ObjectCannedACL enum value
	ObjectCannedACLBucketOwnerRead = "bucket-owner-read"

	// ObjectCannedACLBucketOwnerFullControl is a ObjectCannedACL enum value
	ObjectCannedACLBucketOwnerFullControl = "bucket-owner-full-control"
)

func (s3client *S3Client) PutBucketAcl(bucketName string, acl string, policy s3.AccessControlPolicy) (err error) {
	params := &s3.PutBucketAclInput{
		Bucket:              s3.String(bucketName),
		ACL:                 s3.String(acl),
		AccessControlPolicy: &policy,
	}
	_, err = s3client.Client.PutBucketAcl(params)
	return err
}

func (s3client *S3Client) PutBucketAclWithXml(bucketName string, acl *s3.AccessControlPolicy) (err error) {
	params := &s3.PutBucketAclInput{
		Bucket:              s3.String(bucketName),
		AccessControlPolicy: acl,
	}
	_, err = s3client.Client.PutBucketAcl(params)
	return err
}

func (s3client *S3Client) GetBucketAcl(bucketName string) (out *s3.GetBucketAclOutput, err error) {
	params := &s3.GetBucketAclInput{
		Bucket: s3.String(bucketName),
	}
	out, err = s3client.Client.GetBucketAcl(params)
	return out, err
}

func (s3client *S3Client) PutObjectAcl(bucketName, objName string, acl string) (err error) {
	params := &s3.PutObjectAclInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(objName),
		ACL:    s3.String(acl),
	}
	_, err = s3client.Client.PutObjectAcl(params)
	return err
}

func (s3client *S3Client) PutObjectAclWithXml(bucketName, objName string, acl *s3.AccessControlPolicy) (err error) {
	params := &s3.PutObjectAclInput{
		Bucket:              s3.String(bucketName),
		Key:                 s3.String(objName),
		AccessControlPolicy: acl,
	}
	_, err = s3client.Client.PutObjectAcl(params)
	return err
}

func (s3client *S3Client) GetObjectAcl(bucketName, objName string) (ret string, err error) {
	params := &s3.GetObjectAclInput{
		Bucket: s3.String(bucketName),
		Key:    s3.String(objName),
	}
	out, err := s3client.Client.GetObjectAcl(params)
	return out.String(), err
}
