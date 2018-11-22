package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoServer = "localhost:27017"
var mongoDatabase = "counterBurger"
var mongoCollection = "users"

type User struct {
	ID       string
	Name     string
	Password string
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("USER API ALIVE!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := mgo.Dial(mongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongoDatabase).C(mongoCollection)
	var user []User
	err = c.Find(bson.M{}).All(&user)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	id := r.FormValue("id")

	session, err := mgo.Dial(mongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongoDatabase).C(mongoCollection)

	var user []User
	err = c.Find(bson.M{"ID": id}).One(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TO BE IMPLEMENTED")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ping", ping).Methods("GET")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/user", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users", updateUser).Methods("PUT")
	r.HandleFunc("/api/users", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
