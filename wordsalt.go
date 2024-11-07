package wordsalt

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	// keyLen defines the length of the generated keys.
	keyLen = 64
	// chars defines the set of characters allowed in the keys.
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|"
)

var (
	// keyNames holds the names of the WordPress keys to generate.
	keyNames = []string{
		"AUTH_KEY",
		"SECURE_AUTH_KEY",
		"LOGGED_IN_KEY",
		"NONCE_KEY",
		"AUTH_SALT",
		"SECURE_AUTH_SALT",
		"LOGGED_IN_SALT",
		"NONCE_SALT",
	}

	// generateKeyFunc is a variable that holds the actual key generation function.
	// This can be modified for testing.
	generateKeyFunc = generateKey
)

// GenerateWordPressKeys generates a map of WordPress security keys.
func GenerateWordPressKeys() map[string]string {
	keys := make(map[string]string)

	for _, name := range keyNames {
		key, err := generateKeyFunc(keyLen)
		if err != nil {
			panic(fmt.Errorf("failed to generate key %s: %w", name, err))
		}

		keys[name] = key
	}

	return keys
}

// generateKey generates a random key with the specified length and allowed characters.
func generateKey(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid key length: %d", length)
	}

	key := make([]byte, length)
	for i := range key {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}

		key[i] = chars[n.Int64()]
	}

	return string(key), nil
}
