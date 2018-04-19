package keystore

import "os"

// MasterKey holds the mk while the program is running
var MasterKey []byte

// Keystore holds the file name and its key
type Keystore struct {
	Key      string
	Filename string
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
	return MasterKey != nil
}
