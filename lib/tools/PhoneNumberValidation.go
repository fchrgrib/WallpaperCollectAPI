package tools

import "regexp"

func ValidationNumberPhone(numberPhone string) bool {
	if len(numberPhone) < 9 {
		return false
	}
	var mustNotIn = regexp.MustCompile("^[a-zA-Z!#$%&'*+\\/=?^_`{|}~-]$")

	if !mustNotIn.MatchString(numberPhone) {
		return false
	}
	return true
}
