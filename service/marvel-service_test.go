package service

import (
	"testing"
)

func TestGetAllMarvelCharacters(t *testing.T) {
	marvelResult, err := GetAllMarvelCharacters()
	if err != nil {
		t.Errorf("GetAllGodsByte ERROR: %s\n", err)
	} else {
		t.Logf("%+v\n", marvelResult)
	}
}
