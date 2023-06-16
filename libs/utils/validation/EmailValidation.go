package validation

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func ValidateUserOtherEmail(e string) bool {

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

func ValidateEmail(e string) bool {
	if govalidator.IsNull(e) {
		return false
	}
	if !govalidator.IsEmail(e) {
		return false
	}

	return true
}
