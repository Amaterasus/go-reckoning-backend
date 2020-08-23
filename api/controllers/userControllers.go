package controllers

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"

	"github.com/Amaterasus/go-reckoning-backend/api/models"

	"github.com/gorilla/mux"
	// This is required for using postgres with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllUsers will return a JSON response of all users in the database
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All users endpoint hit")
	
	user := models.User{}

	users := user.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}

// ShowUser will return a JSON response of a single user in the database based on the ID given
func ShowUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show user endpoint hit")

	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	user := models.User{}

	user.FindUserByID(id)

	json.NewEncoder(w).Encode(user)
}

// NewUser will create a new user in the database and return a JSON response of that user
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New user endpoint hit")

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := &models.User{}

	newUser := user.Create(username, email, password)

	fmt.Println("New user added to DataBase")

	json.NewEncoder(w).Encode(newUser)
}

// UpdateUser will find a user in the database and update their email address
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user endpoint hit")

	vars := mux.Vars(r)

	id := vars["id"]

	email := r.FormValue("email")

	user := &models.User{}

	user.Update(id, email)

	json.NewEncoder(w).Encode(user)
}

// DeleteUser will find a user based on their ID and delete them from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user endpoint hit")
	
	vars := mux.Vars(r)

	id := vars["id"]

	var user models.User

	message := user.Destroy(id)

    json.NewEncoder(w).Encode(message)
}

// Login will varify the username and password and eventually should respond with a JWT for future verification
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login endpoint hit")
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user models.User

	if user.Authorise(username, password) {
		json.NewEncoder(w).Encode(user)
	} else {
		m := make(map[string]string)
		m["Message"] = "Username and password do not match"
		json.NewEncoder(w).Encode(m)
	}
}