package main

import (
	"fmt"
	"os"

	"github.com/uos-sdk-go/sample/lib"
)

func HeadObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Put a file
	f, err := os.Open(localFilePath)
	defer f.Close()
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.HeadObject("examplebucket", "v2ray-linux-64.zip")
	//out, err := sc.HeadObject("examplebucket", "exampleobject")
	if err != nil {
		HandleError(err)
	}
	fmt.Println("HeadObject result: ", out)

	fmt.Printf("HeadObjectSample Run Success !\n\n")
}
