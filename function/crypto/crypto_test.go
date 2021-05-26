package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var flogoEncrypt = &fnEncrypt{}

var flogoDecrypt = &fnDecrypt{}

func TestEncryptedAnddecryptedTextValue(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("example plain text")

	ciphertextInterface, err := flogoEncrypt.Eval(key, plaintext)

	var encryptedText []byte = ciphertextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, encryptedText)

	// Decrypt same text
	plaintextInterface, err := flogoDecrypt.Eval(key, encryptedText)
	var decryptedText []byte = plaintextInterface.([]byte)

	assert.Nil(t, err)
	assert.NotNil(t, ciphertextInterface)
	assert.NotNil(t, decryptedText)
	assert.Equal(t, plaintext, decryptedText)
}
