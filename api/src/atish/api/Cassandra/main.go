package Cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
)

var Session *gocql.Session

func init() {
  var err error

  cluster := gocql.NewCluster("13.57.51.85","13.57.199.36","52.53.128.68","54.215.254.103","13.57.243.193")
  cluster.Keyspace = "cmpe281"
  Session, err = cluster.CreateSession()
  if err != nil {
    panic(err)
  }
  fmt.Println("cassandra init done")
}