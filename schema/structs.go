package schema

import (
	"time"
)

type Author struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Dob       string    `json:"dob"`
	Created   time.Time `json:"created"`
}

type Patron struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zip       int       `json:"zip"`
	Dob       string    `json:"dob"`
	Created   time.Time `json:"created"`
}

type Book struct {
	Title       string `json:"title"`
	Pages       int    `json:"pages"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
