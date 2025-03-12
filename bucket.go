package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/tags"
)

// /Bucket Section
// /* This section is Bucket Section */
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

func BucketsExists(NameBucket, minioUrl, accessKeyID, secreetAccessKey string) (bool, error) {
	// Check if a bucket exists.
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	found, err := minioClient.BucketExists(context.Background(), NameBucket)
	if err != nil {
		log.Fatalln(err)
	}
	if found {
		fmt.Println("Bucket found")
	} else {
		fmt.Println("Bucket not found")
	}
	return found, nil
}

func RemoveBuckets(NameBucket, minioUrl, accessKeyID, secreetAccessKey string) (string, error) {
	// Remove a bucket called my-bucketname.
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	err := minioClient.RemoveBucket(context.Background(), NameBucket)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully removed", NameBucket)
	return "Success", nil
}

func SetBucketTagging(NameBucket, minioUrl, accessKeyID, secreetAccessKey string, NewTags map[string]string) (string, error) {
	tags, errTag := tags.NewTags(NewTags, false)
	if errTag != nil {
		log.Fatalln(errTag)
	}
	// Set bucket tagging
	minioClient := NewMinioAPi(minioUrl, accessKeyID, secreetAccessKey)
	err := minioClient.SetBucketTagging(context.Background(), NameBucket, tags)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully set bucket tagging on", NameBucket)
	return "Success", nil
}

///Object Section
/* This section is Object Section */
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
