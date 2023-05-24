package testing

import (
	"github.com/libs/utils/validation"
	"testing"
)

func TestEmailValidation(t *testing.T) {
	result := validation.ValidateEmail("Fahrian.Afdholi@wallpaper.Collect.app")
	expected := true
	if result != expected {
		t.Errorf("checking email isn't valid")
	}

	result = validation.ValidateEmail("Fahrian.Afdholi@gmail.com")
	expected = false
	if result != expected {
		t.Errorf("checking email that not contain @wallpaper.Collect.app isn't valid")
	}

	result = validation.ValidateEmail("@wallpaper.Collect.app")
	expected = false
	if result != expected {
		t.Errorf("checking validation before @ isn't valid")
	}

	result = validation.ValidateEmail("")
	expected = false
	if result != expected {
		t.Errorf("checking isNull isn't valid")
	}

}

func TestPhoneNumberValidation(t *testing.T) {
	result := validation.ValidationNumberPhone("+6287724273282")
	expected := true
	if result != expected {
		t.Errorf("checking phoneNumber isn't valid")
	}

	result = validation.ValidationNumberPhone("")
	expected = false
	if result != expected {
		t.Errorf("checking isNull isn't valid")
	}

	result = validation.ValidationNumberPhone("+0087724273282")
	expected = false
	if result != expected {
		t.Errorf("checking code country isn't valid")
	}

	result = validation.ValidationNumberPhone("+6287724")
	expected = false
	if result != expected {
		t.Errorf("checking length of number isn't valid")
	}
}
