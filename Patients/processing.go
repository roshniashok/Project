// FormToPatient -- fills a User struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// User struct if successful
// array of strings of errors if any errors occur during processing

package Patients

import (
	"net/http"
	"strconv"
)


func FormToPatient(r *http.Request) (Patient, []string) {
	var patient Patient
	var errStr, ageStr, systolicbpStr , diastolicbpStr, heartrateStr  string
	var errs []string
	var err error

  	ageStr, errStr = processFormField(r, "age")
        if len(errStr) != 0 {
                errs = append(errs, errStr)
        } else {
                patient.Age, err = strconv.Atoi(ageStr)
                if err != nil {
                        errs = append(errs, "Parameter 'age' not an integer")
                }
        }


 	diastolicbpStr, errStr = processFormField(r, "diastolicbp")
        if len(errStr) != 0 {
                errs = append(errs, errStr)
        } else {
                patient.DiastolicBP, err = strconv.Atoi(diastolicbpStr)
                if err != nil {
                        errs = append(errs, "Parameter 'diastolicbp' not an integer")
                }
        }

	patient.Gender, errStr = processFormField(r, "gender")
	errs = appendError(errs, errStr)

	heartrateStr, errStr = processFormField(r, "heartrate")
        if len(errStr) != 0 {
                errs = append(errs, errStr)
        } else {
                patient.HeartRate, err = strconv.Atoi(heartrateStr)
                if err != nil {
                        errs = append(errs, "Parameter 'heartrate' not an integer")
                }
        }


	systolicbpStr, errStr = processFormField(r, "systolicbp")
        if len(errStr) != 0 {
                errs = append(errs, errStr)
        } else {
                patient.SystolicBP, err = strconv.Atoi(systolicbpStr)
                if err != nil {
                        errs = append(errs, "Parameter 'systolicbp' not an integer")
                }
	}
	patient.Village, errStr = processFormField(r, "village")
        errs = appendError(errs, errStr)
	 return patient, errs
}

func appendError(errs []string, errStr string) ([]string) {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
