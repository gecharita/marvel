package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mythology/model"
	"net/http"
)

// GetAllGods : bla bla bla
func GetAllGods() (*model.GreekAPIResponse, error) {
	response, err := http.Get("https://anfi.tk/greekApi/person/getAll")
	if err != nil {
		fmt.Printf("HTTP ERROR %s\n", err)
		return nil, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	gods, err := unmarshallGreekGods([]byte(data))

	if err != nil {
		fmt.Printf("Unmarshall error: %s\n", err)
		return nil, err
	}

	// for i := 0; i < len(gods.GodList); i++ {
	// 	fmt.Printf("%d. %s %d\n", i, gods.GodList[i].Name, gods.GodList[i].PersonID)
	// }

	return gods, nil

}

// GetAllGodsByte : bla bla bla
func GetAllGodsByte() ([]byte, error) {
	response, err := http.Get("https://anfi.tk/greekApi/person/getAll")
	if err != nil {
		fmt.Printf("HTTP ERROR %s\n", err)
		return nil, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}

func unmarshallGreekGods(body []byte) (*model.GreekAPIResponse, error) {
	var s = new(model.GreekAPIResponse)
	err := json.Unmarshal(body, &s)
	return s, err
}
