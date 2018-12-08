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
    OrderID     bson.ObjectId  `json:"orderid"`
    TimeStamp   time.Time       `json:"timestamp"`
    Name        string          `json:"name"`
    Number      string          `json:"number"`
    Month       string          `json:"month"`
    Year        string          `json:"year"`
    CVV         string          `json:"cvv"`
    ID          string          `json:"id"`
    Total       string           `json:"total"`
}

type Cart struct {
    //ID            bson.ObjectId   `bson:"_id" json:"id"`
    //TimeStamp     time.Time       `json:"timestamp"`
    Items       string          `json:"items"`
    RestaurantName  string      `json:"res"`
    Price       float32         `json:"price"`
    //SubTotal      float32         `json:"subT"`
    UserId      int             `json:"uid"`
    //CartId        int             `json:"cid"`
}



type Order struct {
    ID          string          `json:"id"`
    OrderID     bson.ObjectId  `json:"orderid"`
    Total       string           `json:"total"`
}

var mongodb_server = "mongodb://admin:yesha@10.0.1.106,10.0.1.186,10.0.1.204,10.0.1.170,10.0.1.247"

var (
    session *mgo.Session
    c       *mgo.Collection
)

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
    session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c = session.DB("counterBurger").C("card")
        
    
    server := NewServer()
    server.Run(":3000")
}


func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
    mx.HandleFunc("/card", getcardHandler(formatter)).Methods("GET")
    mx.HandleFunc("/cart", getcartHandler(formatter)).Methods("GET")
    mx.HandleFunc("/orders", postCard(formatter)).Methods("POST")
    mx.HandleFunc("/cart", postCart(formatter)).Methods("POST")
    mx.HandleFunc("/details", getDetails(formatter)).Methods("GET")
}

func pingHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
    }
}


func getcardHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        result := make([]bson.M,0)
        
        err := c.Find(bson.M{}).All(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Card:", result )
        formatter.JSON(w, http.StatusOK, result)
    }
}


func getcartHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        session, err := mgo.Dial(mongodb_server)
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
    
        var card Card
        objID := bson.NewObjectId()
        card.OrderID =   objID
        err := json.NewDecoder(req.Body).Decode(&card)
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
        // formatter.JSON(w, http.StatusOK, "Welcome to Counfter Burger")
        return
    }
}



func postCart(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB("counterBurger").C("cart")
        var cart Cart
        err = json.NewDecoder(req.Body).Decode(&cart)
        //objID := bson.NewObjectId()
        //cart.ID = objID
        //cart.TimeStamp = time.Now()
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
        //formatter.JSON(w, http.StatusOK, "Welcome to Counter Burger")
        return
    }
}

func getDetails(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        var result []Order
        var card Card
        err := json.NewDecoder(req.Body).Decode(&card)
        err = c.Find(bson.M{"id": card.ID}).All(&result)
        if err != nil {
            log.Fatal(err)
        }
        formatter.JSON(w, http.StatusOK, result)
    }
}
