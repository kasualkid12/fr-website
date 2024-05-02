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
