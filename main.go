package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

}

func NewMinioAPi(urlMinio, accessKeyId, secreetAccessKey string) *minio.Client {
	endpoint := urlMinio
	accessKeyID := accessKeyId
	secretAccessKey := secreetAccessKey
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient // minioClient is now set up
}
