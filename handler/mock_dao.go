package handler

import (
	"database/sql"
	"time"

	"biblio/schema"
)

type MockDAO struct {
	db *sql.DB
}

func (md MockDAO) GetAuthors() []schema.Author {
	return []schema.Author{
		{
			ID:        1,
			FirstName: "jon",
			LastName:  "dick",
			Dob:       "xxx",
			Created:   time.Now(),
		},
		{
			ID:        2,
			FirstName: "maggie",
			LastName:  "baum",
			Dob:       "1997-3-24",
			Created:   time.Now(),
		},
	}
}

func (md MockDAO) CreateAuthor() {

}

func (md MockDAO) DeleteAuthor(id string) {

}