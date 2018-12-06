package main

import (
	"fmt"
	"log"
	"encoding/json"
	"time"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


type Card struct {
	ID			bson.ObjectId   `bson:"_id" json:"id"`
	TimeStamp 	time.Time       `json:"timestamp"`
	Name 		string			`json:"name"`
	Number 		string			`json:"number"`
	CVV 		string			`json:"cvv"`
	Month 	    string     		`json:"month"`
	Year	    int				`json:"year"`
}

type Cart struct {
	ID			bson.ObjectId   `bson:"_id" json:"id"`
	TimeStamp 	time.Time       `json:"timestamp"`
	Name		string			`json:"name"`
	RestaurantName	string		`json:"res"`
	Price		float32			`json:"price"`
	UserId		int			`json:"uid"`
	CartId		int			`json:"cid"`
	
}

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

func main() {

	
	server := NewServer()
	server.Run(":3002")
}


func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/card", cardHandler(formatter)).Methods("GET")
	mx.HandleFunc("/cart", cartHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", postCard(formatter)).Methods("POST")
	mx.HandleFunc("/cart", postCart(formatter)).Methods("POST")
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}


func cardHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("counterBurger").C("card")
		
		result := make([]bson.M,0)
		
        err = c.Find(bson.M{}).All(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Card:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}


func cartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("counterBurger").C("cart")
		
		result := make([]bson.M,0)
		
        err = c.Find(bson.M{}).All(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Cart:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}

func postCard(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("counterBurger").C("card")
		var card Card
		err = json.NewDecoder(req.Body).Decode(&card)


		objID := bson.NewObjectId()

		card.ID = objID

		card.TimeStamp = time.Now()

		err = c.Insert(&card)
		if err != nil {
			panic(err)
		} else {
			log.Printf("Successfully created User:")
		}
		j, err := json.Marshal(card)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
		formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
		return
	}
}


func postCart(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("counterBurger").C("cart")
		var cart Cart
		err = json.NewDecoder(req.Body).Decode(&cart)
		objID := bson.NewObjectId()
		cart.ID = objID
		cart.TimeStamp = time.Now()
		err = c.Insert(&cart)
		if err != nil {
			panic(err)
		} else {
			log.Printf("Successfully created Cart details:")
		}
		j, err := json.Marshal(cart)
		log.Println(cart)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
		formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
		return
	}
}