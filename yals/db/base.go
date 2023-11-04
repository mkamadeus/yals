package db

import bolt "go.etcd.io/bbolt"

type yalsBoltConnection struct {
	Conn       *bolt.DB
	BucketName string
}

func NewBoltDatabaseConnection(conn *bolt.DB, bucketName string) *yalsBoltConnection {
	return &yalsBoltConnection{conn, bucketName}
}
