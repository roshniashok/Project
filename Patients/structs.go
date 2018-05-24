//Holds the various structs necessary for processing

package Patients

import (
	"github.com/gocql/gocql"
)

// User struct to hold profile data for our Patient
type Patient struct {
	ID gocql.UUID `json:"id"`
Age int `json:"age"`
DiastolicBP int `json:"diastolicbp"`
Gender  string `json:"gender"`
HeartRate int `json:"heartrate"`
SystolicBP int `json:"systolicbp"`
Village string `json:"village"`
}

// GetPatientResponse to form payload returning a single Patient struct
type GetPatientResponse struct {
	Patient Patient `json:"patient"`
}

// AllPatientsResponse to form payload of an array of Patient structs
type AllPatientsResponse struct {
	Patient []Patient `json:"patients"`
}

// NewPatientResponse builds a payload of new user resource ID
type NewPatientResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
