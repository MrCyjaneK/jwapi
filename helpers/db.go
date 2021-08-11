package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger/v3"
)

var db *badger.DB

// To store data use:
// /api/db/set/<key>?value
// To get data use:
// /api/db/get/<key>
// Response will be plain text <value>

// DBInit - Load database, called initially
func DBInit(datadir string) {
	Mkdir(datadir + "/db")
	var err error
	db, err = badger.Open(badger.DefaultOptions(datadir + "/db"))
	if err != nil {
		fmt.Println("[helpers][DBInit] Failed!", err)
		os.Exit(1)
	}
}

// Get value from db
func Get(key string) []byte {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err.Error() != "Key not found\n" {
				log.Println(key, "not found")
				return nil
			}
			log.Fatal(err.Error())
		}
		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.
			value = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			log.Fatal("???", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal("a", err)
	}
	return value
}

// Set value
func Set(key string, value []byte) {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}
