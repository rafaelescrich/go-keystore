package database

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
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

// Insert saves on satabase a key value pair
func (db BoltDB) Insert(key string, value string) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("keyBucket"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
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
