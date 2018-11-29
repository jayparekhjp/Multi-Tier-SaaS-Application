package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

// User struct for defining user componenets
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	Password  string        `json:"password"`
	TimeStamp time.Time     `json:"timestamp"`
	Contact   int           `json:"contact"`
	Address   string        `json:"address"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("Ping Function Executed")
	// Send message suggesting API is working
	fmt.Fprint(w, "Login/SignUp API is ALIVE!")
}

func getusers(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("getusers Function Executed")

	var users []User

	// Iteration over collection to get users
	iter := collection.Find(nil).Iter()
	result := User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func login(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("Login Function Executed")
	fmt.Println("TO BE IMPLEMENTED")
}

func signup(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("SignUp Function Executed")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user User
	log.Println(r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Generate a new ID
	objID := bson.NewObjectId()
	// Transfer this ID to user
	user.ID = objID
	// Give TimeStamp to user
	user.TimeStamp = time.Now()
	// Insert user into collection
	err = collection.Insert(&user)
	if err != nil {
		panic(err)
	} else {
		log.Printf("Successfully created User: %s", user.Name)
	}
	j, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	// function execution log
	fmt.Println("updatePassword function executed")
	fmt.Println("TO BE IMPLEMENTED")
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	// function execution log
	fmt.Println("deleteAccount function executed")
	fmt.Println("TO BE IMPLEMENTED")
}

func main() {

	// function execution log
	log.Println("Main function executed")

	// Connecting to MongoDB
	log.Println("Starting MongoDB Session...")
	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("counterBurger").C("users")

	r := mux.NewRouter()

	r.HandleFunc("/api/ping", ping).Methods("GET")
	r.HandleFunc("/api/users", getusers).Methods("GET")
	r.HandleFunc("/api/users/{id}", login).Methods("GET")
	r.HandleFunc("/api/users", signup).Methods("POST")
	r.HandleFunc("/api/users", updatePassword).Methods("PUT")
	r.HandleFunc("/api/users", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
