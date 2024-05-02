package grabenv

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
	key      string
)

var loadEnv = godotenv.Load

func GrabEnv() (string, int, string, string, string, string) {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return readEnv()
}

func readEnv() (string, int, string, string, string, string) {
	host = os.Getenv("PGHOST")
	portStr := os.Getenv("PGPORT")
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname = os.Getenv("PGDBNAME")
	key = os.Getenv("KEY")

	var err error
	port, err = strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Error converting PGPORT to int: %v", err))
	}

	return host, port, user, password, dbname, key
}
