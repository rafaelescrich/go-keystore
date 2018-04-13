package ciphering

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// Salt is hardcoded
const Salt = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// GenerateMasterKey is the method to generate a key from the salt and
// password
func GenerateMasterKey(password string) []byte {
	return pbkdf2.Key([]byte(password), []byte(Salt), 64000, 32, sha512.New)
}

// GenerateNonce generates new nonce to be used on encrypt aes gcm
func GenerateNonce() ([]byte, error) {
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

// EncryptAESGCM encrypt plaintext with the key in aes gcm
func EncryptAESGCM(key []byte, nonce []byte, plaintext []byte) ([]byte, []byte, error) {
	var ciphertext []byte

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ciphertext, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, ciphertext, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// DecryptAESGCM decrypts a ciphertext with a key
func DecryptAESGCM(key []byte, nonce []byte, ciphertext []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
