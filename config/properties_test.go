package config

import (
	"testing"
)

func TestGetMarvelTestValue(t *testing.T) {
	var testVal string = GetGeneralProperty("test.status")
	if testVal == "success" {
		t.Log(testVal)
	} else {
		t.Error(testVal)
	}

}
