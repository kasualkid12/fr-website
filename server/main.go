package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kasualkid12/fr-website/server/handlers/miniohandlers"
	"github.com/kasualkid12/fr-website/server/handlers/personhandlers"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	"github.com/kasualkid12/fr-website/server/modules/metrics"
	miniomodule "github.com/kasualkid12/fr-website/server/modules/minioModule"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// // make new bucket to test
	// bucketName := "test-bucket"
	// location := "us-east-1"

	// err = minioClient.MakeBucket(bucketName, location)
	// if err != nil {
	// 	exists, errBucketExists := minioClient.BucketExists(bucketName)
	// 	if errBucketExists == nil && exists {
	// 		log.Printf("We already own %s\n", bucketName)
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	// 	log.Printf("Successfully created %s\n", bucketName)
	// }

	// // Upload image to bucket
	// objectName := "test.jpg"
	// filePath := "a-file.jpg"
	// // contentType := "image/jpeg"

	// if err := minioClient.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{}); err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("Successfully saved %s\n", objectName)

	host, port, user, password, dbName, minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL := grabenv.GrabEnv()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Successfully connected!")

	// Connect to MinIO
	minioClient, err := minio.New(minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL)
	if err != nil {
		log.Fatal("Error connecting to MinIO:", err)
	}
	fmt.Println("Connected to MinIO")

	obj, err := miniomodule.GetObject(minioClient, "test-bucket", "Trey.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer obj.Close()

	// Download the object from MinIO to a local file
	// localFile, err := os.Create("/tmp/local-file.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer localFile.Close()

	// Add Prometheus metrics endpoint
	router.Handle("/metrics", metrics.MetricsHandler())

	// Handlers

	// Person handlers
	router.HandleFunc("/persons", personhandlers.GetPersonsHandler(db)).Methods("POST")

	// Minio handlers
	router.HandleFunc("/minio/addobject", miniohandlers.AddObjectHandler(minioClient)).Methods("POST")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "body"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := metrics.MetricsMiddleware(c.Handler(router))

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
