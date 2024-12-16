package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"

	"github.com/project-flogo/core/support/log"
)

var logger = log.RootLogger()

// aesEncrypt encrypts data using AES (Advanced Encryption Standard) in Cipher Block Chaining (CBC) mode.
// It generates a random initialization vector (IV) and appends it to the ciphertext.
// The resulting ciphertext is then base64 encoded and returned as a string.
//
// Parameters:
// - data: A byte slice containing the data to be encrypted.
// - aesKey: A byte slice containing the AES key used for encryption. The key must be 16, 24, or 32 bytes long.
//
// Returns:
// - A string representing the base64 encoded AES-encrypted data.
// - An error if the AES key is invalid or if any other encryption error occurs.
func aesEncrypt(data []byte, aesKey []byte) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}




// aesDecrypt decrypts a base64 encoded AES-encrypted data using the provided AES key.
//
// The function takes a base64 encoded string and a byte slice representing the AES key as input.
// It first decodes the base64 encoded string into a byte slice.
// Then, it creates a new cipher block using the provided AES key.
// If the length of the decoded data is less than the AES block size, it returns an error indicating that the ciphertext is too short.
// The initial vector (IV) is extracted from the beginning of the decoded data.//+
// The remaining decoded data is then decrypted using the CFB (Cipher Feedback) mode of operation with the AES cipher block and IV.
// Finally, the decrypted data is returned as a byte slice.//+
//
// Parameters:
// - data: A string representing the base64 encoded AES-encrypted data.//+
// - aesKey: A byte slice containing the AES key used for decryption.//+
//
// Returns:
// - A byte slice representing the decrypted data.//+
// - An error if the AES key is invalid, the ciphertext is too short, or any other decryption error occurs.//+
func aesDecrypt(data string, aesKey []byte) ([]byte, error) {
	decoded, _ := base64.StdEncoding.DecodeString(data)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	if len(decoded) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(decoded, decoded)

	return []byte(fmt.Sprintf("%s", decoded)), nil
}



// hmacValue generates a SHA512 HMAC (Hash-based Message Authentication Code) for the provided data using the given HMAC key.
// The HMAC is then encoded to base64 and returned as a string.
//
// This function uses the SHA512 hash function from the crypto/sha512 package and the HMAC function from the crypto/hmac package.
// The HMAC key is truncated to the length required by the SHA512 hash function (64 bytes).
//
// Parameters:
// - data: A byte slice containing the data for which the HMAC needs to be generated.
// - hmacKey: A byte slice containing the HMAC key used for authentication.
//
// Returns:
// - A string representing the base64 encoded SHA512 HMAC of the provided data.
func hmacValue(data []byte, hmacKey []byte) string {
    h512 := hmac.New(sha512.New, hmacKey[:])
    io.WriteString(h512, string(data))
    hexDigest := fmt.Sprintf("%x", h512.Sum(nil))
    return base64.StdEncoding.EncodeToString([]byte(hexDigest))
}


// Checksum generates a SHA256 checksum of the provided data and returns it as a hexadecimal string.
//
// The SHA256 algorithm is a cryptographic hash function that produces a 256-bit (32-byte) hash value.
// It is widely used for data integrity verification, such as checking if a file has been tampered with.
//
// Parameters:
// - data: A byte slice containing the data for which the checksum needs to be generated.
//
// Returns:
// - A string representing the SHA256 checksum of the provided data in hexadecimal format.
func checksum(data []byte) string {
    h := sha256.New()
    h.Write(data)
    return hex.EncodeToString(h.Sum(nil))
}

// rsaEncrypt encrypts data using the provided RSA public key.
// It uses the RSAES-PKCS1-v1_5 encryption scheme from the RSA Cryptography
// Standard (PKCS#1) and returns the encrypted data as a byte slice.
//
// Parameters:
// - data: A byte slice containing the data to be encrypted.
// - publicKey: A byte slice containing the RSA public key in PEM format.
//
// Returns:
// - A byte slice representing the encrypted data.
// - An error if the public key is invalid or if the encryption process fails.
func rsaEncrypt(data []byte, publicKey []byte) ([]byte, error) {
    block, _ := pem.Decode(publicKey)
    if block == nil {
        return nil, errors.New("Public Key Error")
    }

    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    pub := pubInterface.(*rsa.PublicKey)
    return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}


// rsaDecrypt decrypts data using the provided RSA private key.
// It uses the RSAES-PKCS1-v1_5 encryption scheme from the RSA Cryptography
// Standard (PKCS#1) and returns the decrypted data as a byte slice.
//
// Parameters:
// - ciphertext: A byte slice containing the encrypted data to be decrypted.
// - privateKey: A byte slice containing the RSA private key in PEM format.
//
// Returns:
// - A byte slice representing the decrypted data.
// - An error if the private key is invalid or if the decryption process fails.
func rsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
    block, _ := pem.Decode(privateKey)
    if block == nil {
        return nil, errors.New("Private Key Error")
    }
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}



// encodeBase64 encodes a byte slice to a base64 encoded string.
//
// This function takes a byte slice as input and uses the standard base64 encoding provided by the Go standard library.
// It returns the base64 encoded string.
//
// Parameters:
// - data: A byte slice containing the data to be encoded.
//
// Returns:
// - A string representing the base64 encoded data.
func encodeBase64(data []byte) string {
    res := base64.StdEncoding.EncodeToString(data)
    return res
}



// decodeBase64 decodes a base64 encoded string to a byte slice.
//
// This function takes a base64 encoded string as input and uses the standard base64 decoding provided by the Go standard library.
// It returns the decoded byte slice and an error if any occurred during the decoding process.
//
// Parameters:
// - data: A string representing the base64 encoded data to be decoded.
//
// Returns:
// - A byte slice representing the decoded data.
// - An error if the decoding process fails.
func decodeBase64(data string) ([]byte, error) {
    res, err := base64.StdEncoding.DecodeString(data)
    if err != nil {
        return nil, err
    }
    return res, nil
}
