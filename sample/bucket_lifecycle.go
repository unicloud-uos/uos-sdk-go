package main

import (
	"fmt"
	"time"

	"github.com/unicloud-uos/uos-sdk-go/s3"
	"github.com/unicloud-uos/uos-sdk-go/sample/lib"
)

func BucketLifecycleSample() {
	DeleteTestBucket()
	defer DeleteTestBucket()

	sc := lib.NewClient(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	date, err := time.Parse(time.RFC3339, "2020-04-16T00:00:00+08:00")
	if err != nil {
		HandleError(err)
	}

	lifecycle := &s3.LifecycleConfiguration{
		Rules: []s3.LifecycleRule{
			// AbortIncompleteMultipartUpload,action on all bucket
			{
				AbortIncompleteMultipartUpload: &s3.AbortIncompleteMultipartUpload{
					DaysAfterInitiation: s3.Int64(7),
				},
				ID:     s3.String("id1"),
				Status: s3.String("Enabled"),
			},
			//Expiration,action on "log" dir
			{
				Expiration: &s3.Expiration{
					Date: s3.Time(date),
				},
				Filter: &s3.Filter{
					Prefix: s3.String("log/"),
				},
				ID:     s3.String("id2"),
				Status: s3.String("Enabled"),
			},
			// NoncurrentVersionExpiration,action on "time" dir
			{
				NoncurrentVersionExpiration: &s3.NoncurrentVersionExpiration{
					NoncurrentDays: s3.Int64(30),
				},
				Filter: &s3.Filter{
					And: &s3.And{
						Prefix: s3.String("time/"),
						Tags: []s3.Tag{
							{
								Key:   s3.String("Key1"),
								Value: s3.String("Value1"),
							},
							{
								Key:   s3.String("Key2"),
								Value: s3.String("Value2"),
							},
						},
					},
				},
				ID:     s3.String("id3"),
				Status: s3.String("Enabled"),
			},
			// Transitions, storageClass turn into GLACIER, action on "prefix" dir
			{
				Transitions: []s3.Transition{
					{
						Date:         s3.Time(date),
						StorageClass: s3.String("GLACIER"),
					},
				},
				Filter: &s3.Filter{
					And: &s3.And{
						Tags: []s3.Tag{
							{
								Key:   s3.String("Key1"),
								Value: s3.String("Value1"),
							},
							{
								Key:   s3.String("Key2"),
								Value: s3.String("Value2"),
							},
						},
					},
				},
				ID:     s3.String("id4"),
				Status: s3.String("Enabled"),
			},
			// NoncurrentVersionTransitions,storageClass turn into STANDARD_IA,but it's not work
			{
				Filter: &s3.Filter{
					Prefix: s3.String("test/"),
				},
				ID:     s3.String("id5"),
				Status: s3.String("Disabled"),
				NoncurrentVersionTransitions: []s3.NoncurrentVersionTransition{
					{
						NoncurrentDays: s3.Int64(30),
						StorageClass:   s3.String("STANDARD_IA"),
					},
				},
			},
		},
	}

	err = sc.PutBucketLifecycle(bucketName, lifecycle)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketLifecycle(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket LifecycleConfiguration", out)

	out, err = sc.DeleteBucketLifecycle(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Delete Bucket LifecycleConfiguration", out)

	fmt.Printf("BucketLifecycleSample Run Success !\n\n")
}
