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
}

// UserResource struct
type UserResource struct {
	User User `json:"user"`
}

// UsersResource struct
type UsersResource struct {
	Users []User `json:"users"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("pingHandler Function Executed")
	// Send message suggesting API is working
	fmt.Fprint(w, "API is ALIVE!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("getUsersHandler Function Executed")

	var users []User

	// Iteration over collection to get users
	iter := collection.Find(nil).Iter()
	result := User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(UsersResource{Users: users})
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("getUserHandler Function Executed")
	fmt.Println("TO BE IMPLEMENTED")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	log.Println("createUserHandler Function Executed")

	var userResource UserResource

	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		panic(err)
	}

	user := userResource.User
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
	j, err := json.Marshal(UserResource{User: user})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	fmt.Println("updateUser function executed")
	fmt.Println("TO BE IMPLEMENTED")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// function execution log
	fmt.Println("deleteUser function executed")
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
	r.HandleFunc("/api/ping", pingHandler).Methods("GET")
	r.HandleFunc("/api/users", getUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/api/users", createUserHandler).Methods("POST")
	r.HandleFunc("/api/users", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users", deleteUserHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
