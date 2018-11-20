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
  router.HandleFunc("/put", Data.UpdateHandler)
  router.HandleFunc("/delete", Data.DeleteHandler).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}