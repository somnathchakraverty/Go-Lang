package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var serverPort string

func createFileInBackground(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "heelo")
	w.Header().Set("Content-Type", "application/json")
	go createWrite()
	json.NewEncoder(w).Encode("File creating request received. Please check after 50 seconds.")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create-file-in-background", createFileInBackground).Methods("GET")
	log.Fatal(http.ListenAndServe(serverPort, router))
}

func main() {
	serverPort = ":9002"
	fmt.Printf("Website Started at http://localhost%v", serverPort)
	handleRequests()
}
