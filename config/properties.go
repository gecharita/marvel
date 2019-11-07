package config

import (
	"github.com/magiconair/properties"
)

var generalProperties = properties.MustLoadFile("./general.properties", properties.UTF8)

// GetGeneralProperty generic method to extract property value
func GetGeneralProperty(property string) string {
	return generalProperties.MustGet(property)
}
