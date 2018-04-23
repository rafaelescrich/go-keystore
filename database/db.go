package database

import (
	"errors"

	"github.com/boltdb/bolt"
	"github.com/rafaelescrich/go-keystore/keystore"
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
func (db BoltDB) Insert(ks keystore.Keystore, mk []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(mk)
		if err != nil {
			return err
		}
		sks, err := keystore.SerializeKeystore(ks)
		if err != nil {
			return err
		}
		err = b.Put(ks.Key, sks)
		if err != nil {
			return err
		}
		return nil
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

// GetAllKeys returns all keys from db
func (db BoltDB) GetAllKeys(masterkey []byte) ([]keystore.Keystore, error) {
	keys := []keystore.Keystore{}

	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(masterkey)
		b.ForEach(func(k, v []byte) error {
			dks, err := keystore.DeserializeKeystore(v)
			if err != nil {
				return err
			}
			keys = append(keys, keystore.Keystore{
				Key:      dks.Key,
				Filename: dks.Filename,
			})
			return nil
		})
		return nil
	})
	if err != nil {
		return keys, errors.New("Error on Get keys from DB, Message: " + err.Error())
	}

	return keys, nil
}
