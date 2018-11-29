// Version 2 API
// Using Negroni
// Reference:<https://github.com/paulnguyen/cmpe281/blob/master/golabs/gocloud/gumball/goapi/goapi_v3/src/gumball/server.go>
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
	Name      string        `json:"name"`
	Password  string        `json:"password"`
	TimeStamp time.Time     `json:"timestamp"`
	Contact   int           `json:"contact"`
	Address   string        `json:"address"`
}

// initRoutes for setting the API endpoint routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/api/ping", ping(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/signup", signup(formatter)).Methods("POST")
	mx.HandleFunc("/api/users/login", login(formatter)).Methods("GET")
}

// ping function for checking the status of the API
func ping(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 3.0 alive!"})
	}
}

// signup function fot creatinmg a new user
func signup(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)

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
		formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
		return
	}
}

// login function for getting into the system
func login(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqUser User
		var user User
		err := json.NewDecoder(req.Body).Decode(&reqUser)
		err = collection.Find(bson.M{"name": reqUser.Name}).One(&user)
		if reqUser.Password != user.Password {
			formatter.JSON(w, http.StatusOK, "Password Incorrect")
		} else if reqUser.Password == user.Password {
			formatter.JSON(w, http.StatusOK, "Login Successful")
		}
		if err != (nil) {
			panic(err)
		}
	}
}

// main function for creating the server and running it
func main() {

	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("counterBurger").C("users")

	server := NewServer()
	server.Run(":" + "3000")
}
