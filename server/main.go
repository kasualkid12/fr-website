package main

import (
	"database/sql"
	"fmt"
	"log"

	addperson "github.com/kasualkid12/fr-website/server/modules/addPerson"
	"github.com/kasualkid12/fr-website/server/modules/encryption"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	_ "github.com/lib/pq"
)

func main() {
	host, port, user, password, dbname, key := grabenv.GrabEnv()

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

	newPerson := addperson.Person{
		Name:      "John Doe",
		BirthDate: "1990-01-01",
		DeathDate: nil, // still alive
		Gender:    "male",
		ProfileID: nil, // no linked profile
	}

	encryptedText, err := encryption.Encrypt(newPerson.Name, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encryptedText)
	
	decryptedText, err := encryption.Decrypt(encryptedText, key)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(decryptedText)

	// err = addperson.AddPerson(db, newPerson)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Person added successfully!")

	fmt.Println("Successfully connected!")
}
