package tools

import "regexp"

func ValidationNumberPhone(numberPhone string) bool {
	if len(numberPhone) < 9 {
		return false
	}
	var mustNotIn = regexp.MustCompile("^[a-zA-Z!#$%&'*\\/=?^_`{|}~-]$")

	if mustNotIn.MatchString(numberPhone) {
		return false
	}

	pattern := "08(.+)"
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(numberPhone)

	if len(matches) == 0 {
		return false
	}

	pattern = "+(.+)"
	re = regexp.MustCompile(pattern)
	matches = re.FindStringSubmatch(numberPhone)

	if len(matches) == 0 {
		return false
	}

	return true
}
