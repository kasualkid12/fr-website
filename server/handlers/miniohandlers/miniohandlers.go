package miniohandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	miniomodule "github.com/kasualkid12/fr-website/server/modules/minioModule"
)

type requestBody struct {
	BucketName  string `json:"bucketName"`
	ObjectName  string `json:"objectName"`
	Location    string `json:"location"`
	FilePath    string `json:"filePath"`
	ContentType string `json:"contentType"`
}

// ----------------------BUCKET HANDLERS------------------------
// Make bucket
func MakeBucketHandler(minioClient miniomodule.MinioClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(reqBody)
		if err := miniomodule.MakeBucket(minioClient, reqBody.BucketName, reqBody.Location); err != nil {
			if err.Error() == "Bucket already exists" {
				w.WriteHeader(http.StatusOK)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Remove bucket
func RemoveBucketHandler(minioClient miniomodule.MinioClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := miniomodule.RemoveBucket(minioClient, reqBody.BucketName); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// ----------------------OBJECT HANDLERS------------------------
// Add object
func AddObjectHandler(minioClient miniomodule.MinioClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form, limiting in-memory data to 10MB.
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Retrieve text fields from the form data.
		bucketName := r.FormValue("bucketName")
		objectName := r.FormValue("objectName")
		contentType := r.FormValue("contentType")

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		if err := miniomodule.AddObject(minioClient, bucketName, objectName, file, fileHeader.Size, contentType); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Remove object
func RemoveObjectHandler(minioClient miniomodule.MinioClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := miniomodule.RemoveObject(minioClient, reqBody.BucketName, reqBody.ObjectName); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Get object
func GetObjectHandler(minioClient miniomodule.MinioClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		object, err := miniomodule.GetObject(minioClient, reqBody.BucketName, reqBody.ObjectName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Optionally, set appropriate headers (e.g., Content-Type)
		w.Header().Set("Content-Type", reqBody.ContentType)

		// Stream the object directly to the response
		if _, err := io.Copy(w, object); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
