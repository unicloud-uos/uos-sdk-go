package main

import (
	"bytes"
	"fmt"
	"github.com/unicloud-uos/uos-sdk-go/s3"
	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
)

func MultiPartUploadSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	uploadId, err := sc.CreateMultiPartUpload(bucketName, objectKey, s3.StorageClassStandard)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("UploadId is: ", s3.StringValue(s3.String(uploadId)))
	partCount := 3
	completedUpload := &s3.CompletedMultipartUpload{
		Parts: make([]s3.CompletedPart, partCount),
	}
	for i := 0; i < partCount; i++ {
		partNumber := int64(i + 1)
		etag, err := sc.UploadPart(bucketName, objectKey, lib.GenMinimalPart(), uploadId, partNumber)
		if err != nil {
			HandleError(err)
		}
		completedUpload.Parts[i] = s3.CompletedPart{
			ETag:       s3.String(etag),
			PartNumber: s3.Int64(partNumber),
		}
	}

	err = sc.CompleteMultiPartUpload(bucketName, objectKey, uploadId, completedUpload)
	if err != nil {
		HandleError(err)
		err = sc.AbortMultiPartUpload(bucketName, objectKey, uploadId)
		if err != nil {
			HandleError(err)
		}
	}
	fmt.Printf("MultipartUploadSample Run Success !\n\n")
}

func MultiPartUploadSampleWithForbidOverwrite() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	uploadId, err := sc.CreateMultiPartUpload(bucketName, objectKey, s3.StorageClassStandard)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("UploadId is: ", s3.StringValue(s3.String(uploadId)))
	partCount := 3
	completedUpload := &s3.CompletedMultipartUpload{
		Parts: make([]s3.CompletedPart, partCount),
	}
	for i := 0; i < partCount; i++ {
		partNumber := int64(i + 1)
		etag, err := sc.UploadPart(bucketName, objectKey, lib.GenMinimalPart(), uploadId, partNumber)
		if err != nil {
			HandleError(err)
		}
		completedUpload.Parts[i] = s3.CompletedPart{
			ETag:       s3.String(etag),
			PartNumber: s3.Int64(partNumber),
		}
	}
	err = sc.CompleteMultiPartUpload(bucketName, objectKey, uploadId, completedUpload)
	if err != nil {
		HandleError(err)
		err = sc.AbortMultiPartUpload(bucketName, objectKey, uploadId)
		if err != nil {
			HandleError(err)
		}
	}
	fmt.Printf("MultipartUploadSample Run Success !\n\n")

	// forbid overwrite
	uploadId, err = sc.CreateMultiPartUploadWithForbidOverwrite(bucketName, objectKey, s3.StorageClassStandard, true)
	if err == nil {
		fmt.Println("forbid overwrite success:", err)
	}

	// overwrite
	uploadId, err = sc.CreateMultiPartUploadWithForbidOverwrite(bucketName, objectKey, s3.StorageClassStandard, false)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("UploadId is: ", s3.StringValue(s3.String(uploadId)))
	partCount = 5
	completedUpload = &s3.CompletedMultipartUpload{
		Parts: make([]s3.CompletedPart, partCount),
	}
	for i := 0; i < partCount; i++ {
		partNumber := int64(i + 1)
		etag, err := sc.UploadPart(bucketName, objectKey, lib.GenMinimalPart(), uploadId, partNumber)
		if err != nil {
			HandleError(err)
		}
		completedUpload.Parts[i] = s3.CompletedPart{
			ETag:       s3.String(etag),
			PartNumber: s3.Int64(partNumber),
		}
	}
	err = sc.CompleteMultiPartUpload(bucketName, objectKey, uploadId, completedUpload)
	if err != nil {
		HandleError(err)
		err = sc.AbortMultiPartUpload(bucketName, objectKey, uploadId)
		if err != nil {
			HandleError(err)
		}
	}

	fmt.Printf("overwrite Run Success !\n\n")
}

func MultiPartDownloadSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := lib.NewClient(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Put a 5kib file
	RANGE := make([]byte, 5<<10)
	err = sc.PutObject(bucketName, objectKey, bytes.NewReader(RANGE))
	if err != nil {
		HandleError(err)
	}
	//Slice download by range
	ranges := map[string]string{"0": "1000", "1001": "2000", "2001": "3000", "3001": "4000", "4001": "5119"}
	for range1, range2 := range ranges {
		out, err := sc.GetObjectWithRange(bucketName, objectKey, "bytes="+range1+"-"+range2)
		if err != nil {
			HandleError(err)
		}
		fmt.Println("Download range is :", *out.ContentRange)

	}
	fmt.Printf("MultiPartDownloadSample Run Success !\n\n")
}