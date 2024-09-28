package main

import (
	"log"
	"my/modules/handler"
	"net/http"
)

func main() {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Handle GET /users endpoint
	mux.HandleFunc("GET /users", hdlr.GetUsersHandler)

	// Handle GET /users/{id} endpoint
	mux.HandleFunc("GET /users/{id}", hdlr.GetUserHandler)

	// Handle POST /users endpoint
	mux.HandleFunc("POST /users", hdlr.PostUserHandler)

	// Handle PUT /users/{id} endpoint
	mux.HandleFunc("PUT /users/{id}", hdlr.PutUserHandler)

	// Handle DELETE /users/{id} endpoint
	mux.HandleFunc("DELETE /users/{id}", hdlr.DeleteUserHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", mux))
}
