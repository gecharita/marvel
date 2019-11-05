package utils

import (
	"github.com/magiconair/properties"
)

// GetMarvelProperty generic method to extract property value
func GetMarvelProperty(property string) string {
	return loadMarvelAPIProperties().MustGet(property)
}

func loadMarvelAPIProperties() *properties.Properties {
	return properties.MustLoadFile("${HOME}/go/src/marvel/resources/marvel-api.properties", properties.UTF8)
}
