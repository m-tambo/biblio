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
	Books     []Book
}

type Book struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Pages       int        `json:"pages,omitempty"`
	Year        int        `json:"year,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}
