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

	awsAccessKeyID     string
	awsSecretAccessKey string
	awsRegion          string
	kmsEndpoint        string
	keyAlias           string
)

var loadEnv = godotenv.Load

func GrabEnv() (string, int, string, string, string, string, string, string, bool, string, string, string, string, string) {
	err := loadEnv("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return readEnv()
}

func readEnv() (string, int, string, string, string, string, string, string, bool, string, string, string, string, string) {
	host = os.Getenv("PGHOST")
	portStr := os.Getenv("PGPORT")
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbName = os.Getenv("PGDBNAME")
	minioEndpoint = os.Getenv("MINIOENDPOINT")
	minioAccessKeyID = os.Getenv("MINIOACCESSKEYID")
	minioSecretKey = os.Getenv("MINIOSECRETACCESSKEY")
	minioUseSSLStr := os.Getenv("MINIOUSESSL")

	awsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion = os.Getenv("AWS_REGION")
	if awsRegion == "" {
		awsRegion = os.Getenv("AWS_DEFAULT_REGION")
	}
	kmsEndpoint = os.Getenv("KMS_ENDPOINT")
	keyAlias = os.Getenv("KEY_ALIAS")

	var err error
	port, err = strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Error converting PGPORT to int: %v", err))
	}
	minioUseSSL, err = strconv.ParseBool(minioUseSSLStr)
	if err != nil {
		panic(fmt.Sprintf("Error converting minioUseSSL to bool: %v", err))
	}

	return host, port, user, password, dbName, minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL, awsAccessKeyID, awsSecretAccessKey, awsRegion, kmsEndpoint, keyAlias
}
