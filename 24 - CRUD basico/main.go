package main

import (
	"crud-basico/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetAllUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.GetUserId).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta :5000 🔥 !")
	log.Fatal(http.ListenAndServe(":5000", router))
}
