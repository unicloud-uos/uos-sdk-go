package main

import (
	"fmt"
	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
	"strings"
)

func ObjectMetaSample() {
	DeleteTestBucketAndObject()

	defer DeleteTestBucketAndObject()

	// Set Custom Meta
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	c := make(map[string]string)
	c["a"] = "b"
	err = sc.PutObjectWithMeta(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"), c)
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("ObjectMetaSample Run Success !\n\n")
}
