package main

import (
	"fmt"
	"github.com/alecthomas/units"
	"github.com/unicloud-uos/uos-sdk-go/s3"
	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
	"io"
	"os"
	"strconv"
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
	if err != nil {
		HandleError(err)
	}
	defer f.Close()
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
	if err != nil {
		HandleError(err)
	}
	defer f2.Close()
	io.Copy(f2, out.Body)
	out.Body.Close()

	fmt.Printf("GetObjectSample Run Success !\n\n")
}

func GetObjectByRange() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	f, err := os.Open(localFilePath)
	if err != nil {
		HandleError(err)
	}
	defer f.Close()

	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.HeadObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Object info: ", out)

	var partOffset int64
	var num int
	var failedDownload []string
	objects := make(map[int]*s3.GetObjectOutput)
	objectSize := s3.Int64Value(out.ContentLength)
	// Get object by parts
	partSize := int64(10 * units.KiB)
	fmt.Println(objectSize / partSize)
	for i := 0; i <= int(objectSize/partSize); i++ {
		num++
		if partOffset+partSize-1 <= objectSize {
			rangeString := "bytes=" + strconv.FormatInt(partOffset, 10) + "-" + strconv.FormatInt(partOffset+(partSize-1), 10)
			object, err := sc.GetObjectWithRange(bucketName, objectKey, rangeString)
			if err != nil {
				fmt.Println("part ", num, " download err: ", err, "\n rangeString: ", rangeString)
				failedDownload = append(failedDownload, rangeString)
				continue
			}

			fmt.Println("part ", num, " object info: ", object)
			objects[num] = object
			partOffset += partSize
		} else if partOffset < objectSize {
			rangeString := "bytes=" + strconv.FormatInt(partOffset, 10) + "-" + strconv.FormatInt(objectSize-1, 10)
			fmt.Println("rangeString: ", rangeString)
			object, err := sc.GetObjectWithRange(bucketName, objectKey, rangeString)
			if err != nil {
				fmt.Println("part ", num, " download err: ", err, "\n rangeString: ", rangeString)
				failedDownload = append(failedDownload, rangeString)
				continue
			}
			fmt.Println("part ", num, " object info: ", object)
			objects[num] = object
			break
		}
	}
	// get parts which download failed before
	for i := 0; i < len(failedDownload); i++ {
		object, err := sc.GetObjectWithRange(bucketName, objectKey, failedDownload[i])
		if err != nil {
			fmt.Println("part ", num, " download err: ", err, "\n rangeString: ", failedDownload[i])
			failedDownload = append(failedDownload, failedDownload[i])
			continue
		}
		fmt.Println("part ", num, " object info: ", object)
		objects[num] = object
	}

	f2, err := os.OpenFile("sample/Download.jpg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		HandleError(err)
	}
	defer f2.Close()

	for i := 1; i <= len(objects); i++ {
		io.Copy(f2, objects[i].Body)
		objects[i].Body.Close()
	}

	fmt.Printf("GetObjectByParts Run Success !\n\n")
}
