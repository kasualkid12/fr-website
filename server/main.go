package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kasualkid12/fr-website/server/handlers/personhandlers"
	grabenv "github.com/kasualkid12/fr-website/server/modules/grabEnv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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
	router.HandleFunc("/persons", personhandlers.GetPersonsHandler(db)).Methods("POST")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "body"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
