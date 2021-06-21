package main

import (
	"fmt"
	"log"
	"net/http"

	// Type go get -u github.com/gorilla/mux to install
	"github.com/gorilla/mux" // Unused packages will create compilation error
)

var serverPort string

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Go-Lang GORM tutorial Code Demo")
}

func handleRequests() {
	// := is the short variable declaration operator
	// Automatically determines type for variable
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcomePage).Methods("GET")
	router.HandleFunc("/users", userList).Methods("GET")
	router.HandleFunc("/create-user", createUser).Methods("POST")
	router.HandleFunc("/get-user/{id}", getUser).Methods("GET")
	router.HandleFunc("/delete-user/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/update-user", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(serverPort, router))
}

func main() {
	dbConnection()    // DB connection
	userDBMigration() // User Table migration
	serverPort = ":9000"
	fmt.Printf("Website Started at http://localhost%v", serverPort)
	handleRequests()
}
