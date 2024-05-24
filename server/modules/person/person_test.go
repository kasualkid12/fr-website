package person

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kasualkid12/fr-website/server/modules/customdate"
	"github.com/stretchr/testify/assert"
)

func TestAddPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Setup our expected actions
	mock.ExpectExec("INSERT INTO persons").
		WithArgs("John Doe", "1990-01-01", nil, "male", nil, nil).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Mock the result: 1 row affected

	birthDate := customdate.CustomDate{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)}

	p := Person{
		Name:      "John Doe",
		BirthDate: birthDate,
		DeathDate: nil,
		Gender:    "male",
		ProfileID: nil,
		PhotoURL:  nil,
	}

	err = AddPerson(db, p)
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup expectations
	personID := 1
	mock.ExpectExec("DELETE FROM persons WHERE person_id =").
		WithArgs(personID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Mock the result: 1 row affected

	// Call the function with the mock database
	err = DeletePerson(db, personID)
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletePersonNoRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup expectations for no rows being deleted
	personID := 99
	mock.ExpectExec("DELETE FROM persons WHERE person_id =").
		WithArgs(personID).
		WillReturnResult(sqlmock.NewResult(0, 0)) // Mock the result: no rows affected

	// Call the function with the mock database
	err = DeletePerson(db, personID)
	assert.Error(t, err, "expected an error for no rows affected")
	assert.Contains(t, err.Error(), "no person was deleted")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPersons(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"person_id", "name", "birth_date", "death_date", "gender", "photo_url", "profile_id", "relationship", "parent_object"}
	mock.ExpectQuery("SELECT \\* FROM TREE_CHILD_SPOUSE_VW WHERE parent_object =").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, "John Doe", "2000-01-01", sql.NullTime{Valid: true, Time: time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC)}, "Male", "http://example.com/photo.jpg", 101, "son", 0))

	persons, err := GetPersons(db, 1)
	assert.NoError(t, err)
	assert.Len(t, persons, 1)

	birthDate := customdate.CustomDate{Time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)}
	deathDate := customdate.CustomDate{Time: time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC)}
	photoURL := "http://example.com/photo.jpg"
	profileID := 101

	expectedPerson := Person{
		ID:           1,
		Name:         "John Doe",
		BirthDate:    birthDate,
		DeathDate:    &deathDate,
		Gender:       "Male",
		PhotoURL:     &photoURL,
		ProfileID:    &profileID,
		Relationship: "son",
		ParentObject: 0,
	}

	assert.Equal(t, expectedPerson, persons[0])

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
