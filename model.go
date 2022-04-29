package main

type Merchant struct {
	ID       int64
	Email    string
	Name     string
	PhotoURL string `db:"photo_url"`
}
