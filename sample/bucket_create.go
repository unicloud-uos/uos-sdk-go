package main

import (
	"fmt"

	"github.com/uos-sdk-go/sample/lib"
)

func MakeBucketSample() {
	//DeleteTestBucketAndObject()
	//defer DeleteTestBucketAndObject()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err, out := sc.MakeBucket(bucketName)
	if err != nil {
		fmt.Println(out)
		HandleError(err)
	}

	fmt.Printf("CreateBucketSample Run Success!\n\n", out)
}
