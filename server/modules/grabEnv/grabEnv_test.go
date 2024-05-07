package grabenv

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGrabEnv(t *testing.T) {
	// Mock the godotenv.Load function
	loadEnv = func(filenames ...string) error {
		return nil // Simulate no error without loading actual files
	}
	defer func() { loadEnv = godotenv.Load }() // Reset after the test

	// Set environment variables
	_ = os.Setenv("PGHOST", "localhost")
	_ = os.Setenv("PGPORT", "5432")
	_ = os.Setenv("PGUSER", "testuser")
	_ = os.Setenv("PGPASSWORD", "testpass")
	_ = os.Setenv("PGDBNAME", "testdb")

	host, port, user, password, dbname := GrabEnv()

	assert.Equal(t, "localhost", host)
	assert.Equal(t, 5432, port)
	assert.Equal(t, "testuser", user)
	assert.Equal(t, "testpass", password)
	assert.Equal(t, "testdb", dbname)
}

func TestGrabEnvFailure(t *testing.T) {
	// Mock the godotenv.Load function
	loadEnv = func(filenames ...string) error {
		return nil // Simulate no error without loading actual files
	}
	defer func() { loadEnv = godotenv.Load }() // Reset after the test

	// Clearing the PGPORT to simulate a missing or invalid environment variable
	_ = os.Setenv("PGPORT", "notanumber")

	// Running the function and capturing the log output
	assert.PanicsWithValue(t, "Error converting PGPORT to int: strconv.Atoi: parsing \"notanumber\": invalid syntax", func() {
		_, _, _, _, _ = GrabEnv()
	}, "Expected panic for invalid PGPORT")

	// Resetting PGPORT to a valid value for other tests
	_ = os.Setenv("PGPORT", "5432")
}
