package keystore

import (
	"encoding/json"
	"fmt"
	"os"
)

// MasterKey holds the mk while the program is running
var MasterKey []byte

// Keystore holds the file name and its key
type Keystore struct {
	Key   []byte
	Nonce []byte
}

// CipheredFile is a map between filename and key
type CipheredFile map[string][]byte

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

// SerializeKeystore serialize a keystore
func SerializeKeystore(ks Keystore) ([]byte, error) {
	js, err := json.Marshal(ks)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return js, nil
}

// DeserializeKeystore returns a struct
func DeserializeKeystore(js []byte) (Keystore, error) {
	var ks Keystore
	err := json.Unmarshal(js, ks)
	if err != nil {
		return Keystore{}, err
	}
	return ks, nil
}
