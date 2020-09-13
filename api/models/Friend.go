package models

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	// This is required for using postgres with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Friend struct {
	gorm.Model
	RequesterID uint
	RecieverID uint
	Accepted bool `gorm:"default:'false'"`
}


func InitialFriendsMigration() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.AutoMigrate(&Friend{})
}

func (friend *Friend) Create(requester, reciever uint) bool {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.Create(&Friend{RequesterID: requester, RecieverID: reciever, Accepted: false})

	return true
}