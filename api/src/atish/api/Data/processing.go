package Data

import (
	"net/http"
	// "strconv"
)

// FormToUser -- fills a User struct with submitted form data
// params:
// r - request reader to fetch form data or url params (unused here)
// returns:
// User struct if successful
// array of strings of errors if any occur during processing
func FormToUser(r *http.Request) (User, []string) {
	var user User
	var errStr string
	var errs []string
	// var err error

	user.common_name, errStr = processFormField(r, "common_name")
	errs = appendError(errs, errStr)

	// user.number,errStr = strconv.Atoi(processFormField(r, "number"))
	
	return user, errs
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