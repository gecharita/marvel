package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marvel/model"

	// "marvel/utils"
	"net/http"
	"os"
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

	// ******** REQUEST with extra params ****************
	request, err := http.NewRequest("GET", "http://gateway.marvel.com/v1/public/characters", nil)

	q := request.URL.Query()
	q.Add("ts", "1")
	q.Add("apikey", os.Getenv("marvel_api_key"))
	q.Add("hash", os.Getenv("marvel_api_hash"))
	// q.Add("apikey", utils.GetMarvelProperty("marvel.apikey"))
	// q.Add("hash", utils.GetMarvelProperty("marvel.hash"))

	request.URL.RawQuery = q.Encode()

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
