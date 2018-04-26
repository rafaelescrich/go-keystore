package ciphering

import (
	"encoding/hex"
	"testing"
)

func TestGenerateMasterKey(t *testing.T) {
	t.Log("Testing Master Key in Argon2i")

	// http://antelle.net/argon2-browser/
	// Password: hmrbbi
	// Salt: The Times 03/Jan/2009 Chancellor on brink of second bailout for banks
	// Memory: 32768
	// Iterations: 3
	// Hash length: 32
	// Parallelism: 4

	password := "hmrbbi"

	expectedResult := "7e58dcb20e8c7b4b7ec9fa574f18ae4a44a15c688b1c6b4d831cd12df5ec1673"

	result := GenerateMasterKey(password)

	t.Log(password)
	t.Log(expectedResult)
	t.Log(hex.EncodeToString(result))

	if expectedResult != hex.EncodeToString(result) {
		t.Error("Expected result is different from result")
	} else {
		t.Log("TEST PASSED")
	}
}

func TestEncryptAESGCM(t *testing.T) {
	t.Log("Testing Encrypt AES GCM")

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("exampleplaintext")
	nonce, _ := hex.DecodeString("787b3eae2824fcb3c9351f48")

	ciphertext, err := EncryptAESGCM(key, nonce, plaintext)

	if err != nil {
		t.Error(err)
	}

	expectedResult := "e6666e7047ac9c0a502c86aac59d420a343df303766cc9a93e1990357855227f"

	t.Logf("Plaintext: %s", string(plaintext))
	t.Logf("Ciphertext: %s", hex.EncodeToString(ciphertext))

	if expectedResult != hex.EncodeToString(ciphertext) {
		t.Error("Expected result is different from result")
	} else {
		t.Log("TEST PASSED")
	}

}

func TestDecryptAESGCM(t *testing.T) {
	t.Log("Testing Decrypt AES GCM")

	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")
	ciphertext, _ := hex.DecodeString("2df87baf86b5073ef1f03e3cc738de75b511400f5465bb0ddeacf47ae4dc267d")
	nonce, _ := hex.DecodeString("afb8a7579bf971db9f8ceeed")

	plaintext, err := DecryptAESGCM(key, nonce, ciphertext)

	if err != nil {
		t.Error(err)
	}

	expectedResult := "exampleplaintext"

	t.Logf("Plaintext: %s", string(plaintext))
	t.Logf("Ciphertext: %s", hex.EncodeToString(ciphertext))

	if expectedResult == string(plaintext) {
		t.Log("TEST PASSED")
	} else {
		t.Error("Expected result is different from result")
	}

}
