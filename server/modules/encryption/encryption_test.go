package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "thiskeyis32byteslongandusedhere!"
	data := "Hello, world!"

	// Test encryption
	encryptedData, err := Encrypt(data, key)
	assert.NoError(t, err, "Encryption should not produce an error")
	assert.NotEmpty(t, encryptedData, "Encrypted data should not be empty")

	// Test decryption
	decryptedData, err := Decrypt(encryptedData, key)
	assert.NoError(t, err, "Decryption should not produce an error")
	assert.Equal(t, data, decryptedData, "Decrypted data should match original")
}

func TestEncryptError(t *testing.T) {
	data := "data to encrypt"
	shortKey := "short"

	// Test encryption with a short key
	_, err := Encrypt(data, shortKey)
	assert.Error(t, err, "Encryption should fail with a short key")
}

func TestDecryptError(t *testing.T) {
	key := "thiskeyis32byteslongandusedhere!"
	invalidEncryptedData := "not really encrypted data"

	// Test decryption with invalid data
	_, err := Decrypt(invalidEncryptedData, key)
	assert.Error(t, err, "Decryption should fail with invalid data")
}

