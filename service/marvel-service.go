package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marvel/model"
	"net/http"
	"time"
)

// GetAllMarvelCharacters : bla bla bla
func GetAllMarvelCharacters() (*model.MarvelAPIResult, error) {

	data, err := GetAllMarvelCharactersByte()

	if err != nil {
		fmt.Printf("GetAllGodsByte ERROR: %s\n", err)
		return nil, err
	}

	marvelResult, err := unmarshallMarvelResultAPI(data)

	if err != nil {
		fmt.Printf("Unmarshall error: %s\n", err)
		return nil, err
	}

	return marvelResult, nil

}

// GetAllMarvelCharactersByte : bla bla bla
func GetAllMarvelCharactersByte() ([]byte, error) {

	// ******** DIRECT REQUEST with embedded params ****************
	// response, err := http.Get("http://gateway.marvel.com/v1/public/characters?ts=1&apikey=e2c27b4f76258e3d501e93b7ad4ba0fe&hash=3a6b95f6b43d507afa65fb92c04125f8")

	// ******** REQUEST with extra params ****************
	request, err := http.NewRequest("GET", "http://gateway.marvel.com/v1/public/characters", nil)

	q := request.URL.Query()
	q.Add("ts", "1")
	q.Add("apikey", "e2c27b4f76258e3d501e93b7ad4ba0fe")
	q.Add("hash", "3a6b95f6b43d507afa65fb92c04125f8")

	request.URL.RawQuery = q.Encode()
	fmt.Println(request.URL.String())

	timeout := time.Duration(5 * time.Second)

	client := http.Client{Timeout: timeout}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("HTTP ERROR %s\n", err)
		return nil, err
	}

	data, _ := ioutil.ReadAll(response.Body)

	return data, nil
}

func unmarshallMarvelResultAPI(body []byte) (*model.MarvelAPIResult, error) {
	var s = new(model.MarvelAPIResult)
	err := json.Unmarshal(body, &s)
	return s, err
}
