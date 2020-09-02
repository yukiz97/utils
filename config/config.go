package config

import (
	"encoding/json"
	"log"
	"os"
)

func ParseJsonConfigToStruct(inputStruct interface{}, pathConfig string) interface{} {
	fileConfig, errOpen := os.Open(pathConfig)

	if errOpen != nil {
		log.Fatal("Cannot open " + pathConfig + ", please check and try again!")
	}
	defer fileConfig.Close()

	decoder := json.NewDecoder(fileConfig)
	errParse := decoder.Decode(&inputStruct)

	if errParse != nil {
		log.Fatal("Cannot load json config file into input struct, please try again!")
	}

	return inputStruct
}
