package config

import (
	"encoding/json"
	"log"
	"os"
)

//ParseJSONConfigToMap read .json file on input path and decode it to a map[string]interface{}
func ParseJSONConfigToMap(pathConfig string) map[string]interface{} {
	result := make(map[string]interface{})
	fileConfig, errOpen := os.Open(pathConfig)

	if errOpen != nil {
		log.Fatal("Cannot open " + pathConfig + ", please check and try again!")
	}
	defer fileConfig.Close()

	decoder := json.NewDecoder(fileConfig)
	errParse := decoder.Decode(&result)

	if errParse != nil {
		log.Fatal("Cannot parsing " + pathConfig + ", please check and try again!")
	}

	return result
}
