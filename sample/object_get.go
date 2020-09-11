package main

import (
	"fmt"
	"io"
	"os"

	"github.com/uos-sdk-go/sample/lib"
)

func GetObjectSample() {
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

	// Get the reader
	out, err := sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println(out)

	//Download to a file
	f2, err := os.OpenFile("sample/Download.jpg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer f2.Close()
	if err != nil {
		HandleError(err)
	}
	io.Copy(f2, out.Body)
	out.Body.Close()

	fmt.Printf("GetObjectSample Run Success !\n\n")
}
