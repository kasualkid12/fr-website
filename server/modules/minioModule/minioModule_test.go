package miniomodule

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
)

type mockMinioClient struct {
	minio.Client
	BucketExistsFunc func(string) (bool, error)
	MakeBucketFunc   func(string, string) error
	RemoveBucketFunc func(string) error
	PutObjectFunc    func(string, string, io.Reader, int64, minio.PutObjectOptions) (interface{}, error)
	RemoveObjectFunc func(string, string) error
	GetObjectFunc    func(string, string, minio.GetObjectOptions) (io.ReadCloser, error)
}

func (m *mockMinioClient) BucketExists(bucketName string) (bool, error) {
	return m.BucketExistsFunc(bucketName)
}
func (m *mockMinioClient) MakeBucket(bucketName, location string) error {
	return m.MakeBucketFunc(bucketName, location)
}
func (m *mockMinioClient) RemoveBucket(bucketName string) error {
	return m.RemoveBucketFunc(bucketName)
}
func (m *mockMinioClient) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (interface{}, error) {
	return m.PutObjectFunc(bucketName, objectName, reader, objectSize, opts)
}
func (m *mockMinioClient) RemoveObject(bucketName, objectName string) error {
	return m.RemoveObjectFunc(bucketName, objectName)
}
func (m *mockMinioClient) GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
	return m.GetObjectFunc(bucketName, objectName, opts)
}

func TestMakeBucket(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return false, nil },
		MakeBucketFunc:   func(name, loc string) error { return nil },
	}
	err := MakeBucket(client, "test-bucket", "us-east-1")
	assert.NoError(t, err)
}

func TestMakeBucket_AlreadyExists(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return true, nil },
	}
	err := MakeBucket(client, "test-bucket", "us-east-1")
	assert.NoError(t, err)
}

func TestMakeBucket_Error(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return false, nil },
		MakeBucketFunc:   func(name, loc string) error { return errors.New("fail") },
	}
	err := MakeBucket(client, "test-bucket", "us-east-1")
	assert.Error(t, err)
}

func TestRemoveBucket(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return true, nil },
		RemoveBucketFunc: func(name string) error { return nil },
	}
	err := RemoveBucket(client, "test-bucket")
	assert.NoError(t, err)
}

func TestRemoveBucket_NotExists(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return false, nil },
	}
	err := RemoveBucket(client, "test-bucket")
	assert.NoError(t, err)
}

func TestRemoveBucket_Error(t *testing.T) {
	client := &mockMinioClient{
		BucketExistsFunc: func(name string) (bool, error) { return true, nil },
		RemoveBucketFunc: func(name string) error { return errors.New("fail") },
	}
	err := RemoveBucket(client, "test-bucket")
	assert.Error(t, err)
}

func TestAddObject(t *testing.T) {
	client := &mockMinioClient{
		PutObjectFunc: func(bucket, object string, r io.Reader, size int64, opts minio.PutObjectOptions) (interface{}, error) {
			return nil, nil
		},
	}
	err := AddObject(client, "test-bucket", "obj", strings.NewReader("data"), 4, "image/jpeg")
	assert.NoError(t, err)
}

func TestAddObject_Error(t *testing.T) {
	client := &mockMinioClient{
		PutObjectFunc: func(bucket, object string, r io.Reader, size int64, opts minio.PutObjectOptions) (interface{}, error) {
			return nil, errors.New("fail")
		},
	}
	err := AddObject(client, "test-bucket", "obj", strings.NewReader("data"), 4, "image/jpeg")
	assert.Error(t, err)
}

func TestRemoveObject(t *testing.T) {
	client := &mockMinioClient{
		RemoveObjectFunc: func(bucket, object string) error { return nil },
	}
	err := RemoveObject(client, "test-bucket", "obj")
	assert.NoError(t, err)
}

func TestRemoveObject_Error(t *testing.T) {
	client := &mockMinioClient{
		RemoveObjectFunc: func(bucket, object string) error { return errors.New("fail") },
	}
	err := RemoveObject(client, "test-bucket", "obj")
	assert.Error(t, err)
}

func TestGetObject(t *testing.T) {
	client := &mockMinioClient{
		GetObjectFunc: func(bucket, object string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
			return ioutil.NopCloser(strings.NewReader("mockdata")), nil
		},
	}
	obj, err := GetObject(client, "test-bucket", "obj")
	assert.NoError(t, err)
	assert.NotNil(t, obj)
}

func TestGetObject_Error(t *testing.T) {
	client := &mockMinioClient{
		GetObjectFunc: func(bucket, object string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
			return nil, errors.New("fail")
		},
	}
	obj, err := GetObject(client, "test-bucket", "obj")
	assert.Error(t, err)
	assert.Nil(t, obj)
}
