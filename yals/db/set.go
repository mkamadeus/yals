package db

import "go.etcd.io/bbolt"

func (y *yalsBoltConnection) SetURLToAlias(alias, url string) error {
	err := y.Conn.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(y.BucketName))
		err := bucket.Put([]byte(alias), []byte(url))
		return err
	})

	return err
}
