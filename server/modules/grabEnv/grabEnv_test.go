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
	_ = os.Setenv("MINIOENDPOINT", "testendpoint")
	_ = os.Setenv("MINIOACCESSKEYID", "testkeyid")
	_ = os.Setenv("MINIOSECRETACCESSKEY", "testsecret")
	_ = os.Setenv("MINIOUSESSL", "true")

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "testaccesskey")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "testsecretkey")
	_ = os.Setenv("AWS_REGION", "us-east-1")
	_ = os.Setenv("KMS_ENDPOINT", "http://localhost:4566")
	_ = os.Setenv("KEY_ALIAS", "alias/test-key")

	host, port, user, password, dbname, minioEndpoint, minioAccessKeyID, minioSecretKey, minioUseSSL, awsAccessKeyID, awsSecretAccessKey, awsRegion, kmsEndpoint, keyAlias := GrabEnv()

	assert.Equal(t, "localhost", host)
	assert.Equal(t, 5432, port)
	assert.Equal(t, "testuser", user)
	assert.Equal(t, "testpass", password)
	assert.Equal(t, "testdb", dbname)
	assert.Equal(t, "testendpoint", minioEndpoint)
	assert.Equal(t, "testkeyid", minioAccessKeyID)
	assert.Equal(t, "testsecret", minioSecretKey)
	assert.Equal(t, true, minioUseSSL)
	assert.Equal(t, "testaccesskey", awsAccessKeyID)
	assert.Equal(t, "testsecretkey", awsSecretAccessKey)
	assert.Equal(t, "us-east-1", awsRegion)
	assert.Equal(t, "http://localhost:4566", kmsEndpoint)
	assert.Equal(t, "alias/test-key", keyAlias)
}

func TestGrabEnvFailure(t *testing.T) {
	// Mock the godotenv.Load function
	loadEnv = func(filenames ...string) error {
		return nil // Simulate no error without loading actual files
	}
	defer func() { loadEnv = godotenv.Load }() // Reset after the test

	// Clearing the PGPORT to simulate a missing or invalid environment variable
	_ = os.Setenv("PGPORT", "notanumber")
	_ = os.Setenv("MINIOENDPOINT", "testendpoint")
	_ = os.Setenv("MINIOACCESSKEYID", "testkeyid")
	_ = os.Setenv("MINIOSECRETACCESSKEY", "testsecret")
	_ = os.Setenv("MINIOUSESSL", "true")

	// Running the function and capturing the log output
	assert.PanicsWithValue(t, "Error converting PGPORT to int: strconv.Atoi: parsing \"notanumber\": invalid syntax", func() {
		_, _, _, _, _, _, _, _, _, _, _, _, _, _ = GrabEnv()
	}, "Expected panic for invalid PGPORT")

	// Resetting PGPORT to a valid value for other tests
	_ = os.Setenv("PGPORT", "5432")
}
