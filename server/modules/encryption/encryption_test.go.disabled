package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptAndDecrypt(t *testing.T) {
	// Test data
	originalData := "hello world"

	// Encrypt data
	encryptedData, err := EncryptWithKMS(originalData)
	assert.NoError(t, err, "EncryptWithKMS should not return an error")

	// Decrypt data
	decryptedData, err := DecryptWithKMS(encryptedData)
	assert.NoError(t, err, "DecryptWithKMS should not return an error")

	// Check if decrypted data matches original data
	assert.Equal(t, originalData, decryptedData, "Decrypted data should match original data")
}

func TestEncryptionFailure(t *testing.T) {
	// Test data
	invalidData := ""

	// Encrypt invalid data
	_, err := EncryptWithKMS(invalidData)
	assert.Error(t, err, "EncryptWithKMS should return an error for invalid data")
	assert.Contains(t, err.Error(), "failed to encrypt data", "Error message should contain expected message")
}

func TestDecryptionFailure(t *testing.T) {
	// Test data
	invalidEncryptedData := "invalid encrypted data"

	// Decrypt invalid encrypted data
	_, err := DecryptWithKMS(invalidEncryptedData)
	assert.Error(t, err, "DecryptWithKMS should return an error for invalid encrypted data")
}
