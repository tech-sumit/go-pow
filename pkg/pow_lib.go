package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

const difficulty = 4

func SolvePoW(nonce string) (string, error) {
	var solution string
	for {
		hash := sha256.Sum256([]byte(nonce + solution))
		hexHash := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hexHash, strings.Repeat("0", difficulty)) {
			return solution, nil
		}
		solution = hex.EncodeToString(hash[:]) // Generate next solution candidate
	}
}

func GenerateNonce() string {
	rand.Seed(time.Now().UnixNano())
	nonce := make([]byte, 16)
	for i := range nonce {
		nonce[i] = byte(rand.Intn(256))
	}
	return hex.EncodeToString(nonce)
}

func ValidatePoW(nonce, hash string) bool {
	target := strings.Repeat("0", difficulty)
	hashed := sha256.Sum256([]byte(nonce + hash))
	return strings.HasPrefix(hex.EncodeToString(hashed[:]), target)
}
