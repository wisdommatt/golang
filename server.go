package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type user struct {
	ID          int    `json:"id"`
	FullName    string `json:"name"`
	Email       string `json:"email"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
}

var users = []user{}

func (usr *user) save() {
	usr.ID = len(users) + 1
	users = append(users, *usr)
}

func main() {
	http.HandleFunc("/users/new", CreateUserHandler)
	http.HandleFunc("/users/all", ReadUsersHandler)
	http.HandleFunc("/users/single", ReadUserHandler)
	http.HandleFunc("/users/single/update", UpdateUserHandler)

	http.ListenAndServe(":8000", nil)
}

func CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	var submittedUser user

	err := json.NewDecoder(request.Body).Decode(&submittedUser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submitted !"))
		return
	}

	submittedUser.save()
	json.NewEncoder(response).Encode(submittedUser)
}

func ReadUsersHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(users)
}

func ReadUserHandler(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("userid")
	userID, err := strconv.Atoi(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid user ID !"))
		return
	}

	if userID > len(users) || userID < 1 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("User id is out of range"))
		return
	}

	userWithID := users[userID-1]
	json.NewEncoder(response).Encode(userWithID)
}

func UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("userid")
	userID, err := strconv.Atoi(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid user ID !"))
		return
	}

	if userID > len(users) || userID < 1 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("User id is out of range"))
		return
	}

	userWithID := &users[userID-1]
	err = json.NewDecoder(request.Body).Decode(userWithID)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid data submited"))
		return
	}

	json.NewEncoder(response).Encode(userWithID)
}
