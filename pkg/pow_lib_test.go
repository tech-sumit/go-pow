package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"testing"
)

func TestSolvePoW(t *testing.T) {
	nonce := "testnonce"
	solution, err := SolvePoW(nonce)
	if err != nil {
		t.Errorf("Error solving PoW: %v", err)
	}

	// Validate the solution
	hash := sha256.Sum256([]byte(nonce + solution))
	hexHash := hex.EncodeToString(hash[:])
	if !strings.HasPrefix(hexHash, strings.Repeat("0", difficulty)) {
		t.Error("PoW solution is incorrect")
	}
}

func TestGenerateNonce(t *testing.T) {
	nonce := GenerateNonce()
	if len(nonce) == 0 {
		t.Errorf("Generated nonce is empty")
	}
	if len(nonce) != 32 { // 16 bytes in hex is 32 characters
		t.Errorf("Expected nonce length of 32, got %d", len(nonce))
	}
}

func TestValidatePoW(t *testing.T) {
	nonce := "testnonce"

	// Directly create a hash that should satisfy the difficulty requirement
	for {
		testString := randStringBytes(10) // Random string for the hash
		hash := sha256.Sum256([]byte(nonce + testString))
		hexHash := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hexHash, strings.Repeat("0", difficulty)) {
			if !ValidatePoW(nonce, testString) {
				t.Errorf("validatePoW should have returned true for a valid hash")
			}
			break
		}
	}

	invalidHash := "invalidhashvalue"
	if ValidatePoW(nonce, invalidHash) {
		t.Errorf("validatePoW should have returned false for an invalid hash")
	}
}

// Helper function to generate a random string of a given length
func randStringBytes(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}
