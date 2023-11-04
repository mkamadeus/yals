package db

import (
	bolt "go.etcd.io/bbolt"
)

func (y *yalsBoltConnection) IsAliasUsed(alias string) (bool, error) {
	var exists bool
	y.Conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(y.BucketName))
		entry := bucket.Get([]byte(alias))
		exists = entry != nil
		return nil
	})

	return exists, nil
}
