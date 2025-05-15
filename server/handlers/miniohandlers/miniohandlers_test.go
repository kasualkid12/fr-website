package miniohandlers

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	minio "github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
)

type mockMinioClient struct {
	makeBucketFunc   func(string, string) error
	removeBucketFunc func(string) error
	addObjectFunc    func(string, string, io.Reader, int64, string) error
	removeObjectFunc func(string, string) error
	getObjectFunc    func(string, string, minio.GetObjectOptions) (io.ReadCloser, error)
}

func (m *mockMinioClient) BucketExists(bucketName string) (bool, error) { return true, nil }
func (m *mockMinioClient) MakeBucket(bucketName, location string) error {
	if m.makeBucketFunc != nil {
		return m.makeBucketFunc(bucketName, location)
	}
	return nil
}
func (m *mockMinioClient) RemoveBucket(bucketName string) error {
	if m.removeBucketFunc != nil {
		return m.removeBucketFunc(bucketName)
	}
	return nil
}
func (m *mockMinioClient) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (interface{}, error) {
	if m.addObjectFunc != nil {
		return nil, m.addObjectFunc(bucketName, objectName, reader, objectSize, opts.ContentType)
	}
	return nil, nil
}
func (m *mockMinioClient) RemoveObject(bucketName, objectName string) error {
	if m.removeObjectFunc != nil {
		return m.removeObjectFunc(bucketName, objectName)
	}
	return nil
}
func (m *mockMinioClient) GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
	if m.getObjectFunc != nil {
		return m.getObjectFunc(bucketName, objectName, opts)
	}
	return ioutil.NopCloser(strings.NewReader("testdata")), nil
}

func TestMakeBucketHandler_Success(t *testing.T) {
	mockClient := &mockMinioClient{
		makeBucketFunc: func(bucket, loc string) error { return nil },
	}
	h := MakeBucketHandler(mockClient)
	body := `{"bucketName":"test-bucket","location":"us-east-1"}`
	req := httptest.NewRequest("POST", "/minio/makebucket", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMakeBucketHandler_Error(t *testing.T) {
	mockClient := &mockMinioClient{
		makeBucketFunc: func(bucket, loc string) error { return errors.New("Bucket already exists") },
	}
	h := MakeBucketHandler(mockClient)
	body := `{"bucketName":"test-bucket","location":"us-east-1"}`
	req := httptest.NewRequest("POST", "/minio/makebucket", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRemoveBucketHandler_Success(t *testing.T) {
	mockClient := &mockMinioClient{
		removeBucketFunc: func(bucket string) error { return nil },
	}
	h := RemoveBucketHandler(mockClient)
	body := `{"bucketName":"test-bucket"}`
	req := httptest.NewRequest("DELETE", "/minio/removebucket", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRemoveBucketHandler_Error(t *testing.T) {
	mockClient := &mockMinioClient{
		removeBucketFunc: func(bucket string) error { return errors.New("fail") },
	}
	h := RemoveBucketHandler(mockClient)
	body := `{"bucketName":"test-bucket"}`
	req := httptest.NewRequest("DELETE", "/minio/removebucket", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestAddObjectHandler_Success(t *testing.T) {
	mockClient := &mockMinioClient{
		addObjectFunc: func(bucket, object string, r io.Reader, size int64, ct string) error { return nil },
	}
	h := AddObjectHandler(mockClient)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	_ = w.WriteField("bucketName", "test-bucket")
	_ = w.WriteField("objectName", "test.jpg")
	_ = w.WriteField("contentType", "image/jpeg")
	fw, _ := w.CreateFormFile("file", "test.jpg")
	fw.Write([]byte("imagedata"))
	w.Close()
	req := httptest.NewRequest("POST", "/minio/addobject", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAddObjectHandler_Error(t *testing.T) {
	mockClient := &mockMinioClient{
		addObjectFunc: func(bucket, object string, r io.Reader, size int64, ct string) error { return errors.New("fail") },
	}
	h := AddObjectHandler(mockClient)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	_ = w.WriteField("bucketName", "test-bucket")
	_ = w.WriteField("objectName", "test.jpg")
	_ = w.WriteField("contentType", "image/jpeg")
	fw, _ := w.CreateFormFile("file", "test.jpg")
	fw.Write([]byte("imagedata"))
	w.Close()
	req := httptest.NewRequest("POST", "/minio/addobject", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestRemoveObjectHandler_Success(t *testing.T) {
	mockClient := &mockMinioClient{
		removeObjectFunc: func(bucket, object string) error { return nil },
	}
	h := RemoveObjectHandler(mockClient)
	body := `{"bucketName":"test-bucket","objectName":"test.jpg"}`
	req := httptest.NewRequest("DELETE", "/minio/removeobject", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRemoveObjectHandler_Error(t *testing.T) {
	mockClient := &mockMinioClient{
		removeObjectFunc: func(bucket, object string) error { return errors.New("fail") },
	}
	h := RemoveObjectHandler(mockClient)
	body := `{"bucketName":"test-bucket","objectName":"test.jpg"}`
	req := httptest.NewRequest("DELETE", "/minio/removeobject", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetObjectHandler_Success(t *testing.T) {
	mockClient := &mockMinioClient{
		getObjectFunc: func(bucket, object string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
			return ioutil.NopCloser(strings.NewReader("imagedata")), nil
		},
	}
	h := GetObjectHandler(mockClient)
	body := `{"bucketName":"test-bucket","objectName":"test.jpg","contentType":"image/jpeg"}`
	req := httptest.NewRequest("POST", "/minio/getobject", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))
	assert.Equal(t, "imagedata", rec.Body.String())
}

func TestGetObjectHandler_Error(t *testing.T) {
	mockClient := &mockMinioClient{
		getObjectFunc: func(bucket, object string, opts minio.GetObjectOptions) (io.ReadCloser, error) {
			return nil, errors.New("fail")
		},
	}
	h := GetObjectHandler(mockClient)
	body := `{"bucketName":"test-bucket","objectName":"test.jpg","contentType":"image/jpeg"}`
	req := httptest.NewRequest("POST", "/minio/getobject", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
