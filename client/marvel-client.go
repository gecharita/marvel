package client

import (
	"log"
	"marvel/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port string = "9090"

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/marvel/getall").
		HandlerFunc(getMarvelCharactes)
}

func getMarvelCharactesByte(w http.ResponseWriter, r *http.Request) {
	marvelCharactersByte, err := service.GetAllMarvelCharactersByte()
	if err != nil {
		w.Write([]byte("ERROR \n"))
	} else {
		w.Write(marvelCharactersByte)
	}
}

func getMarvelCharactes(w http.ResponseWriter, r *http.Request) {
	marvelAPIResult, err := service.GetAllMarvelCharacters()
	if err != nil {
		w.Write([]byte("ERROR \n"))
	} else {
		w.Write([]byte(marvelAPIResult.Status + " "))
		codeStr := strconv.FormatInt(marvelAPIResult.Code, 10)
		w.Write([]byte(codeStr))
	}
}

// StartClient  bla bla
func StartClient() {
	log.Println("Server is running at localhost:" + port)
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
