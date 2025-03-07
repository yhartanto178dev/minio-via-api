package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
)

func MakeBucket(NameBucket, minioUrl, accessKeyID, secreetAccessKey string) (string, error) {
	// Make a new bucket called my-bucketname.
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	err := minioClient.MakeBucket(context.Background(), NameBucket, minio.MakeBucketOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully created", NameBucket)
	return "Success", nil
}

func ListBuckets(minioUrl, accessKeyID, secreetAccessKey string) ([]minio.BucketInfo, error) {
	// List all buckets from the server
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return buckets, nil
}
