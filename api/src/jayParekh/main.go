package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id       string
	Name     string
	Password string
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func upadteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ping", ping).Methods("GET")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users", updateUser).Methods("PUT")
	r.HandleFunc("/api/users", deleteUSer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
