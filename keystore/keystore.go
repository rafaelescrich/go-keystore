package keystore

import (
	"encoding/json"
	"os"
)

// MasterKey holds the mk while the program is running
var MasterKey []byte

// Keystore holds the file name and its key
type Keystore struct {
	Key   []byte
	Nonce []byte
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

// SerializeKeystore serialize a keystore
func SerializeKeystore(ks Keystore) ([]byte, error) {
	kss, err := json.Marshal(ks)
	if err != nil {
		return nil, err
	}
	return kss, nil
}

// DeserializeKeystore returns a struct
func DeserializeKeystore(kss []byte) (Keystore, error) {
	var ks Keystore
	err := json.Unmarshal(kss, &ks)
	if err != nil {
		return Keystore{}, err
	}
	return ks, nil
}
