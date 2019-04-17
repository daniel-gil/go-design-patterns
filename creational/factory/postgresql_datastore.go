package factory

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// PostgreSQLDataStore is an implementation of DataStore interface
type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
}

// NewPostgreSQLDataStore is a factory method (constructor) that return the common interface DataStore
func NewPostgreSQLDataStore(conf map[string]string) (DataStore, error) {
	dsn := ""
	ok := false
	if dsn, ok = conf["DATASTORE_POSTGRES_DSN"]; !ok {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN"))
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

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error) {
	var username string
	rows, err := pds.DB.Query("SELECT username FROM users WHERE id=$1", id)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}

	err = rows.Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errUserNotFound
		}
		return "", err
	}
	return username, nil
}
