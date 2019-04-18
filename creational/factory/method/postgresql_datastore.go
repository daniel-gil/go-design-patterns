package factory

import (
	"database/sql"
	"fmt"
)

// PostgreSQLDataStore is an implementation of DataStore interface
type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
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
