package miniomodule

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
)

// ----------------------BUCKET METHODS------------------------
// Check if bucket exists
func checkBucket(minioClient *minio.Client, bucketName string) (bool, error) {
	exists, err := minioClient.BucketExists(bucketName)
	if err != nil {
		log.Fatal("Error checking bucket: ", err)
		return exists, err
	}
	if !exists {
		fmt.Println("Bucket does not exist: ", bucketName)
		return exists, nil
	}
	return exists, nil
}

// Make bucket
func MakeBucket(minioClient *minio.Client, bucketName string, location string) error {
	exists, err := checkBucket(minioClient, bucketName)
	if err != nil {
		log.Fatal("Error checking bucket: ", err)
	}
	if exists {
		fmt.Println("Bucket already exists: ", bucketName)
		return nil
	} else if err := minioClient.MakeBucket(bucketName, location); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created bucket: ", bucketName)
	return nil
}

// Remove bucket
func RemoveBucket(minioClient *minio.Client, bucketName string) error {
	exists, err := checkBucket(minioClient, bucketName)
	if err != nil {
		log.Fatal("Error checking bucket: ", err)
	}
	if !exists {
		fmt.Println("Bucket does not exist: ", bucketName)
		return nil
	} else if err := minioClient.RemoveBucket(bucketName); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully removed bucket: ", bucketName)
	return nil
}

// ----------------------OBJECT METHODS------------------------
// Add object
func AddObject(minioClient *minio.Client, bucketName string, objectName string, file io.Reader, fileSize int64, contentType string) error {
	uploadInfo, err := minioClient.PutObject(bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		fmt.Println("Error uploading object: ", err)
		return err
	}
	fmt.Fprintf(os.Stdout, "Successfully uploaded object %s of size %d\n", objectName, uploadInfo)
	return nil
}

// Remove object
func RemoveObject(minioClient *minio.Client, bucketName string, objectName string) error {
	if err := minioClient.RemoveObject(bucketName, objectName); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully removed object: ", objectName)
	return nil
}

// Get object
func GetObject(minioClient *minio.Client, bucketName string, objectName string) (*minio.Object, error) {
	object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully got object: ", objectName)
	return object, nil
}
