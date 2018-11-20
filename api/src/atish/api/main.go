/*package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("18.144.62.162", "54.215.252.168", "13.57.7.146","52.53.128.5","13.56.81.44")
	cluster.Keyspace = "cmpe281"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// var t,s,pk,v string
	var common_name string

	// list all tweets
	for i := 0; i < 50; i++ {
		iter := session.Query(`SELECT * FROM t`).Iter()
		for iter.Scan(&common_name) {
			fmt.Println("common_name:", common_name)
		}
		if err := iter.Close(); err != nil {
			log.Fatal(err)
		}
	}

	//Insert
	for i := 0; i < 10; i++ {
		if err := session.Query("INSERT INTO cmpe281 (number,common_name) VALUES (?,'test')", i).Exec(); err != nil {
			log.Fatal(err)
		}else{
			fmt.Println("Inserted row:%d",i);		
		}
	}

	//Update
	var x string
	x="test1"
	if err := session.Query("UPDATE cmpe281 set common_name=?", x).Exec(); err != nil {
			log.Fatal(err)
	}

	//Delete
	var y string
	y="test2"
	if err := session.Query("DELETE *from cmpe281 where common_name=?", y).Exec(); err != nil {
			log.Fatal(err)
	}	
	fmt.Println("Done")
}
*/

package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
   "atish/api/Cassandra"
   "atish/api/Data"
)


type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
  CassandraSession := Cassandra.Session
  defer CassandraSession.Close()
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", heartbeat)
  router.HandleFunc("/data", Data.Post)
  router.HandleFunc("/get/{number}", Data.Get)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}