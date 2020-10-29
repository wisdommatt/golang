package main

import (
	"net/http"

	"github.com/wisdommatt/internweb/routers"
)

func main() {
	http.HandleFunc("/users/new", routers.CreateUserHandler)
	http.HandleFunc("/users/all", routers.ReadUsersHandler)
	http.HandleFunc("/users/single", routers.ReadUserHandler)
	http.HandleFunc("/users/single/update", routers.UpdateUserHandler)

	http.ListenAndServe(":8000", nil)
}
