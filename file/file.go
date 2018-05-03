package file

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/rafaelescrich/go-keystore/keystore"
)

// Filename holds the file name
type Filename struct {
	Fn string
}

// CipheredFile is a map between filename and key
type CipheredFile map[Filename]keystore.Keystore

// WriteFile writes a slice of bytes to a file
func WriteFile(path string, content []byte) error {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile returns the content of a file
func ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content, err
}

// DeleteFile deletes a file from hard drive
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
