package testing

import (
	"github.com/lib/tools"
	"testing"
)

func TestEmailValidation(t *testing.T) {
	result := tools.ValidateEmail("Fahrian.Afdholi@wallpaper.Collect.app")
	expected := false
	if result == expected {
		t.Errorf("checking email isn't valid")
	}

	result = tools.ValidateEmail("Fahrian.Afdholi@gmail.com")
	expected = true
	if result == expected {
		t.Errorf("checking email that not contain @wallpaper.Collect.app isn't valid")
	}

	result = tools.ValidateEmail("@wallpaper.Collect.app")
	expected = true
	if result == expected {
		t.Errorf("checking validation before @ isn't valid")
	}

	result = tools.ValidateEmail("")
	expected = true
	if result == expected {
		t.Errorf("checking isNull isn't valid")
	}

}
