package main

import (
	"fmt"

	"github.com/uos-sdk-go/s3"
	"github.com/uos-sdk-go/sample/lib"
)

func BucketACLSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucket()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Set Bucket CannedACL 'PublicRead'
	policy := s3.AccessControlPolicy{
		AccessControlList: []s3.Grant{
			{
				Grantee: &s3.Grantee{
					ID:   s3.String("dsgsfdhgfdhd"),
					Type: s3.TypeGroup,
				},
				Permission: s3.PermissionFullControl,
			},
			{
				Grantee: &s3.Grantee{
					ID:   s3.String("dagfsgbfdshfj"),
					Type: s3.TypeGroup,
				},
				Permission: s3.PermissionFullControl,
			},
		},
		Owner: &s3.Owner{
			ID: s3.String("ewrtrreytu"),
		},
	}
	err = sc.PutBucketAcl(bucketName, lib.ObjectCannedACLPublicReadWrite, policy)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketAcl(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)

	fmt.Printf("BucketACLSample Run Success!\n\n")
}
