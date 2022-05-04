package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email_id"`
}

func userDBMigration() {
	DB.AutoMigrate(&User{})
	fmt.Println(" User DB was successfully Migrated. ")
}

func userList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " ******* User List ********* ")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " ******* User List ********* ")
	var user User
	params := mux.Vars(r)
	DB.Find(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, " ******* Create New User ********* ")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " ******* Delete User ********* ")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " ******* Update User ********* ")
	var user User
	params := mux.Vars(r)
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user is deleted successfully")
}
