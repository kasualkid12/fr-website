package person

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Person struct to hold personal data
type Person struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	BirthDate string  `json:"birthDate"` // in "YYYY-MM-DD" format
	DeathDate *string `json:"deathDate"` // can be nil if person is alive
	Gender    string  `json:"gender"`
	PhotoURL  *string `json:"photoUrl"`  // can be nil if no photo URL is provided
	ProfileID *int    `json:"profileId"` // can be nil if no profile is linked
}

// AddPerson inserts a new person into the database
func AddPerson(db *sql.DB, p Person) error {
	query := `INSERT INTO persons (name, birth_date, death_date, gender, profile_id, photo_url) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, p.Name, p.BirthDate, p.DeathDate, p.Gender, p.ProfileID, p.PhotoURL)
	if err != nil {
		return fmt.Errorf("AddPersons insert error: %v", err)
	}
	return nil
}

// DeletePerson removes a person from the database by ID
func DeletePerson(db *sql.DB, personID int) error {
	query := `DELETE FROM persons WHERE person_id = $1`
	result, err := db.Exec(query, personID)
	if err != nil {
		return fmt.Errorf("DeletePerson delete error: %v", err)
	}

	// Optional: Check if the row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeletePerson rows affected error: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no person was deleted with the ID: %d", personID)
	}

	return nil
}

func GetPersons(db *sql.DB) ([]Person, error) {
	query := `
	WITH RECURSIVE descendents AS (
    SELECT DISTINCT person_id, name, birth_date, death_date, gender, photo_url, profile_id
    FROM persons
    WHERE person_id = 1
    
    UNION
    
    SELECT DISTINCT p.person_id, p.name, p.birth_date, p.death_date, p.gender, p.photo_url, p.profile_id
    FROM persons p
    JOIN relationships r ON p.person_id = r.person2_id
    JOIN descendents d ON r.person1_id = d.person_id
    WHERE r.relationship_type IN ('Parent', 'Spouse')
)
SELECT * FROM descendents;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, fmt.Errorf("GetPersons rows error: %v", err)
	}
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.Name, &person.BirthDate, &person.DeathDate, &person.Gender, &person.PhotoURL, &person.ProfileID)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("GetPersons scan error: %v", err)
		}
		persons = append(persons, person)
	}

	return persons, nil
}
