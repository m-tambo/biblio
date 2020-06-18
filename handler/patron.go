package handler

import "biblio/schema"

//type dao struct {
//	db *sql.DB
//}

type PatronsDAO interface {
	GetPatrons() []schema.Patron
}

func NewGetPatronsHandler {

}