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

// UpdatePersonHandler updates one or more fields of a person
func UpdatePersonHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type updateRequest struct {
			ID      int                    `json:"id"`
			Updates map[string]interface{} `json:"updates"`
		}
		var req updateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.ID == 0 || len(req.Updates) == 0 {
			http.Error(w, "Missing id or updates", http.StatusBadRequest)
			return
		}
		if err := person.UpdatePerson(db, req.ID, req.Updates); err != nil {
			log.Printf("Error updating person: %v", err)
			http.Error(w, "Error updating person", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Person updated successfully"))
	}
}
