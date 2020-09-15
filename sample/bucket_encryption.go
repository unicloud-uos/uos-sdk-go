package main

import (
	"fmt"
	"github.com/uos-sdk-go/s3"
	"github.com/uos-sdk-go/sample/lib"
)

func BucketEncryptionSample() {
	DeleteTestBucket()
	defer DeleteTestBucket()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	encryption := &s3.ServerSideEncryptionConfiguration{
		Rules: []s3.ServerSideEncryptionRule{
			{
				ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
					SSEAlgorithm: s3.ServerSideEncryptionAes256,
				},
			},
		},
	}

	err = sc.PutBucketEncryption(bucketName,encryption)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketEncryption(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket Encryption", out)

	out, err = sc.DeleteBucketEncryption(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Delete Bucket Encryption", out)

	fmt.Printf("BucketLifecycleSample Run Success !\n\n")
}