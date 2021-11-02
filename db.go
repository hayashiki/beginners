package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
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

func seed(db *sqlx.DB) error {
	merchant := Merchant{
		Email:    "hayashiki@gmail.com",
		Name:     "aioue",
		PhotoURL: "https://hoge.com",
	}

	ctx := context.Background()
	result := db.MustExecContext(ctx,
		"INSERT INTO merchants(email, name, photo_url)\nVALUES (?,?,?)",
		merchant.Email, merchant.Name, merchant.PhotoURL)

	log.Println("result.LastInsertId()")
	log.Println(result.LastInsertId())

	return nil
}
