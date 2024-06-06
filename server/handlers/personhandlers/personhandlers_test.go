package personhandlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kasualkid12/fr-website/server/modules/customdate"
	"github.com/kasualkid12/fr-website/server/modules/person"
	"github.com/stretchr/testify/assert"
)

func TestGetPersonsHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock the database response
	columns := []string{"person_id", "name", "birth_date", "death_date", "gender", "photo_url", "profile_id", "relationship", "parent_object"}
	mock.ExpectQuery("SELECT \\* FROM TREE_CHILD_SPOUSE_VW WHERE parent_object =").
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, "John Doe", "2000-01-01", sql.NullTime{Valid: true, Time: time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC)}, "Male", "http://example.com/photo.jpg", 101, "son", 3))

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/persons", nil)
	assert.NoError(t, err)

	// Record the response
	rr := httptest.NewRecorder()
	handler := GetPersonsHandler(db)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect
	expectedPersons := []person.Person{
		{
			ID:           1,
			FirstName:    "John",
			LastName:     "Doe",
			BirthDate:    customdate.CustomDate{Time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
			DeathDate:    &customdate.CustomDate{Time: time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC)},
			Gender:       "Male",
			PhotoURL:     stringPtr("http://example.com/photo.jpg"),
			ProfileID:    intPtr(101),
			Relationship: "son",
			ParentObject: 3,
		},
	}
	expectedBody, err := json.Marshal(expectedPersons)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedBody), rr.Body.String())

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
