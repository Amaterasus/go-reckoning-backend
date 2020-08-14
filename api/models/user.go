package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// This is required for using postgres with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User is the structure of the class being used for the database
type User struct {
	gorm.Model
	Username string 
	Email string 
	HashedPassword string 
}

// InitialUserMigration will use GORM to migrate the tables in the database.
func InitialUserMigration() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

// GetAllUsers Queries the database and returns all users
func (user *User) GetAllUsers() *[]User {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	users := []User{}


	db.Find(&users)

	return &users
}