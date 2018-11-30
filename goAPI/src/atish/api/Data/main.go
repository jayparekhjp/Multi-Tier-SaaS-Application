package Data

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"atish/api/Cassandra"
"github.com/gorilla/mux"
"fmt"
"strconv"
// "log"
)


type Identification struct {
  number int `json:"number"`
  name string `json:"name"`
}


func Post(w http.ResponseWriter, r *http.Request) {
  	var errs []string
  	var gocqlUuid gocql.UUID

  	r.ParseForm()
  	number := r.FormValue("number")
  	common_name :=r.FormValue("common_name")
    if err := Cassandra.Session.Query("INSERT INTO cmpe281 (number,common_name) VALUES (?,?)", number,common_name).Exec(); err != nil {
      errs = append(errs, err.Error())
    }
    if len(errs) == 0 {
	    gocqlUuid = gocql.TimeUUID()
	    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
	  }
}

func Get(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
  number := vars["number"]
  var common_name string
	number_num,_ := strconv.Atoi(number)
  iter := Cassandra.Session.Query(`SELECT * FROM cmpe281.cmpe281 WHERE number = ?`,number_num).Iter()
  for iter.Scan(&number,&common_name) {
   // fmt.Println("number:", number_num)
   var  data = Identification{
    number : number_num,
    name : common_name,
   }
   fmt.Println(data)
   json.NewEncoder(w).Encode(data)
  }
}


func UpdateHandler(w http.ResponseWriter, r *http.Request){
  var errs []string
  var gocqlUuid gocql.UUID

  r.ParseForm()
  number := r.FormValue("number")
  common_name :=r.FormValue("common_name")
  if err := Cassandra.Session.Query("UPDATE cmpe281.cmpe281 set common_name=? WHERE number = ?",common_name,number).Exec(); err != nil {
    errs = append(errs, err.Error())
  }
  if len(errs) == 0 {
    gocqlUuid = gocql.TimeUUID()
    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
  }
}


/*func DeleteHandler(w http.ResponseWriter, r *http.Request){
  var errs []string
  var gocqlUuid gocql.UUID
  fmt.Println("here")
  r.ParseForm()
  number := r.FormValue("number")
  // common_name :=r.FormValue("common_name")
  fmt.Println(number);
  if err := Cassandra.Session.Query("DELETE FROM cmpe281.cmpe281 WHERE number = ?",number).Exec(); err != nil {
    errs = append(errs, err.Error())
  }
  if len(errs) == 0 {
    gocqlUuid = gocql.TimeUUID()
    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
  }
}*/