package tools

import "github.com/asaskevich/govalidator"

func ValidationNumberPhone(numberPhone string) bool {

	if govalidator.IsE164(numberPhone) {
		return true
	}
	return false
}
