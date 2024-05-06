package relationships

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// TestAddRelationship tests the AddRelationship function
func TestAddRelationship(t *testing.T) {
	db, mock, err := sqlmock.New() // Create a new instance of sqlmock
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rel := Relationship{
		Person1ID:        1,
		Person2ID:        2,
		RelationshipType: "sibling",
	}

	// Setup expectations
	mock.ExpectExec("INSERT INTO relationships").
		WithArgs(rel.Person1ID, rel.Person2ID, rel.RelationshipType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function with the mock database
	err = AddRelationship(db, rel)
	assert.NoError(t, err)

	// Make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestGetRelationships tests the GetRelationships function
func TestGetRelationships(t *testing.T) {
	db, mock, err := sqlmock.New() // Create a new mock database
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"person1_id", "person2_id", "relationship_type"}
	// Expect a query to fetch data, set up mock rows
	mock.ExpectQuery("SELECT person1_id, person2_id, relationship_type FROM relationships").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, 2, "sibling").AddRow(1, 3, "parent"))

	relationships, err := GetRelationships(db, 1)
	assert.NoError(t, err)
	assert.Len(t, relationships, 2) // Expecting two relationships
	assert.Equal(t, "sibling", relationships[0].RelationshipType)
	assert.Equal(t, "parent", relationships[1].RelationshipType)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
