package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFileJson(filepath string) map[string]interface{} {
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config map[string]interface{}
	json.Unmarshal([]byte(byteValue), &config)

	return config
}
