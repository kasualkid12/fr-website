package grabenv

import (
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
)

func GrabEnv() (string, int, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host = os.Getenv("PGHOST")
	portStr := os.Getenv("PGPORT")
	user = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname = os.Getenv("PGDBNAME")

	var errConv error
	port, errConv = strconv.Atoi(portStr)
	if errConv != nil {
		log.Fatalf("Error converting PGPORT to int: %v", errConv)
	}

	return host, port, user, password, dbname
}
