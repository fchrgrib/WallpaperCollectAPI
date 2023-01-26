package tools

import (
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidateEmail(e string) bool {
	if len(e) < 3 || len(e) > 254 { //for checking the size of valid email
		return false
	}

	if !emailRegex.MatchString(e) { // for checking is email have the symbol like emailRegex or not
		return false
	}

	return true
}
