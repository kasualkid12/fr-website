package miniomodule

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
)

// Define an interface for the MinIO client to allow mocking in tests
// Only include the methods used in this module
type MinioClient interface {
	BucketExists(bucketName string) (bool, error)
	MakeBucket(bucketName, location string) error
	RemoveBucket(bucketName string) error
	PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (interface{}, error)
	RemoveObject(bucketName, objectName string) error
	GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.ReadCloser, error)
}

// RealMinioClient is a wrapper to adapt *minio.Client to the MinioClient interface
// This is needed because GetObject's return type is *minio.Object, which implements io.ReadCloser
// but the interface expects io.ReadCloser

type RealMinioClient struct {
	*minio.Client
}

func (r *RealMinioClient) BucketExists(bucketName string) (bool, error) {
	return r.Client.BucketExists(bucketName)
}
func (r *RealMinioClient) MakeBucket(bucketName, location string) error {
	return r.Client.MakeBucket(bucketName, location)
}
func (r *RealMinioClient) RemoveBucket(bucketName string) error {
	return r.Client.RemoveBucket(bucketName)
}
func (r *RealMinioClient) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (interface{}, error) {
	return r.Client.PutObject(bucketName, objectName, reader, objectSize, opts)
}
func (r *RealMinioClient) RemoveObject(bucketName, objectName string) error {
	return r.Client.RemoveObject(bucketName, objectName)
}
func (r *RealMinioClient) GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
	obj, err := r.Client.GetObject(bucketName, objectName, opts)
	return obj, err
}

// ----------------------BUCKET METHODS------------------------
// Check if bucket exists
func checkBucket(minioClient MinioClient, bucketName string) (bool, error) {
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
func MakeBucket(minioClient MinioClient, bucketName string, location string) error {
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
func RemoveBucket(minioClient MinioClient, bucketName string) error {
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
func AddObject(minioClient MinioClient, bucketName string, objectName string, file io.Reader, fileSize int64, contentType string) error {
	_, err := minioClient.PutObject(bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		fmt.Println("Error uploading object: ", err)
		return err
	}
	fmt.Fprintf(os.Stdout, "Successfully uploaded object %s\n", objectName)
	return nil
}

// Remove object
func RemoveObject(minioClient MinioClient, bucketName string, objectName string) error {
	if err := minioClient.RemoveObject(bucketName, objectName); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully removed object: ", objectName)
	return nil
}

// Get object
func GetObject(minioClient MinioClient, bucketName string, objectName string) (io.ReadCloser, error) {
	object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully got object: ", objectName)
	return object, nil
}
