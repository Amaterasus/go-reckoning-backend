package main 

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"	
	"github.com/Amaterasus/reckoning-backend/api/models"
	"github.com/Amaterasus/reckoning-backend/api/router"
)

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	initialMigrations()
	router.HandleRequests(port)
}

func initialMigrations() {
	models.InitialUserMigration()
}