package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func setupDatabase() (*bolt.DB, error) {
	db, err := bolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func updateDatabase(key string, val string) {
	fmt.Println(key)
	fmt.Println(val)
	db, _ := setupDatabase()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(DB_BUCKET_NAME))
		return b.Put([]byte(key), []byte(val))
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_BUCKET_NAME))
		storedVal := b.Get([]byte(key))
		fmt.Printf("Value Stored: %s\n", append(storedVal, (" With Key: "+key)...))
		return nil
	})
}
