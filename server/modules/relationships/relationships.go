package relationships

import (
	"database/sql"
	"fmt"
)

type Relationship struct {
	Person1ID        int
	Person2ID        int
	RelationshipType string
}

// AddRelationship inserts a new relationship into the database
func AddRelationship(db *sql.DB, rel Relationship) error {
	query := `INSERT INTO relationships (person1_id, person2_id, relationship_type) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, rel.Person1ID, rel.Person2ID, rel.RelationshipType)
	if err != nil {
		return fmt.Errorf("AddRelationship error: %v", err)
	}
	return nil
}

// GetRelationships retrieves relationships for a given person ID
func GetRelationships(db *sql.DB, personID int) ([]Relationship, error) {
	query := `SELECT person1_id, person2_id, relationship_type FROM relationships WHERE person1_id = $1 OR person2_id = $1`
	rows, err := db.Query(query, personID)
	if err != nil {
		return nil, fmt.Errorf("GetRelationships query error: %v", err)
	}
	defer rows.Close()

	var relationships []Relationship
	for rows.Next() {
		var r Relationship
		err := rows.Scan(&r.Person1ID, &r.Person2ID, &r.RelationshipType)
		if err != nil {
			return nil, fmt.Errorf("GetRelationships scan error: %v", err)
		}
		relationships = append(relationships, r)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetRelationships rows error: %v", err)
	}
	return relationships, nil
}
