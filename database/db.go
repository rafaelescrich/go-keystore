package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/rafaelescrich/go-keystore/keystore"
)

// BoltDB pointer
type BoltDB struct {
	DB *bolt.DB
}

// InitDB opens or create db file
func InitDB() *BoltDB {

	db, err := bolt.Open("keystore.db", 0600, nil)
	if err != nil {
		fmt.Printf("BoltDB Error: %s \r\n", err)
		os.Exit(1)
	}

	return &BoltDB{db}
}

// Insert saves on daatabase a key value pair and the bucket name is the pbkdf2 master key
func (db BoltDB) Insert(keystore keystore.Keystore, bucket []byte) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		err = b.Put([]byte(keystore.Key), []byte(keystore.Filename))
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
func (db BoltDB) Delete(key string) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("keyBucket"))
		return bucket.Delete([]byte(key))
	})
}

// GetAllKeys returns all keys from db
func (db BoltDB) GetAllKeys(masterkey []byte) ([]keystore.Keystore, error) {
	keys := []keystore.Keystore{}

	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(masterkey)
		b.ForEach(func(k, v []byte) error {
			keys = append(keys, keystore.Keystore{
				Key:      string(k),
				Filename: string(v),
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
