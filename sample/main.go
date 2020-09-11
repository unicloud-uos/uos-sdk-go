package main

import "fmt"

func main() {
	// Read S3 config
	ReadConfig()

	MakeBucketSample()
	//BucketACLSample()


	//PutObjectSample()

	//GetObjectSample()

	fmt.Println("All samples completed !")
}