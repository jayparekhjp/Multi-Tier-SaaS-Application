package Data

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"atish/api/Cassandra"
"github.com/gorilla/mux"
"fmt"
"strconv"
)


type Identification struct {
   number    int
   common_name string
}


func Post(w http.ResponseWriter, r *http.Request) {
  	var errs []string
  	var gocqlUuid gocql.UUID

  	r.ParseForm()
  	number := r.FormValue("number")
  	common_name :=r.FormValue("common_name")
  	// fmt.Println(r.FormValue("number"))
  	fmt.Println(common_name)
  	fmt.Println(strconv.Atoi(number))
    if err := Cassandra.Session.Query("INSERT INTO cmpe281 (number,common_name) VALUES (?,?)", number,common_name).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      // created := true
    }
    // generate a unique UUID for this user
    if len(errs) == 0 {
	    fmt.Println("creating a new user")
	    gocqlUuid = gocql.TimeUUID()
	    json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
	}
  /*


    // write data to Cassandra
  }

  // depending on whether we created the user, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    fmt.Println("user_id", gocqlUuid)
  } else {
    fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }*/
}

func Get(w http.ResponseWriter, r *http.Request){

	var errs []string
  	// var gocqlUuid gocql.UUID
  	var idents []Identification

  	vars := mux.Vars(r)
    number := vars["number"]
    var common_name string
    // var jsonText = []byte(`[
        // {""}]`)

  	// number := r.FormValue("number")
  	// common_name :=r.FormValue("common_name")
  	// fmt.Println(r.FormValue("number"))
  	// fmt.Println(common_name)
  	// fmt.Println(strconv.Atoi(number))
  	if err := Cassandra.Session.Query(`SELECT * FROM cmpe281.cmpe281 WHERE number = ?`,
		number).Consistency(gocql.One).Scan(&number); err != nil {
		errs = append(errs, err.Error())
	}

	iter := Cassandra.Session.Query(`SELECT * FROM cmpe281.cmpe281 WHERE number = ?`, number).Iter()
	for iter.Scan(&common_name) {
		fmt.Println("Tweet:", common_name)
		// if err := json.Unmarshal([]byte(jsonText), &idents); err != nil {
		// 		errs = append(errs, err.Error())
		// }
		// errs = append(errs, err.Error())
		number_num,_ := strconv.Atoi(number)
	    idents = append(idents, Identification{number: number_num, common_name: common_name})
	}
	if err := iter.Close(); err != nil {
		errs = append(errs, err.Error())
	}
    // generate a unique UUID for this user
    if len(errs) == 0 {
	    // gocqlUuid = gocql.TimeUUID()
	    result,_ := json.Marshal(idents)
	    json.NewEncoder(w).Encode(result)
	}


}


func Put(w http.ResponseWriter, r *http.Request) {
	var errs []string
	r.ParseForm()
	number := r.FormValue("number")
  common_name :=r.FormValue("common_name")
	 fmt.Printf("Updating Emp with id = %s\n", emp.id)
	 if err := Cassandra.Session.Query("UPDATE cmpe281.cmpe281 SET number = ?, common_name = ? WHERE number = ?",
		 number,common_name).Exec(); err != nil {
		 fmt.Println("Error while updating user")
		 fmt.Println(err)
	}
}


func Delete(w http.ResponseWriter, r *http.Request) {
	var errs []string
	r.ParseForm()
	number := r.FormValue("number")
  common_name :=r.FormValue("common_name")
	fmt.Printf("Deleting User with number = %s\n", number_num)
	  if err := Cassandra.Session.Query("DELETE FROM cmpe281.cmpe281 WHERE number = ?", number_num).Exec(); err != nil {
		 fmt.Println("Error while deleting user")
		 fmt.Println(err)
	}
}