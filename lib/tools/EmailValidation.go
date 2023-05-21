package tools

import "regexp"

func ValidateEmail(e string) bool {

	pattern := "(.+)@wallpaper.Collect.app"
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(e)

	if len(matches) == 3 && len(matches) > 254 {
		return false
	}

	return true
}
