package main

import (
	"fmt"
	"strings"

	"github.com/uos-sdk-go/sample/lib"
)

func ObjectACLSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Test ObjectACL public-read
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"))
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObjectAcl(bucketName, objectKey, lib.BucketCannedACLPublicRead)
	if err != nil {
		HandleError(err)
	}
	out, err := sc.GetObjectAcl(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)
	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// Test ObjectACL public-read-write
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"))
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObjectAcl(bucketName, objectKey, lib.BucketCannedACLPublicReadWrite)
	if err != nil {
		HandleError(err)
	}
	out, err = sc.GetObjectAcl(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)
	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("ObjectACLSample Run Success!\n\n")
}