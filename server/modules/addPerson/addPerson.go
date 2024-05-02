package addperson

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Person struct to hold personal data
type Person struct {
	Name      string
	BirthDate string  // in "YYYY-MM-DD" format
	DeathDate *string // can be nil if person is alive
	Gender    string
	ProfileID *int    // can be nil if no profile is linked
	PhotoURL  *string // can be nil if no photo URL is provided
}

// AddPerson inserts a new person into the database
func AddPerson(db *sql.DB, p Person) error {
	query := `INSERT INTO persons (name, birth_date, death_date, gender, profile_id, photo_url) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, p.Name, p.BirthDate, p.DeathDate, p.Gender, p.ProfileID, p.PhotoURL)
	if err != nil {
		return err
	}
	return nil
}
