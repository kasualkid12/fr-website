package personhandlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kasualkid12/fr-website/server/modules/person"
)

type requestBody struct {
	ID int `json:"id"`
}

func GetPersonsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		persons, err := person.GetPersons(db, reqBody.ID)
		if err != nil {
			log.Printf("Error getting persons: %v", err)
			http.Error(w, "Error getting persons", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}
