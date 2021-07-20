package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// ValidateHash return true if hash correcspond to the new hash Generated
func ValidateHash(algo string, hash string, data string) bool {
	newHash := calculatesha256(data)
	return newHash == hash
}

// GenerateHash generates de hash string
func GenerateHash(algo string, data string) string {
	return calculatesha256(data)
}

func calculatesha256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
