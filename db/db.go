package db

import (
	"database/sql"
	"fmt"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

type StorageDriver interface {
	QuoteDriver
	UserDriver
}

type DB struct {
	*sql.DB
}

func DialDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[DialDB] %s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[DialDB] %s", err)
	}

	return &DB{db}, nil
}
