package crypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	passphrase := "passphrase"
	input := "Text test"
	encrypt := EncryptString(input, passphrase)

	output, _ := DecryptString(encrypt, passphrase)

	assert.Equal(t, input, output)
	fmt.Printf("Input: %v Encrypted %s EncryptedLength: %d Output: %s\n", input, encrypt, len(encrypt), output)
}

func TestDecryptError(t *testing.T) {
	passphrase := "passphrase"
	encrypt := "5558e21d7dfe11bdf8249bffb24c1dc9211e3e677d2788bf0fa4f07c830a93cceba9a203c4sdf"
	fmt.Printf("Encrypted %s EncryptedLength: %d\n", encrypt, len(encrypt))

	_, err := DecryptString(encrypt, passphrase)
	assert.Error(t, err)
}
