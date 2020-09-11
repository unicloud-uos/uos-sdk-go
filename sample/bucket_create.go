package main

import (
	"fmt"

	"github.com/uos-sdk-go/sample/lib"
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

	fmt.Printf("CreateBucketSample Run Success!\n\n")
}
