package jsonutil

import (
	"io/ioutil"
	"os"
)

func OpenJsonFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
