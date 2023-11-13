// cmd/cryptoFunctions.go

package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Helper function to hash the passphrase to a 32-byte key using SHA-256
func hashTo32Bytes(input string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hasher.Sum(nil) // Returns a 32-byte hash
}

// Encrypt takes a byte slice and a passphrase, then returns a base64 encoded string of the encrypted data.
func Encrypt(data []byte) (string, error) {
	passphrase := os.Getenv("WS1_KEY")
	if passphrase == "" {
		panic("WS1_KEY environment variable not set")
	}
	key := hashTo32Bytes(passphrase) // Hash the passphrase to get a 32-byte key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	encrypted := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptValues is a helper function that decrypts a base64-encoded string
func DecryptValues(encryptedBase64Value string) (string, error) {
	if encryptedBase64Value == "" {
		return "", nil // If the value is empty, return an empty string without error
	}
	decryptedBytes, err := decrypt(encryptedBase64Value)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}
	return string(decryptedBytes), nil
}

// decrypt takes a base64 encoded string, decodes, and then decrypts it to return the original data.
func decrypt(encryptedBase64Str string) ([]byte, error) {
	passphrase := os.Getenv("WS1_KEY")
	if passphrase == "" {
		return nil, fmt.Errorf("WS1_KEY environment variable not set")
	}
	key := hashTo32Bytes(passphrase) // Hash the passphrase to get a 32-byte key

	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedBase64Str)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertext := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
