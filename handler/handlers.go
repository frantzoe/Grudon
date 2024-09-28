package hdlr

import (
	"log"
	"my/modules/utility"
	"net/http"
)

var userList *utlt.Data = &utlt.Data{}

func init() {
	// Load users from file
	users, usersErr := utlt.LoadUsersFromFile()

	if usersErr != nil {
		log.Fatal("Error loading users:", usersErr)
		return
	}

	userList = users
}

// GetAllUsersHandler handles the GET /users endpoint
func GetUsersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	log.Println("Endpoint Hit: GET /users")

	utlt.EncodeResponse(writer, userList)
}

// GetUserByIdHandler handles the GET /users/{id} endpoint
func GetUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id := request.PathValue("id")

	log.Printf("Endpoint Hit: GET /users/%s", id)

	userFound := userList.FindUser(id)

	if userFound != nil {
		utlt.EncodeResponse(writer, userFound)
	} else {
		http.Error(writer, "User with id {" + id + "} not found.", http.StatusNotFound)
	}
}

// AddUserHandler handles the POST /users endpoint
func PostUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	log.Println("Endpoint Hit: POST /users")

	user, userErr := utlt.DecodeUser(request.Body)

	if userErr != nil {
		http.Error(writer, "Invalid request body.", http.StatusBadRequest)
		return
	}
	
	if utlt.AppendUser(userList, *user) {
		utlt.EncodeResponse(writer, userList)
	} else {
		http.Error(writer, "Error adding user.", http.StatusInternalServerError)
	}
}

// UpdateUserHandler handles the PUT /users/{id} endpoint
func PutUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id := request.PathValue("id")

	log.Printf("Endpoint Hit: PUT /users/%s", id)

	user, userErr := utlt.DecodeUser(request.Body)

	if userErr != nil {
		http.Error(writer, "Invalid request body.", http.StatusBadRequest)
		return
	}

	user.Id = id

	if utlt.UpdateUser(userList, *user) {
		utlt.EncodeResponse(writer, userList)
	} else {
		http.Error(writer, "User not found.", http.StatusNotFound)
	}
}

// DeleteUserHandler handles the DELETE /users/{id} endpoint
func DeleteUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	id := request.PathValue("id")

	log.Printf("Endpoint Hit: /users/%s", id)

	if utlt.DeleteUser(userList, id) {
		utlt.EncodeResponse(writer, userList)
	} else {
		http.Error(writer, "User not found.", http.StatusNotFound)
	}
}
