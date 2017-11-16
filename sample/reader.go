package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JSONReader struct {
	configFile string
}

func NewJSONReader(configFile string) *JSONReader {
	return &JSONReader{configFile: configFile}
}

func (this *JSONReader) Read() (interface{}, error) {
	config := Config{} // TODO: be able to create an instance of this and populate it generically

	file, err := os.Open(this.configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		return nil, err
	}

	return config, nil
}
