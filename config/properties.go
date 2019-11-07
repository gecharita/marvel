package config

import (
	"log"
	"os"

	"github.com/magiconair/properties"
)

// var generalProperties = properties.MustLoadFile("${HOME}/go/src/marvel/config/general.properties", properties.UTF8)
var generalProperties = properties.MustLoadFile("./general.properties", properties.UTF8)

// GetGeneralProperty generic method to extract property value
func GetGeneralProperty(property string) string {
	log.Println(os.Getwd())
	return generalProperties.MustGet(property)
}
