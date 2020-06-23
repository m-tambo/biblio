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
	Books []PatronBook
}

type PatronBook struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
}

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pages       int       `json:"pages"`
	Year        int       `json:"year"`
	Created     time.Time `json:"created"`
}
