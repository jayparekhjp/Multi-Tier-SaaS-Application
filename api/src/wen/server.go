package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

// MongoDB Config
// var mongodb_server = "mongodb://admin:cmpe281@10.0.1.119,10.0.1.67,10.0.1.158,10.0.1.99,10.0.1.34"
var mongodb_server = "localhost:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "burgerOrder"

type Order struct {
	UserId string	`json:"UserId"`
	OrderId string   `json:"OrderId"`
	TotalPrice string  `json:"TotalPrice"`
}


// NewServer configures and returns a Server.
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

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/list", userOdersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/insert", burgerNewOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")


}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}


//*********************************** Burger New Order Handler*********************
func burgerNewOrderHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		var ord Order
		err := json.NewDecoder(req.Body).Decode(&ord)

		session, err := mgo.Dial(mongodb_server)
		if err != nil{
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		// var result bson.M
		// req.ParseForm()
		// x := req.Form.Get("OrderId")
		// y := req.Form.Get("UserId")
		// z := req.Form.Get("totalPrice")
		// result = bson.M{"OrderId" : x, "UserId" : y, "totalPrice" : z}
		err = c.Insert(&ord)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Burger orders:", result )
		formatter.JSON(w, http.StatusOK, ord)
	}
}




//*********************************** search user orders *********************

func userOdersHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		var ord Order
		session, err := mgo.Dial(mongodb_server)
		if err != nil{
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result [] Order
		// req.ParseForm()                     // Parses the request body
		// var uid string= req.Form.Get("UserId")
		// err = c.Find(bson.M{"UserId" : uid}).All(&result)
		err = json.NewDecoder(req.Body).Decode(&ord)
		err = c.Find(bson.M{"userid": ord.UserId}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Burger orders:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := NewServer()
	server.Run(":" + port)


}
