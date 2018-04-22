package ciphering

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/golang/crypto/argon2"
)

// Salt is hardcoded
const Salt = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// GenerateMasterKey is the method to generate a key from the salt and
// password
func GenerateMasterKey(password string) []byte {
	// The draft RFC recommends time=3, and memory=32*1024
	// is a sensible number. If using that amount of memory (32 MB) is
	// not possible in some contexts then the time parameter can be increased
	//  to compensate.
	// Key(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32)
	return argon2.Key([]byte(password), []byte(Salt), 3, 32*1024, 4, 32)

}

// GenerateNonce generates new nonce to be used on encrypt aes gcm
func GenerateNonce(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
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
