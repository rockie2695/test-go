package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var envkey map[string]string

func EnvSetup() {

	jsonFile, err := os.Open("config/env.json")

	if err != nil {
		panic("failed to open json file: " + err.Error())
	}

	rawJSON, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		panic("failed to read json file : " + err.Error())
	}

	err = json.Unmarshal(rawJSON, &envkey)
	if err != nil {
		panic("failed to unmarhsall env.json : " + err.Error())
	}

	defer jsonFile.Close()

	for key, value := range envkey {
		os.Setenv(key, value)
	}
}
