package handler

import (
	"time"

	"biblio/schema"
)

type MockDAO struct {

}

func (d MockDAO) GetAuthors() []schema.Author {
	return []schema.Author{
		{
			ID:        1,
			FirstName: "jon",
			LastName:  "dick",
			Dob:       "xxx",
			Created:   time.Now(),
		},
	}
}