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

type Item struct {
	ID	  		string 			`json:"id"`
	CartId	  bson.ObjectId `bson:"cart_id" json:"cid"`
	TimeStamp time.Time     `json:"timestamp"`
	RestuarantName string   `json:"res"`
	ItemName       string   `json:"iname"`
	ItemId         int      `json:"iid"`
	Price          float32      `json:"price"`
	//Address   string        `json:"address"`
}

var (
	session    *mgo.Session
	c *mgo.Collection
)


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
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c = session.DB("counter").C("burger")

	server := NewServer()
	server.Run(":3000")
}



// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/api/cart/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/cart/itemDisplay", cartDisplay(formatter)).Methods("GET")
	//mx.HandleFunc("/gumball", gumballUpdateHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/api/cart/itemSave", cartSave(formatter)).Methods("POST")
	//mx.HandleFunc("/order/{id}", gumballOrderStatusHandler(formatter)).Methods("GET")

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
		cartID := bson.NewObjectId()
		item.CartId = cartID
		
		item.TimeStamp = time.Now()
		err = c.Insert(&item)
		if err != nil {
			panic(err)
		} else {
			log.Printf("Successfully created User:")
		}
		j, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
		formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
		return
	}
}


func cartDisplay(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {	
		var result []Item
		var item Item
		err := json.NewDecoder(req.Body).Decode(&item)
		err = c.Find(bson.M{"uid":item.ID}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		/*	for i := range result {
				if result[i].ID == item.ID {
					resu[i] = result[i]
				}
			}
		*/
		//fmt.Println("Burger orders:", resu )
		formatter.JSON(w, http.StatusOK, result)
	}
}
