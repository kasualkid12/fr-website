package miniomodule

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

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

func AddObject(minioClient *minio.Client, bucketName string, objectName string, filePath string, contentType string) error {
	uploadInfo, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)
	return nil
}

func RemoveObject(minioClient *minio.Client, bucketName string, objectName string) error {
	if err := minioClient.RemoveObject(bucketName, objectName); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully removed object: ", objectName)
	return nil
}

func GetObject(minioClient *minio.Client, bucketName string, objectName string) (*minio.Object, error) {
	object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully got object: ", objectName)
	return object, nil
}
