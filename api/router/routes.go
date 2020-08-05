package router

import (
	"log"
	"net/http"

	"github.com/Amaterasus/reckoning-backend/api/controllers"

	"github.com/gorilla/mux"
)

// HandleRequests creates the router for the applicaation 
func HandleRequests(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", controllers.AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", controllers.ShowUser).Methods("GET")
	myRouter.HandleFunc("/users", controllers.NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PATCH")
	myRouter.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":" + port, myRouter))
}