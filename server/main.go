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
	// Create router
	router := mux.NewRouter()

	// Read environment variables
	host, port, user, password, dbName, minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL, _, _, _, _, _ := grabenv.GrabEnv()

	//---------------------DATABASE CONNECTIONS------------------------
	// Connect to PostgreSQL
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
	fmt.Println("Successfully connected to PostgreSQL")

	// Connect to MinIO
	minioClient, err := minio.New(minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL)
	if err != nil {
		log.Fatal("Error connecting to MinIO:", err)
	}
	fmt.Println("Connected to MinIO")

	// Wrap minioClient with RealMinioClient to satisfy MinioClient interface
	realMinioClient := &miniomodule.RealMinioClient{Client: minioClient}

	//---------------------METRICS------------------------
	// Add Prometheus metrics endpoint
	router.Handle("/metrics", metrics.MetricsHandler())

	//---------------------HANDLERS------------------------
	// Person handlers
	router.HandleFunc("/persons", personhandlers.GetPersonsHandler(db)).Methods("POST")
	router.HandleFunc("/persons/update", personhandlers.UpdatePersonHandler(db)).Methods("POST")

	// Minio handlers
	router.HandleFunc("/minio/makebucket", miniohandlers.MakeBucketHandler(realMinioClient)).Methods("POST")
	router.HandleFunc("/minio/removebucket", miniohandlers.RemoveBucketHandler(realMinioClient)).Methods("DELETE")
	router.HandleFunc("/minio/addobject", miniohandlers.AddObjectHandler(realMinioClient)).Methods("POST")
	router.HandleFunc("/minio/removeobject", miniohandlers.RemoveObjectHandler(realMinioClient)).Methods("DELETE")
	router.HandleFunc("/minio/getobject", miniohandlers.GetObjectHandler(realMinioClient)).Methods("POST")

	//---------------------SERVER STARTUP------------------------
	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "body"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := metrics.MetricsMiddleware(c.Handler(router))

	// Start server
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
