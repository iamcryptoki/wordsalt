package wordsalt

import (
	"fmt"
	"strings"
	"testing"
)

// TestGenerateWordPressKeys tests that GenerateWordPressKeys
// produces the expected number of keys and valid key values.
func TestGenerateWordPressKeys(t *testing.T) {
	keys := GenerateWordPressKeys()

	// Verify that the number of keys is correct (should be equal to len(keyNames)).
	if len(keys) != len(keyNames) {
		t.Errorf("Expected %d keys, got %d", len(keyNames), len(keys))
	}

	// Verify that each key is present and non-empty.
	for _, name := range keyNames {
		key, exists := keys[name]
		if !exists {
			t.Errorf("Expected key %s to be generated, but it is missing", name)
		}

		if len(key) == 0 {
			t.Errorf("Key %s is empty, expected a non-empty value", name)
		}

		// Verify that the key contains only valid characters.
		for _, char := range key {
			if !strings.Contains(chars, string(char)) {
				t.Errorf("Key %s contains invalid character: %c", name, char)
			}
		}
	}

	// Verify the uniqueness of the keys.
	seenKeys := make(map[string]bool)
	for name, key := range keys {
		if seenKeys[key] {
			t.Errorf("Duplicate key found for %s: %s", name, key)
		}

		seenKeys[key] = true
	}
}

// TestGenerateKeyLength tests that the key generated has the correct length.
func TestGenerateKeyLength(t *testing.T) {
	keys := GenerateWordPressKeys()

	// Check that each key has the expected length.
	for _, name := range keyNames {
		key := keys[name]
		if len(key) != keyLen {
			t.Errorf("Expected key %s to have length %d, but got length %d", name, keyLen, len(key))
		}
	}
}

// TestGenerateKeyCharacterSet tests that all characters in the generated key
// are from the allowed character set.
func TestGenerateKeyCharacterSet(t *testing.T) {
	keys := GenerateWordPressKeys()

	// Verify that each character in each key is from the allowed character set.
	for _, name := range keyNames {
		key := keys[name]
		for _, char := range key {
			if !strings.Contains(chars, string(char)) {
				t.Errorf("Key %s contains invalid character: %c", name, char)
			}
		}
	}
}

// TestGenerateKeyInvalidLength tests that generateKey
// returns an error when passed an invalid length.
func TestGenerateKeyInvalidLength(t *testing.T) {
	_, err := generateKey(0) // Passing 0 should result in an error.
	if err == nil {
		t.Errorf("Expected error for key length 0, but got nil")
	}

	_, err = generateKey(-1) // Negative length should also result in an error.
	if err == nil {
		t.Errorf("Expected error for negative key length, but got nil")
	}
}

// TestGenerateKeyRandomness tests that generateKey
// generates unique keys for different calls.
func TestGenerateKeyRandomness(t *testing.T) {
	key1, err1 := generateKey(64)
	key2, err2 := generateKey(64)

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error in key generation: %v, %v", err1, err2)
	}

	// Verify that the keys are different.
	if key1 == key2 {
		t.Errorf("Expected different keys, but got identical keys: %s", key1)
	}
}

// TestGenerateKeyPanicOnError tests that GenerateWordPressKeys
// panics when generateKey returns an error.
func TestGenerateKeyPanicOnError(t *testing.T) {
	// Save the original generateKeyFunc so we can restore it later.
	originalGenerateKeyFunc := generateKeyFunc
	defer func() { generateKeyFunc = originalGenerateKeyFunc }()

	// Simulate a failure in generating a key (force an error).
	generateKeyFunc = func(length int) (string, error) {
		return "", fmt.Errorf("simulated error in key generation")
	}

	// Use defer and recover to check for panic.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to error in generateKey, but got none")
		}
	}()

	// Call GenerateWordPressKeys, which will panic due to the error in generateKey.
	GenerateWordPressKeys()
}
