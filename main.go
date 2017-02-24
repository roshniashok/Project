package main

import (
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/roshniashok/restapi/Cassandra"
  "github.com/roshniashok/restapi/Patients"
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
  router.HandleFunc("/patients/new",Patients.Post)
  router.HandleFunc("/patients", Patients.Get)
  router.HandleFunc("/patients/{patient_uuid}", Patients.GetOne)
  log.Fatal(http.ListenAndServe(":8005", router))


}
func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

