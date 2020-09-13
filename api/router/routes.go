package router

import (
	"log"
	"net/http"

	"github.com/Amaterasus/go-reckoning-backend/api/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// HandleRequests creates the router for the applicaation 
func HandleRequests(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", controllers.AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", controllers.ShowUser).Methods("GET")
	myRouter.HandleFunc("/users", controllers.Signup).Methods("POST")
	myRouter.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PATCH")
	myRouter.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	myRouter.HandleFunc("/authorised", controllers.Authorised).Methods("GET")
	myRouter.HandleFunc("/login", controllers.Login).Methods("POST")


	corsOrigins := handlers.AllowedOrigins([]string{"https://reckoning.netlify.app", "*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Authorised"})
	
	
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(corsOrigins, corsHeaders)(myRouter)))
}