package Data

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"atish/api/Cassandra"
"github.com/gorilla/mux"
"fmt"
"strconv"
"log"
)


type Identification struct {
  number string `json:"number"`
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
    } else {
      // created := true
    }
    // generate a unique UUID for this user
    if len(errs) == 0 {
	    // fmt.Println("creating a new user")
      // fmt.Println(common_name)
      // fmt.Println(strconv.Atoi(number))
	    gocqlUuid = gocql.TimeUUID()
	    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
	}
}

func Get(w http.ResponseWriter, r *http.Request){

  var errs []string
	// var gocqlUuid gocql.UUID
	var data string

	vars := mux.Vars(r)
  number := vars["number"]
  var common_name string
	number_num,_ := strconv.Atoi(number)
  iter := Cassandra.Session.Query(`SELECT * FROM cmpe281.cmpe281 WHERE number = ?`,number_num).Iter()
  // fmt.Println("Outside",number_num  )
  for iter.Scan(&number,&common_name) {
    // fmt.Println("In")
    // fmt.Println(number,common_name)
   fmt.Println("number:", number_num)
    // json.NewEncoder(w).Encode(number_num)
   data +="{number:" +number+","+ "common_name:"+ common_name+"}"
	 // idents = append(idents, )
    // json.NewEncoder(w).Encode(ident)
   json.NewEncoder(w).Encode(Identification{number: number, name: common_name})
  }
  if err := iter.Close(); err != nil {
    errs = append(errs, err.Error())
    log.Fatal(err)
  }
    // generate a unique UUID for this user
  // gocqlUuid = gocql.TimeUUID()
  // err := json.Unmarshal([]byte(data), identification)
  // fmt.Println(data)
  // s2, _ := json.Marshal(data)
  // result,_ := json.Marshal(data)
  json.NewEncoder(w).Encode(data)
}