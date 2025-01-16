package miniomodule

import (
	"fmt"

	"github.com/minio/minio-go"
)

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