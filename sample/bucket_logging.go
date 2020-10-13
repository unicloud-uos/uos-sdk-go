package main

import (
	"fmt"
	"github.com/unicloud-uos/uos-sdk-go/s3"
	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
)

func BucketLoggingSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	err = sc.MakeBucket("targetbucket")
	if err != nil {
		HandleError(err)
	}

	// SetBucketLogging(bucketName, logBucketName, "prefix")
	rules := &s3.LoggingEnabled{
		TargetBucket: s3.String("targetbucket"),
		TargetPrefix: s3.String("MyBucketLogs/"),
	}
	err = sc.PutBucketLogging(bucketName, rules)
	if err != nil {
		HandleError(err)
	}
	// GetBucketLogging(bucketName)
	a, err := sc.GetBucketLogging(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println(a)
	// DeleteBucketLogging(bucketName)
	err = sc.PutBucketLogging(bucketName, nil)
	if err != nil {
		HandleError(err)
	}
	err = sc.DeleteBucket("targetbucket")
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("BucketLoggingSample Run Success !\n\n")
}
