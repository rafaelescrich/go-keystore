package keystore

import "os"

var masterkey []byte

// Keystore holds the file name and its key
type Keystore struct {
	Filename string
	Key      string
}

// FileExists return true if a file with that name exists
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// MasterkeyExists returns true if master key exists
func MasterkeyExists() bool {
	return masterkey != nil
}
