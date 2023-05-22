package tools

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func ValidateEmail(e string) bool {

	if govalidator.IsNull(e) {
		return false
	}
	if !govalidator.IsEmail(e) {
		return false
	}

	pattern := `(.+)@wallpaper.Collect.app`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(e)

	if len(matches) == 0 {
		return false
	}

	return true
}
