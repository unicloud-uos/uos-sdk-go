package main

import (
	"fmt"

	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
)

func MakeBucketSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// TODO: Make bucket with ACL

	// Delete a bucket
	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	//Make bucket with ACL
	err = sc.MakeBucketWithAcl(bucketName, lib.BucketCannedACLPublicRead)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketAcl(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)

	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("CreateBucketSample Run Success!\n\n")
}
