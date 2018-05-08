package database

import (
	"errors"

	"github.com/rafaelescrich/go-keystore/file"
	"github.com/rafaelescrich/go-keystore/keystore"

	"github.com/boltdb/bolt"
)

// BoltDB pointer
type BoltDB struct {
	DB *bolt.DB
}

// InitDB opens or create db file
func InitDB() (*BoltDB, error) {

	db, err := bolt.Open("keystore.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltDB{db}, nil
}

// Insert a key value pair in the db with the bucket name being the pbkdf2 master key
func (db BoltDB) Insert(filename []byte, kss []byte, mk []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(mk)
		if err != nil {
			return err
		}
		err = b.Put(filename, kss)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete the key from database
func (db BoltDB) Delete(key []byte, mk []byte) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(mk)
		return bucket.Delete(key)
	})
}

// Get returns nonce from filename
func (db BoltDB) Get(fl []byte, masterkey []byte) ([]byte, error) {
	var kss []byte

	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(masterkey)
		if b == nil {

			return errors.New("bucket not found")
		}
		kss = b.Get(fl)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return kss, nil
}

// GetAllKeys returns all keys from db
func (db BoltDB) GetAllKeys(masterkey []byte) ([]file.CipheredFile, error) {
	fks := make([]file.CipheredFile, 1, 1)
	var cf file.CipheredFile
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(masterkey)
		b.ForEach(func(k, v []byte) error {
			filename := string(k)
			key, err := keystore.DeserializeKeystore(v)
			if err != nil {
				return err
			}
			cf[filename] = key
			fks = append(fks, cf)
			return nil
		})
		return nil
	})
	if err != nil {
		return fks, errors.New("Error on Get keys from DB, Message: " + err.Error())
	}

	return fks, nil
}
