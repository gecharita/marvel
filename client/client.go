package client

import (
	"log"
	"mythology/service"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = "9090"

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/mythology/gods/getall").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing!")
	// w.Write([]byte("SUCCESS!! \n"))
	//gods, err := service.GetAllGods()

	gods, err := service.GetAllGodsByte()

	if err != nil {
		w.Write([]byte("ERROR \n"))
	} else {

		w.Write(gods)

		// w.Write([]byte("TODO RETURN GODS \n"))
		// for i := 0; i < len(gods.GodList); i++ {
		// 	fmt.Printf("%d. %s %d\n", i, gods.GodList[i].Name, gods.GodList[i].PersonID)
		// }
	}

}

// StartClient  bla bla
func StartClient() {
	log.Println("Server is running at localhost:" + port)
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
