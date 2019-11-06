package test

import (
	"log"
	"os"
	"testing"
)

func TestEnvVariables(t *testing.T) {
	log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	log.Println(os.Getenv("super_secret"))
	log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
