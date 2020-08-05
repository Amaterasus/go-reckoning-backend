package controllers

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"

	"github.com/Amaterasus/reckoning-backend/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// This is required for using postgres with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllUsers will return a JSON response of all users in the database
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All users endpoint hit")
	
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	var users []models.User

	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// ShowUser will return a JSON response of a single user in the database based on the ID given
func ShowUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show user endpoint hit")

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	vars := mux.Vars(r)

	id := vars["id"]
	var user models.User

	db.First(&user, id)

	json.NewEncoder(w).Encode(user)
}

// NewUser will create a new user in the database and return a JSON response of that user
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new user endpoint hit")
	
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	name := r.FormValue("name")
	email := r.FormValue("email")

	user := db.Create(&models.User{Username: name, Email: email})

	fmt.Println("New user added to DataBase")

	json.NewEncoder(w).Encode(user.Value)
}

// UpdateUser will find a user in the database and update their email address
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user endpoint hit")

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	vars := mux.Vars(r)

	id := vars["id"]

	email := r.FormValue("email")

	var user models.User
	db.Where("id = ?", id).Find(&user)
	
	user.Email = email

	db.Save(&user)

	fmt.Println("User successfully updated")

	json.NewEncoder(w).Encode(user)
}

// DeleteUser will find a user based on their ID and delete them from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user endpoint hit")
	
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	vars := mux.Vars(r)

	id := vars["id"]

	var user models.User

	db.Where("id = ?", id).Find(&user)
	db.Delete(&user)

	fmt.Println("User successfully deleted")
	m := make(map[string]string)
    m["Message"] = "User Deleted!"
    json.NewEncoder(w).Encode(m)
}