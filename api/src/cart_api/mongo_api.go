package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	ID             string    `json:"id"`
	TimeStamp      time.Time `json:"timestamp"`
	RestuarantId   string    `json:"rid"`
	RestuarantName string    `json:"res"`
	ItemName       string    `json:"iname"`
	ItemId         string    `json:"iid"`
	Price          float32   `json:"price"`
	//Address   string        `json:"address"`
}
type Send struct {
	ItemId string `json:"iid"`
}

var (
	session *mgo.Session
	c       *mgo.Collection
)

var mongodb_server = "mongodb://admin:cmpe281@10.0.1.13,10.0.1.64,10.0.1.174,10.0.1.112,10.0.1.115"
var mongodb_database = "counter"
var mongodb_collection = "burger"

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
func main() {
	var err error
	session, err = mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c = session.DB(mongodb_database).C(mongodb_collection)

	server := NewServer()
	server.Run(":3000")
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/api/cart/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/cart/itemDisplay", cartDisplay(formatter)).Methods("GET")
	mx.HandleFunc("/api/cart/itemSave", cartSave(formatter)).Methods("POST")
	mx.HandleFunc("/api/cart/cartDelete", cartDelete(formatter)).Methods("DELETE")
	mx.HandleFunc("/atish", sendAtish(formatter)).Methods("GET")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

func cartSave(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var item Item
		err := json.NewDecoder(req.Body).Decode(&item)
		item.TimeStamp = time.Now()
		err = c.Insert(&item)
		if err != nil {
			panic(err)
		} else {
			log.Printf("inserted to cart")
		}
		j, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
		formatter.JSON(w, http.StatusOK, "saved to cart")
		return
	}
}

func cartDisplay(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var result []Item
		var item Item
		err := json.NewDecoder(req.Body).Decode(&item)
		err = c.Find(bson.M{"id": item.ID}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)
	}
}

func cartDelete(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var item Item
		var result []Item
		err := json.NewDecoder(req.Body).Decode(&item)
		err = c.Find(bson.M{"id": item.ID, "restuarantid": item.RestuarantId, "itemid": item.ItemId}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		for i := range result {
			fmt.Print(i)
			err = c.Remove(bson.M{"id": item.ID, "restuarantid": item.RestuarantId, "itemid": item.ItemId})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func sendAtish(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var item Item
		var result []Send

		err := json.NewDecoder(req.Body).Decode(&item)
		/*result := &map[string]interface{}{
			"res" : &res,
			"id" : &id,
			"iid" : &iid,
			"rid" : &rid,
			"iname" : &iname,
			"price" : &price,
		}*/

		err = c.Find(bson.M{"id": item.ID, "restuarantid": item.RestuarantId}).Select(bson.M{"itemid": 1}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)
	}
}
