package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "github.com/kasualkid12/fr-website/server/modules/encryption"
	"github.com/gorilla/mux"
	"github.com/kasualkid12/fr-website/server/handlers/personhandlers"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	host, port, user, password, dbname := grabenv.GrabEnv()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	
	// Use the handler from the person package
	router.HandleFunc("/persons", personhandlers.GetPersonsHandler(db)).Methods("GET")
	
	fmt.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
	
}
