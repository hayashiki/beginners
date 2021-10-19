package main

type Merchant struct {
	ID       int
	Email    string
	Name     string
	PhotoURL string `db:"photo_url"`
}
