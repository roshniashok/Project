//initiates cassandra cluster
//creates a new keyspace called patientapi

package Cassandra

import (
  "github.com/gocql/gocql"
  "fmt"
)

var Session *gocql.Session

func init() {

  var err error
  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "patientapi"
  Session, err = cluster.CreateSession()

  if err != nil {
    panic(err)
  }

  fmt.Println("cassandra init done")
}
