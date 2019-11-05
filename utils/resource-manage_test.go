package utils

import (
	"testing"
)

func TestGetMarvelTestValue(t *testing.T) {
	var testVal string = GetMarvelProperty("marvel.test")
	if testVal == "success" {
		t.Log(testVal)
	} else {
		t.Error(testVal)
	}

}
