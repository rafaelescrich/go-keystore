package ciphering

import (
	"encoding/hex"
	"testing"
)

func TestGenerateMasterKey(t *testing.T) {
	t.Log("Testing")

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
