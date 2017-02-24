package Patients

import (
  "net/http"
  "github.com/gocql/gocql"
  "encoding/json"
  "github.com/roshniashok/restapi/Cassandra"
  "github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
  var patientList []Patient
  m := map[string]interface{}{}

  query := "SELECT id,gender,systolicbp,diastolicbp,age FROM patients"
  iterable := Cassandra.Session.Query(query).Iter()
  for iterable.MapScan(m) {
    patientList = append(patientList, Patient{
      ID: m["id"].(gocql.UUID),
      Gender: m["gender"].(string),
      SystolicBP: m["systolicbp"].(int),
      DiastolicBP: m["diastolicbp"].(int),
      Age: m["age"].(int),
    })
    m = map[string]interface{}{}
  }

  json.NewEncoder(w).Encode(AllPatientsResponse{Patient: patientList})
}

func GetOne(w http.ResponseWriter, r *http.Request) {
  var patient Patient
  var errs []string
  var found bool = false

  vars := mux.Vars(r)
  id := vars["patient_uuid"]

  uuid, err := gocql.ParseUUID(id)
  if err != nil {
    errs = append(errs, err.Error())
  } else {
    m := map[string]interface{}{}
    query := "SELECT id,gender,systolicbp,diastolicb,age  FROM users WHERE id=? LIMIT 1"
    iterable := Cassandra.Session.Query(query, uuid).Consistency(gocql.One).Iter()
    for iterable.MapScan(m) {
      found = true
      patient = Patient{
        ID: m["id"].(gocql.UUID),
        Gender: m["gender"].(string),
        SystolicBP: m["systolicbp"].(int),
        DiastolicBP: m["diastolicbp"].(int),
        Age: m["age"].(int),
      }
    }
    if !found {
      errs = append(errs, "Patient not found")
    }
  }

  if found {
    json.NewEncoder(w).Encode(GetPatientResponse{Patient: patient})
  } else {
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}

