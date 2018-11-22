package main

import (
	
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	
	"github.com/gorilla/mux"
)
	/*
	type User struct {
		ID		string	`json:"id"`
		Name  	string	`json:"name"`
	}
	*/
	type Burger struct {
		Name  string
	}

func createBurger(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	name := r.FormValue("name")
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("burgerPlace").C("burgers")
	err = c.Insert(&Burger{Name: name})
	/*
		if err != nil {
			log.Fatal
		}
	*/
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/burgers", createBurger).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", r))
}