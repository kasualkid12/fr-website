package main

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/kasualkid12/fr-website/server/modules/encryption"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	"github.com/kasualkid12/fr-website/server/modules/person"
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

	rows, err := person.GetPersons(db)
	if err != nil {
		log.Fatalf("Error getting persons: %s", err)
	}

	if !rows.Next() {
		fmt.Println("No rows returned from the query.")
	} else {
		for {
			var new_person person.Person
			var personID int
			if err := rows.Scan(&personID, &new_person.Name, &new_person.BirthDate, &new_person.DeathDate, &new_person.Gender, &new_person.PhotoURL, &new_person.ProfileID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Person ID: %d, Name: %s, Birth Date: %s, Death Date %s, Gender: %s, Photo URL: %s, Profile ID: %d\n", personID, new_person.Name, new_person.BirthDate, new_person.DeathDate, new_person.Gender, new_person.PhotoURL, new_person.ProfileID)

			// Check if there are more rows; if not, break the loop
			if !rows.Next() {
				break
			}
		}
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
