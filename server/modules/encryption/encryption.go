package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	_ "github.com/lib/pq"
)

func Encrypt(data string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Encrypt cypher error: %v", err)
	}

	b := []byte(data)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("Encrypt read error: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedString string, key string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", fmt.Errorf("Decrypt decode error: %v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Decrypt cypher error: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("Decrypt block size error: %v", err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
