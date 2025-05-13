package grabenv

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	host             string
	port             int
	user             string
	password         string
	dbName           string
	minioEndpoint    string
	minioAccessKeyID string
	minioSecretKey   string
	minioUseSSL      bool
)

var loadEnv = godotenv.Load

func GrabEnv() (string, int, string, string, string, string, string, string, bool) {
	err := loadEnv("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return readEnv()
}

func readEnv() (string, int, string, string, string, string, string, string, bool) {
	host = os.Getenv("PGHOST")
	portStr := os.Getenv("PGPORT")
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbName = os.Getenv("PGDBNAME")
	minioEndpoint = os.Getenv("MINIOENDPOINT")
	minioAccessKeyID = os.Getenv("MINIOACCESSKEYID")
	minioSecretKey = os.Getenv("MINIOSECRETACCESSKEY")
	minioUseSSLStr := os.Getenv("MINIOUSESSL")

	var err error
	port, err = strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Error converting PGPORT to int: %v", err))
	}
	minioUseSSL, err = strconv.ParseBool(minioUseSSLStr)
	if err != nil {
		panic(fmt.Sprintf("Error converting minioUseSSL to bool: %v", err))
	}

	return host, port, user, password, dbName, minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL
}
