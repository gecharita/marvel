package test

import (
	"log"
	"os"
	"testing"
)

func TestEnvVariables(t *testing.T) {

	// var testVal = "success"
	// if testVal == "success" {
	// 	t.Log(testVal)
	// } else {
	// 	t.Error(testVal)
	// }
	// for _, e := range os.Environ() {
	// 	pair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(e)
	// 	fmt.Println(pair[0], pair[1])
	// }

	log.Println(os.Getenv("testCaseSecret"))
	log.Println("^^^^^^^^^^^^1^^^^^^^^^^^^^^^")
	log.Println(os.Getenv("super_secret"))
	log.Println("^^^^^^^^^^^^2^^^^^^^^^^^^^^^")

	log.Println(os.Environ())

}
