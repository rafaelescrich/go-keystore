package controller

import (
	"errors"

	"github.com/rafaelescrich/go-keystore/ciphering"
	"github.com/rafaelescrich/go-keystore/keystore"
)

// CreateMK creates a master key
func CreateMK(password string) error {
	keystore.MasterKey = ciphering.GenerateMasterKey(password)
	if keystore.MasterKey == nil {
		return errors.New("Error while creating master key")
	} else {
		return nil
	}
}
