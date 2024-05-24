package personhandlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kasualkid12/fr-website/server/modules/person"
)

func GetPersonsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		persons, err := person.GetPersons(db, 3)
		if err != nil {
			log.Printf("Error getting persons: %v", err)
			http.Error(w, "Error getting persons", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}
