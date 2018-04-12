package ciphering

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

// Salt is hardcoded
const Salt = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// TODO: Every time a new nonce is created have to save on db

// Nonce to use in encrypt/decrypt AES GCM
const Nonce = "64a9433eae7ccceee2fc0eda"

// GenerateMasterKey is the method to generate a key from the salt and
// password
func GenerateMasterKey(password string) []byte {
	return pbkdf2.Key([]byte(password), []byte(Salt), 4096, 32, sha512.New)
}

// EncryptAESGCM encrypt plaintext with the key in aes gcm
func EncryptAESGCM(key string, plaintext string) ([]byte, error) {
	keyEnconded, err := hex.DecodeString(key)
	pT := []byte(plaintext)
	var cT []byte

	block, err := aes.NewCipher(keyEnconded)
	if err != nil {
		return cT, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	// nonce := make([]byte, 12)
	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	return cT, err
	// }
	nonce, _ := hex.DecodeString(Nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return cT, err
	}

	cT = aesgcm.Seal(nil, nonce, pT, nil)
	return cT, nil
}

// DecryptAESGCM decrypts a ciphertext with a key
func DecryptAESGCM(key string, ciphertext string) (string, error) {
	keyEncoded, _ := hex.DecodeString(key)
	cT, _ := hex.DecodeString(ciphertext)
	nonce, _ := hex.DecodeString(Nonce)

	block, err := aes.NewCipher(keyEncoded)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, cT, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext[:]), nil
}
