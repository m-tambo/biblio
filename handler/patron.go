package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"biblio/schema"
	"database/sql"
)

type PatronsDAO interface {
	GetPatrons() []schema.Patron
}

func NewPatronDAO(db *sql.DB) PatronsDAO {
	return dao{db: db}
}

type getPatrons struct {
	pdao PatronsDAO
}

func NewGetPatronsHandler(pdao PatronsDAO) http.Handler {
	return getPatrons{pdao: pdao}
}

func (gp getPatrons) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	patrons := gp.pdao.GetPatrons()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patrons)
}

func (d dao) GetPatrons() []schema.Patron {
	var patrons []schema.Patron
	qry := "SELECT * FROM patron"

	rows, err := d.db.Query(qry)
	if err != nil {
		log.Printf("Error retrieving patrons %v", err)
	}

	for rows.Next() {
		var p schema.Patron
		err = rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Street, &p.City, &p.State, &p.Zip, &p.Dob, &p.Created)
		if err != nil {
			log.Printf("Error reading Patrons response: %v", err)
		}

		patrons = append(patrons, p)
	}

	return patrons
}
