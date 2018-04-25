package file

import (
	"fmt"
	"io"
	"os"

	"github.com/rafaelescrich/go-keystore/keystore"
)

// Filename holds the file name
type Filename struct {
	Fn string
}

// CipheredFile is a map between filename and key
type CipheredFile map[Filename]keystore.Keystore

func createFile(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return err
		}
		defer file.Close()
	}

	return err
}

func writeFile(path string, content []byte) error {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.Write(content)
	if isError(err) {
		return err
	}
	// save changes
	err = file.Sync()
	if isError(err) {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func readFile(path string) ([]byte, error) {
	var content = make([]byte, 1024)
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return content, err
	}
	defer file.Close()

	// read file, line by line

	for {
		_, err = file.Read(content)
		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}
		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	if err != nil {
		return content, err
	}
	return content, nil
}

func deleteFile(path string) error {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
