package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/rafaelescrich/go-keystore/ciphering"
	"github.com/rafaelescrich/go-keystore/database"
	"github.com/rafaelescrich/go-keystore/file"
	"github.com/rafaelescrich/go-keystore/keystore"
)

// DB returns db
var db *database.BoltDB

// InitDB instantiate a db
func InitDB() {
	var err error
	db, err = database.InitDB()
	if err != nil {
		fmt.Printf("BoltDB Error: %s \r\n", err)
		os.Exit(1)
	}
}

// CreateMK creates a master key
func CreateMK(password string) error {
	keystore.MasterKey = ciphering.GenerateMasterKey(password)
	if keystore.MasterKey == nil {
		return errors.New("Error while creating master key")
	} else {
		return nil
	}
}

// GetAllKeys returns all keys
func GetAllKeys() ([]file.CipheredFile, error) {
	fks, err := db.GetAllKeys(keystore.MasterKey)
	if err != nil {
		return nil, err
	}
	return fks, nil
}

// EncryptFile receives a file and encrypts do another one
func EncryptFile(fn string) error {
	fl, err := file.ReadFile(fn)
	if err != nil {
		return err
	}
	nonce := ciphering.GenerateStreamBytes(12)
	key := ciphering.GenerateStreamBytes(16)
	ks := keystore.Keystore{Key: key, Nonce: nonce}
	kss, err := keystore.SerializeKeystore(ks)
	if err != nil {
		return err
	}
	ct, err := ciphering.EncryptAESGCM(key, nonce, fl)
	if err != nil {
		return err
	}
	newFilename := fn + ".aes"
	err = file.WriteFile(newFilename, ct)
	if err != nil {
		return err
	}
	err = db.Insert([]byte(fn), kss, keystore.MasterKey)
	if err != nil {
		return err
	}
	return nil
}

// DecryptFile receives a file and decrypts to the original one
func DecryptFile(fn string) error {
	kss, err := db.Get([]byte(fn), keystore.MasterKey)
	if err != nil {
		return err
	}
	ks, err := keystore.DeserializeKeystore(kss)
	if err != nil {
		return err
	}
	newFilename := fn + ".aes"
	ct, err := file.ReadFile(newFilename)
	if err != nil {
		return err
	}
	pt, err := ciphering.DecryptAESGCM(ks.Key, ks.Nonce, ct)
	if err != nil {
		return err
	}
	newFilename = "decrypted" + fn
	err = file.WriteFile(newFilename, pt)
	if err != nil {
		return err
	}
	return nil
}
