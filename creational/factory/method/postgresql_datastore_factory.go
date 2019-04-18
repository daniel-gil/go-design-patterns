package factory

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type PostgreSQLDataStoreFactory struct{}

// Create is a factory method that return the common interface DataStore
func (pdsf *PostgreSQLDataStoreFactory) Create(conf map[string]string) (DataStore, error) {
	dsn := ""
	ok := false
	if dsn, ok = conf["DATASTORE_POSTGRES_DSN"]; !ok {
		return nil, fmt.Errorf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, fmt.Errorf("Failed to connect to datastore: %s", err.Error())
	}

	return &PostgreSQLDataStore{
		DSN: dsn,
		DB:  *db.DB,
	}, nil
}
