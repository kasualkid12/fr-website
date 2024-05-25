package person

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/kasualkid12/fr-website/server/modules/customdate"
	_ "github.com/lib/pq"
)

// Person struct to hold personal data
type Person struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	BirthDate    customdate.CustomDate  `json:"birthDate"`           // in "YYYY-MM-DD" format
	DeathDate    *customdate.CustomDate `json:"deathDate,omitempty"` // can be nil if person is alive
	Gender       string                 `json:"gender"`
	PhotoURL     *string                `json:"photoUrl,omitempty"`  // can be nil if no photo URL is provided
	ProfileID    *int                   `json:"profileId,omitempty"` // can be nil if no profile is linked
	Relationship string                 `json:"relationship"`        // relationship to the root person
	ParentObject int                    `json:"parentObject"`
}

// AddPerson inserts a new person into the database
func AddPerson(db *sql.DB, p Person) error {
	query := `INSERT INTO persons (name, birth_date, death_date, gender, profile_id, photo_url) VALUES ($1, $2, $3, $4, $5, $6)`
	birthDate := p.BirthDate.Time.Format("2006-01-02")
	var deathDate *string
	if p.DeathDate != nil {
		d := p.DeathDate.Time.Format("2006-01-02")
		deathDate = &d
	}
	_, err := db.Exec(query, p.Name, birthDate, deathDate, p.Gender, p.ProfileID, p.PhotoURL)
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

func GetPersons(db *sql.DB, personID int) ([]Person, error) {
	query := `
	SELECT * FROM TREE_CHILD_SPOUSE_VW WHERE parent_object = $1;
	`

	rows, err := db.Query(query, personID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, fmt.Errorf("GetPersons rows error: %v", err)
	}
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var person Person
		var birthDate customdate.CustomDate
		var deathDateValid sql.NullTime

		err := rows.Scan(&person.ID, &person.Name, &birthDate, &deathDateValid, &person.Gender, &person.PhotoURL, &person.ProfileID, &person.Relationship, &person.ParentObject)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("GetPersons scan error: %v", err)
		}

		person.BirthDate = birthDate
		if deathDateValid.Valid {
			person.DeathDate = &customdate.CustomDate{Time: deathDateValid.Time}
		} else {
			person.DeathDate = nil
		}

		persons = append(persons, person)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %v", err)
		return nil, fmt.Errorf("GetPersons rows error: %v", err)
	}

	// Rearrange persons to move spouses above their respective persons
	persons = RearrangePersons(persons)

	return persons, nil
}

func RearrangePersons(persons []Person) []Person {
	personMap := make(map[string]int)
	for i, person := range persons {
		personMap[person.Name] = i
	}

	rearrange := make([]Person, 0, len(persons))
	rearranged := make([]Person, 0, len(persons))
	spouses := make(map[int]Person)

	for i, person := range persons {
		if strings.Contains(person.Relationship, "Spouse") {
			parentName := strings.TrimSuffix(person.Relationship, " Spouse")
			if index, exists := personMap[parentName]; exists {
				if i-1 == index {
					// Spouse is already in the correct spot
					rearrange = append(rearrange, person)
					continue
				}
				spouses[index] = person
			}
		} else {
			rearrange = append(rearrange, person)
		}
	}

	// Insert spouses above their respective persons
	for i := 0; i < len(rearrange); i++ {
		rearranged = append(rearranged, rearrange[i])
		if spouse, exists := spouses[i]; exists {
			rearranged = append(rearranged, spouse)
		}
	}

	return rearranged
}
