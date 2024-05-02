package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kasualkid12/fr-website/server/modules/relationships"
	// "github.com/kasualkid12/fr-website/server/modules/encryption"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	_ "github.com/lib/pq"
)

func main() {
	host, port, user, password, dbname, _ := grabenv.GrabEnv()

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

	// rel := relationships.Relationship{
	// 	Person1ID:        1,
	// 	Person2ID:        2,
	// 	RelationshipType: "sibling",
	// }

	// err = relationships.AddRelationship(db, rel)
	// if err != nil {
	// 	log.Fatal("Error adding relationship:", err)
	// }
	// fmt.Println("Relationship added successfully!")

	rels, err := relationships.GetRelationships(db, 1)
	if err != nil {
		log.Fatal("Error getting relationships:", err)
	}
	for _, rel := range rels {
		fmt.Printf("Relationship: %d -> %d (%s)\n", rel.Person1ID, rel.Person2ID, rel.RelationshipType)
	}

	// newPerson := person.Person{
	// 	Name:      "Jeff Doe",
	// 	BirthDate: "1995-01-01",
	// 	DeathDate: nil, // still alive
	// 	Gender:    "male",
	// 	ProfileID: nil, // no linked profile
	// }

	// encryptedText, err := encryption.Encrypt(newPerson.Name, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(encryptedText)

	// decryptedText, err := encryption.Decrypt(encryptedText, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(decryptedText)

	// err = person.AddPerson(db, newPerson)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Person added successfully!")

}
