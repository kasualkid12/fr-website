package personhandlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kasualkid12/fr-website/server/modules/person"
)

func GetPersonsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Grabbing and sending Persons data\n")
		persons, err := person.GetPersons(db)
		if err != nil {
			log.Printf("Error getting persons: %v", err)
			http.Error(w, "Error getting persons", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}
