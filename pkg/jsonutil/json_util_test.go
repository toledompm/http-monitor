package jsonutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOpenJSONFile(t *testing.T) {
	fileName := "test.json"
	jsonData := `{"testKey": "testValue", "intTestKey": 1}`
	expectedJsonByteData := []byte(jsonData)
	err := ioutil.WriteFile(fileName, expectedJsonByteData, 0644)

	if err != nil {
		t.Errorf("Error writing json data to file: %v", err)
	}

	jsonByteData, err := OpenJsonFile(fileName)

	if err != nil {
		t.Errorf("Error opening json file: %v", err)
	}

	if string(jsonByteData) != jsonData {
		t.Errorf("Expected json data: %v, but got: %v", expectedJsonByteData, jsonByteData)
	}

	os.Remove(fileName)
}
