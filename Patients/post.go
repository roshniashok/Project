//The post function helps post the data received from the FormtoPatient
//function into the Cassandra Database if we had no errors. It also generates a
//unique UUID, a special functionality of the GOCQL module which helps in
//creating a new ID for every entry.And depending on whether we created an
//entry we return the resource ID as a JSON payload else we return our errors.
//Encoder Function is a serialiser which converts the Go structs into JSON
//data for enabling communication between any given entities over the internet.


package Patients

import (
"net/http"
"github.com/gocql/gocql"
"encoding/json"
"github.com/roshniashok/restapi/Cassandra"
"fmt"
)

func Post(w http.ResponseWriter, r *http.Request) {
  var errs []string
  var gocqlUuid gocql.UUID

  // FormToPatient() is included in Users/processing.go
  patient, errs := FormToPatient(r)

  // have we created a user correctly
  var created bool = false

  // if we had no errors from FormToPatient, we will
  // attempt to save our data to Cassandra
  if len(errs) == 0 {
    fmt.Println("creating a new user")

    // generate a unique UUID for this user
    gocqlUuid = gocql.TimeUUID()

    // write data to Cassandra
    if err := Cassandra.Session.Query(`
      INSERT INTO patients (id,age, diastolicbp,gender,heartrate,systolicbp,village) VALUES (?, ?, ?, ?, ?, ? ,?)`,
      gocqlUuid, patient.Age, patient.DiastolicBP, patient.Gender,patient.HeartRate,patient.SystolicBP, patient.Village).Exec(); err != nil {
      errs = append(errs, err.Error())
    } else {
      created = true
    }
  }

  // depending on whether we created the user, return the
  // resource ID in a JSON payload, or return our errors
  if created {
    fmt.Println("patient_id", gocqlUuid)
    json.NewEncoder(w).Encode(NewPatientResponse{ID: gocqlUuid})
  } else {
    fmt.Println("errors", errs)
    json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
  }
}
