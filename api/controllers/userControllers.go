package controllers

import (
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"net/http"

	"github.com/Amaterasus/go-reckoning-backend/api/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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


	var user models.User

	foundUser := user.FindUserByID(id)

	json.NewEncoder(w).Encode(foundUser)
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

// Login will varify the username and password and eventually should respond with a JWT for future verification
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login endpoint hit")
	
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user models.User

	db.Where("Username = ?", username).Find(&user)


	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if err != nil {
		m := make(map[string]string)
		m["Message"] = "Username and password do not match"
		json.NewEncoder(w).Encode(m)
	} else {
		json.NewEncoder(w).Encode(user)

	}

}