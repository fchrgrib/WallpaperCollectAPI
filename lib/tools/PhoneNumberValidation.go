package tools

import "github.com/asaskevich/govalidator"

func ValidationNumberPhone(numberPhone string) bool {

	if len(numberPhone) < 11 {
		return false
	}
	if govalidator.IsE164(numberPhone) {
		return true
	}
	return false
}
