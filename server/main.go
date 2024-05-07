package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kasualkid12/fr-website/server/modules/encryption"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	_ "github.com/lib/pq"
)

func main() {
	host, port, user, password, dbname := grabenv.GrabEnv()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	originalText := "Hello, world!"

	encryptedText, err := encryption.EncryptWithKMS(originalText)
	if err != nil {
		log.Fatal("Encryption error:", err)
	}
	fmt.Println("Encrypted:", encryptedText)

	decryptedText, err := encryption.DecryptWithKMS(encryptedText)
	if err != nil {
		log.Fatal("Decryption error:", err)
	}
	fmt.Println("Decrypted:", decryptedText)

}
