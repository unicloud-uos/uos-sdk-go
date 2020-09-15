package main

import (
	"fmt"
	"github.com/uos-sdk-go/sample/lib"
)

func BucketWebsiteSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	//put bucket website
	err = sc.PutBucketWebsite(bucketName)
	if err != nil {
		HandleError(err)
	}

	//Get Bucket Website
	err = sc.GetBucketWebsite(bucketName)
	if err != nil {
		HandleError(err)
	}

	//Delete Bucket Website
	err = sc.DeleteBucketWebsite(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("BucketWebsiteSample Run Success !\n\n")
}
