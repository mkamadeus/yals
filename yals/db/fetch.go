package db

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func (y *yalsBoltConnection) FetchURLFromAlias(alias string) (string, error) {
	url := ""
	err := y.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("urls"))
		entry := bucket.Get([]byte(alias))
		if entry == nil {
			return fmt.Errorf("alias not found")
		}
		url = string(entry)
		return nil
	})

	return url, err
}
