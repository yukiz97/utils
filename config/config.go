package config

import (
	"encoding/json"
	"log"
	"os"
)

func ParseJsonConfigToStruct(pathConfig string) interface{} {
	var result interface{}
	fileConfig, errOpen := os.Open(pathConfig)

	if errOpen != nil {
		log.Fatal("Cannot open " + pathConfig + ", please check and try again!")
	}
	defer fileConfig.Close()

	decoder := json.NewDecoder(fileConfig)
	errParse := decoder.Decode(&result)

	if errParse != nil {
		log.Fatal("Cannot load json config file into input struct, please try again!")
	}

	return result
}
