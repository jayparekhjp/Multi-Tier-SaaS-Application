// Version 3 API
// Using Negroni

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

// MongoDB connections
// var mongoServer = "mongodb://admin:cmpe281@10.0.1.115,10.0.1.165,10.0.1.175,10.0.1.107,10.0.1.211"
var mongoServer = "localhost:27017"
var mongoDatabase = "counterBurger"
var mongoCollection = "users"

// NewServer configures and returns a server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// User struct defines elements of user
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Username  string        `bson:"username" json:"username"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Contact   string        `json:"contact"`
	Address   string        `json:"address"`
	Password  string        `json:"password"`
	TimeStamp time.Time     `json:"timestamp"`
}

// initRoutes for setting the API endpoint routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/api/ping", ping(formatter)).Methods("GET")
	mx.HandleFunc("/api/users", getAllUsers(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/signup", signup(formatter)).Methods("POST")
	mx.HandleFunc("/api/users/login", login(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/{id}", updatePassword(formatter)).Methods("PUT")
	mx.HandleFunc("/api/users/{id}", deleteAccount(formatter)).Methods("DELETE")
}

// ping function for checking the status of the API
func ping(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 3.0 alive!"})
	}
}

// getAllUsers function for testing purposes only
func getAllUsers(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqUser User
		var users []User
		err := json.NewDecoder(req.Body).Decode(&reqUser)
		if reqUser.Username == "admin" && reqUser.Password == "cmpe281" {
			formatter.JSON(w, http.StatusOK, "Welcome Master")
			iter := collection.Find(nil).Iter()
			result := User{}
			for iter.Next(&result) {
				users = append(users, result)
			}
			j, err := json.Marshal(users)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(j)

		} else {
			formatter.JSON(w, http.StatusOK, "You are not admin.")
			return
		}
		if err != (nil) {
			panic(err)
		}
	}
}

// signup function fot creatinmg a new user
func signup(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user User
		var testUser User
		err := json.NewDecoder(req.Body).Decode(&user)
		// Check if the username already exists
		// testUser.Username = user.Username
		err = collection.Find(bson.M{"username": user.Username}).One(&testUser)
		if testUser.Username != "" {
			formatter.JSON(w, http.StatusOK, "Username taken. Please try different username.")
			return
		} else {
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
			// j, err := json.Marshal(user)
			// if err != nil {
			// 	panic(err)
			// }
			w.Header().Set("Content-Type", "application/json")
			// w.Write(j)
			formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
			return
		}
	}
}

// login function for getting into the system
func login(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqUser User
		var user User
		err := json.NewDecoder(req.Body).Decode(&reqUser)
		err = collection.Find(bson.M{"username": reqUser.Username}).One(&user)
		if reqUser.Password != user.Password {
			formatter.JSON(w, http.StatusOK, "Password Incorrect")
			return
		} else if reqUser.Password == user.Password {
			formatter.JSON(w, http.StatusOK, user.ID)
			return
		}
		if err != (nil) {
			panic(err)
		}
	}
}

// updatePassword function for updating the password
func updatePassword(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, "updatePassword function not implemented yet")
		return
	}
}

// deleteAccount function for updating the password
func deleteAccount(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, "deleteAccount function not implemented yet")
		return
	}
}

// main function for creating the server and running it
func main() {

	var err error
	session, err = mgo.Dial(mongoServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB(mongoDatabase).C(mongoCollection)

	server := NewServer()
	server.Run(":" + "3000")
}
