package person

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	p := Person{
		Name:      "John Doe",
		BirthDate: "1990-01-01",
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

	columns := []string{"person_id", "name"}
	mock.ExpectQuery("WITH RECURSIVE descendents").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "John Doe"))

	rows, err := GetPersons(db)
	assert.NoError(t, err)

	var personID int
	var name string
	for rows.Next() {
		err := rows.Scan(&personID, &name)
		assert.NoError(t, err)
		assert.Equal(t, 1, personID)
		assert.Equal(t, "John Doe", name)
	}

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}