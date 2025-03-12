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

func FPutObjects(bucketName, objectName, filePath, minioUrl, accessKeyID, secreetAccessKey string) (string, error) {
	// Upload the zip file
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	_, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return "Success", nil
}

func FGetObjects(bucketName, objectName, filePath, minioUrl, accessKeyID, secreetAccessKey string) (string, error) {
	// Download the zip file
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	err := minioClient.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return "Success", nil
}

func ListObjects(bucketName, minioUrl, accessKeyID, secreetAccessKey, Prefix string) ([]minio.ObjectInfo, error) {
	// List all objects from the server
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	objectsCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Prefix:    Prefix,
		Recursive: true,
	})
	var objects []minio.ObjectInfo
	for object := range objectsCh {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		objects = append(objects, object)
	}

	return objects, nil
}
