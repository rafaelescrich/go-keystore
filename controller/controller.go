package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/rafaelescrich/go-keystore/ciphering"
	"github.com/rafaelescrich/go-keystore/database"
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
func GetAllKeys() ([]keystore.Keystore, error) {
	var keys []keystore.Keystore
	keys, err := db.GetAllKeys(keystore.MasterKey)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// CreateKey creates and insert new key
func CreateKey() {

}
