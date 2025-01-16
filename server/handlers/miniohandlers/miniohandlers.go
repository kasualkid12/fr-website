package miniohandlers

import (
	"encoding/json"
	"net/http"

	miniomodule "github.com/kasualkid12/fr-website/server/modules/minioModule"
	"github.com/minio/minio-go"
)

type requestBody struct {
	bucketName  string `json:"bucketName"`
	objectName  string `json:"objectName"`
	filePath    string `json:"filePath"`
	contentType string `json:"contentType"`
}

func AddObjectHandler(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := miniomodule.AddObject(minioClient, reqBody.bucketName, reqBody.objectName, reqBody.filePath, reqBody.contentType); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
