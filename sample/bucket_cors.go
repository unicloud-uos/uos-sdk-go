package main

import (
	"fmt"
	"github.com/uos-sdk-go/sample/lib"

	"github.com/uos-sdk-go/s3"
)

func BucketCORSSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	rule1 := s3.CORSRule{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST"},
		AllowedHeaders: []string{},
		ExposeHeaders:  []string{},
		MaxAgeSeconds:  s3.Int64(100),
	}

	rule2 := s3.CORSRule{
		AllowedOrigins: []string{"http://www.a.com", "http://www.b.com"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Authorization"},
		ExposeHeaders:  []string{"x-amz-test", "x-amz-test1"},
		MaxAgeSeconds:  s3.Int64(100),
	}

	err = sc.SetBucketCORS(bucketName, []s3.CORSRule{rule1})
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketCORS(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println(out)

	err = sc.SetBucketCORS(bucketName, []s3.CORSRule{rule1, rule2})
	if err != nil {
		HandleError(err)
	}

	out, err = sc.GetBucketCORS(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println(out)

	err = sc.DeleteBucketCORS(bucketName)
	if err != nil {
		HandleError(err)
	}

	out, err = sc.GetBucketCORS(bucketName)
	fmt.Println(out, err)

	fmt.Printf("BucketCORSSample Run Success !\n\n")
}
