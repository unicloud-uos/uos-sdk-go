package main

import (
	"fmt"
	"github.com/uos-sdk-go/sample/lib"
	"os"
)

func HandleError(err error) {
	fmt.Println("sample err:", err)
	err = DeleteTestBucketAndObject()
	if err != nil {
		fmt.Println("DeleteTestBucketAndObject err:", err)
	}
	os.Exit(-1)
}

func DeleteTestBucketAndObject() error {
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		return err
	}
	err = sc.DeleteBucket(bucketName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTestBucket() error {
	sc := lib.NewClient(endpoint, accessKey, secretKey)

	err := sc.DeleteBucket(bucketName)
	if err != nil {
		return err
	}
	return nil
}
