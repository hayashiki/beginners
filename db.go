package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS merchants
(
	id INTEGER PRIMARY KEY,
	name TEXT,
	email TEXT,
	photo_url TEXT,
	timestamp DATETIME
)`

func prepareSchema(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("could not create http_requests table: %w", err)
	}
	return nil
}

