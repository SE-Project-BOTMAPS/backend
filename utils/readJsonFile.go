package utils

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJsonFile(fileName string, model interface{}) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, model)
	if err != nil {
		panic(err)
	}
}
