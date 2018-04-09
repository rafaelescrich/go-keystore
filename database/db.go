package database

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// InitDB opens or create db file
func InitDB() {

	db, err := bolt.Open("keystore.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB open")
	defer db.Close()
}
